package main

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	log "github.com/sirupsen/logrus"
)

func (r *Repository) cloneRepo(path string) {
	repo, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}
	r.repoHandle = repo

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

type Repository struct {
	repoHandle *git.Repository
	useSemtag  bool
}

// get all tags with prefix

func (r *Repository) getAllTagsWithGlob(glob string) ([]string, error) {
	ref, err := r.repoHandle.Tag(glob)
	if err != nil {
		log.WithError(err).Errorf("err")
		return nil, err
	}
	log.WithField("got", ref.String()).Info("ref")
	return nil, nil
}
