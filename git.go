package main

import (
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/plumbing/object"
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
    r, err := git.PlainOpen(".")
    if err != nil {
        log.Fatal(err)
    }

    // ... retrieving the HEAD reference
    ref, err := r.Head()


    // ... retrieves the commit history
    cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})

    r.Tags()

    // ... just iterates over the commits
    var cCount int
    err = cIter.ForEach(func(c *object.Commit) error {
        cCount++
        log.Info(c.String())
        return nil
    })
}
