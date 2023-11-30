const puppeteer = require('puppeteer');
const fs = require('fs');

const filePath = '/tmp/httpreqr.pid';
const cur_input = 'fuzzing-0/output/.cur_input';

(async () => {
    let parsedStrings; // Declare parsedStrings outside the try block
    let parsedWords;
    try {
        // 파일을 동기적으로 읽어서 내용을 가져옴
        const data = fs.readFileSync(cur_input, 'utf8');

        // Null 문자를 기준으로 문자열을 파싱
        parsedStrings = data.split('\x00');

    } catch (err) {
        console.error(`Error reading file: ${err}`);
    }

    if (parsedStrings[1].includes('<iframe') == true || parsedStrings[1].includes('<script>') == true) {
        const browser = await puppeteer.launch({
            args: ['--no-sandbox'],
            headless: 'new',
        });
        const page = await browser.newPage();

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
                        if (pid && Number.isInteger(pid, '11')) {
                            try {
                                process.kill(pid);
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


        const url = process.argv[2] + '?' + parsedStrings[1];
        await page.goto(url);

        await browser.close();
    } else {
        console.log('Not XSS mutated. Headless browser not launched.');
    }
})();
