package main

import (
    "os/exec"
)

func runBot(URL string) {
    cmd := exec.Command("node", "bot.js", URL)

    go exitAFL(cmd)
}
