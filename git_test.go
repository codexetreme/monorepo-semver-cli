package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func Test_cloneRepo(t *testing.T) {
	r := Repository{}
	r.cloneRepo("/tmp/foo")
}

func TestShowLog(t *testing.T) {
	ShowLog()
}

func TestRepository_getAllTagsWithGlob(t *testing.T) {
	r := Repository{}
	err := os.RemoveAll("/tmp/foo")
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	r.cloneRepo("/tmp/foo")
	// works, it gets a hash object
	r.getAllTagsWithGlob("v5.4.2")

	r.getAllTagsWithGlob("v5*")

}
