package main

import (
    "github.com/go-git/go-git/v5"
    log "github.com/sirupsen/logrus"
    "os"
)

func cloneRepo() {
    _, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
        URL:      "https://github.com/go-git/go-git",
        Progress: os.Stdout,
    })
    if err != nil {
        log.Fatal(err)
    }
}

func ShowLog() {
    git.
    _, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
        URL:      "https://github.com/go-git/go-git",
        Progress: os.Stdout,
    })
    if err != nil {
        log.Fatal(err)
    }
}
