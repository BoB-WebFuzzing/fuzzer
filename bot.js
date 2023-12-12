const puppeteer = require('puppeteer');
const fs = require('fs'); //for working with files
const getenv = require('getenv');
const url = require('url');

const filePath = '/tmp/httpreqr.pid';
const cur_input = 'fuzzing-0/output/.cur_input';

//load cookie function
const loadCookie = async (page) => {
    if (getenv('LOGIN_COOKIE')) {
        const cookiesArray = getenv('LOGIN_COOKIE').split(';').map(cookie => {
            const [key, value] = cookie.trim().split('=');
            // 추가하기 전에 key와 value가 비어있는지 확인
            if (key && value) {
                return { name: key, value: value, domain:url.parse(process.argv[2]).hostname};
            }
            return null; // key 또는 value가 비어있다면 null 반환
        }).filter(Boolean); // null이 아닌 항목만 필터링

        await page.setCookie(...cookiesArray);
    }
}

let parsedStrings; // Declare parsedStrings outside the try block
try {
    // 파일을 동기적으로 읽어서 내용을 가져옴
    const data = fs.readFileSync(cur_input, 'utf8');

    // Null 문자를 기준으로 문자열을 파싱
    parsedStrings = data.split('\x00');

} catch (err) {
    console.error(`Error reading file: ${err}`);
}

(async () => {
    const browser = await puppeteer.launch({
        args: [
            '--disable-features=site-per-process', '--no-sandbox', '--disable-setuid-sandbox'
            ],
        "defaultViewport": null,
        headless: 'new',
    });

    const page = await browser.newPage();
    
    const url = process.argv[2] + '?' + parsedStrings[1];

    await loadCookie(page); //load cookie

    page.on('dialog', async (dialog) => {
        console.log(`Dialog message: ${dialog.message()}`);
        if (dialog.type() === 'alert') {
            if (dialog.message() === 'WTFTEST') {
                console.log('XSS detected');
                // 파일을 읽어서 내용을 가져옴
                fs.readFile(filePath, 'utf8', (err, data) => {
                    if (err) {
                        console.error(`Error reading file: ${err}`);
                        return;
                    }

                    // 파일 내용을 정수로 변환
                    const pid = parseInt(data.trim(), 10);

                    // 정수값 출력 또는 다른 작업 수행
                    console.log(`Read PID from file: ${pid}`);

                    // Check if the process with the specified PID exists before killing it
                    if (pid && Number.isInteger(pid)) {
                        try {
                            process.kill(pid, 'SIGSEGV');
                            console.log(`Process with PID ${pid} killed.`);
                        } catch (killError) {
                            console.error(`Error killing process with PID ${pid}: ${killError.message}`);
                        }
                    } else {
                        console.log(`Invalid PID: ${pid}`);
                    }
                });
            } else {
                console.log('alert But not XSS');
            }
        }
        await dialog.dismiss();
    });

    await page.goto(url);

    await browser.close();
  })();
