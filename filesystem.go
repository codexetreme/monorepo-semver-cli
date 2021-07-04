package main

import "io"

type PersistenceConfig struct {
}

type PersistenceOps interface {
	Save(writer io.Writer) error
	Read(reader io.Reader, location string) ([]byte, error)
	MkDir(name string, location string) error
}
