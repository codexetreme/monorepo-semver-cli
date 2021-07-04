package main

import "path"

// config for all operations

type ConfigFileOptions struct {
	ConfigFileName    string
	ConfigDirName     string
	ConfigDirLocation string
}

func (c ConfigFileOptions) GetFullPath() string {
	return path.Join(c.ConfigDirLocation,
		c.ConfigDirName,
		c.ConfigFileName)
}

type ConfigFileOptionsOperations interface {
	GetFullPath() string
}

type ProjectConfigOptions struct {
	ConfigFileOptions
}
type CliConfigOptions struct {
	ConfigFileOptions
	envPrefix string
}

type GlobalConfigOptions struct {
	CliConfigOptions
	ProjectConfigOptions
}

var (
	DefaultConfigOptions = GlobalConfigOptions{
		CliConfigOptions: CliConfigOptions{
			ConfigFileOptions: ConfigFileOptions{
				configFileName:    "cli.config.yml",
				configDirName:     ".msc",
				configDirLocation: ".",
			},
			envPrefix: "MSC",
		},
		ProjectConfigOptions: ProjectConfigOptions{
			ConfigFileOptions{
				configFileName:    "projects.config.yml",
				configDirName:     ".msc",
				configDirLocation: ".",
			},
		},
	}
)
