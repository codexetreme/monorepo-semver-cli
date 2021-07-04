package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type DelimChar uint8

const (
	COLON DelimChar = iota
	SLASH
)

var delimCharMap = map[DelimChar]string{
	COLON: ":",
	SLASH: "/",
}

type ProjectId string
type registryMap map[ProjectId]*Project
type Project struct {
	id     ProjectId
	name   string
	prefix string
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

func NewProject(name string, prefix string) *Project {
	p := &Project{
		id:     ProjectId(uuid.New().String()),
		name:   name,
		prefix: prefix,
	}
	return p
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
)

func (r Registry) AddProject(p *Project) error {

	if err := r.isPrefixUniqueForProject(p.prefix); err != nil {
		return fmt.Errorf("cannot add project to registry: %s", err)
	}
	log.Debugf("adding project with id: %s", p.id)
	r.projects[p.id] = p
	return nil
}

func (r Registry) RemoveProject(id ProjectId) {
	log.Debugf("removing project with id: %s", id)
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
			return errors.New(fmt.Sprintf("prefix already exists for project with id %s", prjInReg.id))
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
