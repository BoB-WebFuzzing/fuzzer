# fuzzer
[WebTheFuzzer](https://github.com/BoB-WebFuzzing/WTFuzzer-PHP) control tower

# Usage
## Install Golang

```bash
wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

## Make Instructions

```bash
make # make fuzzer
make clean # clean all
make clean-dir # clean all except fuzzer
```

## Required Files
### config.json

```json
{
    "testname": "test",
    "afl_path": "/afl",
    "target_binary": "/usr/local/bin/php-cgi ",
    "base_url": "http://localhost:{PORT}/",
    "base_port": 80,
    "timeout" : 60,
    "ld_library_path": "/lib",
    "ld_preload": "/lib/hook_recv.so",
    "memory": "8G",
    "first_crash": true,
    "cores": 1,
    "login": {
        "url": "http://localhost:{PORT}/login",
        "port": 80,
        "postData": "id=admin&pw=admin",
        "getData": "id=guest&pw=guest",
        "positiveHeaders": {"content-type": "Application/json"},
        "positiveBody": "",
        "method": "POST",
        "loginSessionCookie" : "PHPSESSID"
    }
}
```

### request_data.json
This file can be generated by [WTF-Crawler](https://github.com/BoB-WebFuzzing/WTF-Crawler)

## Run fuzzer

```bash
./fuzzer /path/to/config.json /path/to/request_data.json
```
