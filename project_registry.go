package main

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io"
)

type DelimChar uint8

const (
	COLON DelimChar = iota
	SLASH
)

var (
	delimCharMap = map[DelimChar]string{
		COLON: ":",
		SLASH: "/",
	}
	validate = validator.New()
)

type ProjectId string
type registryMap map[ProjectId]*Project
type Project struct {
	id           ProjectId `validate:"uuid4"`
	name         string    `validate:"required,uuid4"`
	prefix       string    `validate:"required"`
	relativePath string    `validate:"required"`
}

func (p Project) Name() string {
	return p.name
}

func (p Project) Prefix() string {
	return p.prefix
}

func (p Project) Id() ProjectId {
	return p.id
}

func NewProject(name string, prefix string, relativePath string) (*Project, error) {

	if err := validate.Var(name, "required,hostname_rfc1123"); err != nil {
		return nil, fmt.Errorf("err: %s", err)
	}

	if err := validate.Var(prefix, "required,hostname_rfc1123"); err != nil {
		return nil, fmt.Errorf("err: %s", err)
	}

	if err := validate.Var(relativePath, "required"); err != nil {
		return nil, fmt.Errorf("err: %s", err)
	}

	p := &Project{
		id:           ProjectId(uuid.New().String()),
		name:         name,
		prefix:       prefix,
		relativePath: relativePath,
	}

	return p, nil
}

type (
	Registry struct {
		projects registryMap
		delim    DelimChar
	}
	RegistryRead interface {
		GetProjectById(id ProjectId) (*Project, error)
	}
	RegistryWriter interface {
		AddProject(p *Project) error
		RemoveProject(id ProjectId)
		SetDelim(delim string) error
		ClearRegistry()
	}
	RegistryPersistence interface {
		toYaml() []byte
		save(writer io.Writer)
		load(location string)
		generate(location string)
	}
)

func (r Registry) AddProject(p *Project) error {

	if err := r.isPrefixUniqueForProject(p.prefix); err != nil {
		return fmt.Errorf("cannot add project to registry: %s", err)
	}
	log.Debugf("adding project with Id: %s", p.id)
	r.projects[p.id] = p
	return nil
}

func (r Registry) RemoveProject(id ProjectId) {
	log.Debugf("removing project with Id: %s", id)
	delete(r.projects, id)
}

func (r Registry) GetProjectById(id ProjectId) (*Project, error) {
	if val, ok := r.projects[id]; ok {
		return val, nil
	}
	return nil, errors.New("cannot find project")
}

func (r Registry) isPrefixUniqueForProject(prefix string) error {
	for _, prjInReg := range r.projects {
		if prjInReg.prefix == prefix {
			return fmt.Errorf("prefix already exists for project with Id %s", prjInReg.id)
		}
	}
	return nil
}

func (r Registry) ClearRegistry() {
	r.projects = make(registryMap)
}

func NewRegistry(delim DelimChar) *Registry {
	r := &Registry{
		projects: make(registryMap),
		delim:    delim,
	}
	return r
}
