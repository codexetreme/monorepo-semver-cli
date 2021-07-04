package main

import (
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

var (
	prj1, _      = NewProject("Prj1", "prj", "a")
	prj2, _      = NewProject("Prj2", "prj", "b")
	testRegistry = NewRegistry(SLASH)
)

func TestRegistry_AddProject(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	type args struct {
		p *Project
	}
	tests := []struct {
		name     string
		registry *Registry
		args     args
		wantErr  bool
	}{
		{
			name:     "test-1",
			registry: testRegistry,
			args: args{
				p: prj1,
			},
			wantErr: false,
		},
		{
			name:     "test-1",
			registry: testRegistry,
			args: args{
				p: prj2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.registry
			if err := r.AddProject(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("AddProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegistry_RemoveProject(t *testing.T) {
	log.SetLevel(log.InfoLevel)
	// delete valid Object
	t.Run("vaild Id", func(t *testing.T) {
		testRegistry.ClearRegistry()
		_ = testRegistry.AddProject(prj1)
		objectsInRegistry := len(testRegistry.projects)
		testRegistry.RemoveProject(prj1.id)
		if objectsInRegistry-1 != len(testRegistry.projects) {
			t.Errorf("did not delete object")
		}

	})
	// delete invalid Object
	t.Run("invaild Id", func(t *testing.T) {
		testRegistry.ClearRegistry()
		_ = testRegistry.AddProject(prj1)
		objectsInRegistry := len(testRegistry.projects)
		testRegistry.RemoveProject("non existent Id")
		if objectsInRegistry-1 == len(testRegistry.projects) {
			t.Errorf("object was deleted even when it did not exist")
		}

	})
}

func TestRegistry_isPrefixUniqueForProject(t *testing.T) {
	_ = testRegistry.AddProject(prj1)

	type args struct {
		prefix string
	}
	tests := []struct {
		name     string
		registry *Registry
		args     args
		wantErr  bool
	}{
		{
			name:     "valid prefixes",
			registry: testRegistry,
			args:     args{prefix: "prj2"},
			wantErr:  false,
		},
		{
			name:     "invalid prefix",
			registry: testRegistry,
			args:     args{prefix: "prj"},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.registry
			if err := r.isPrefixUniqueForProject(tt.args.prefix); (err != nil) != tt.wantErr {
				t.Errorf("isPrefixUniqueForProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegistry_GetProjectById(t *testing.T) {

	testRegistry.ClearRegistry()
	_ = testRegistry.AddProject(prj1)
	type args struct {
		id ProjectId
	}
	tests := []struct {
		name     string
		registry *Registry
		args     args
		want     *Project
		wantErr  bool
	}{
		{
			name:     "valid project Id",
			registry: testRegistry,
			args:     args{id: prj1.id},
			want:     prj1,
			wantErr:  false,
		},
		{
			name:     "valid project Id",
			registry: testRegistry,
			args:     args{id: prj2.id},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.registry
			got, err := r.GetProjectById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjectById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProjectById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProject(t *testing.T) {
	t.Run("valid project creation", func(t *testing.T) {
		_, err := NewProject("test_1", "test_1", ".")
		if err != nil {
			t.Errorf("failed to create project: %s", err)
		}
	})
}
