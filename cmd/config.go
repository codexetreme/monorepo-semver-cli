package cmd

import "github.com/spf13/viper"

var yaml = []byte(`
version: v1
giturl: https://github.com/codexetreme/monorepo-semver-cli
useV: false # make tags like vX.Y.Z, set it globally here, but each project can override this
projects:
- name: prj1 # name of the project
  relativePath: docs # location of the project relative to root
  prefix: prj1-tags # used to override the default prefix
  useV: true

- name: prj2 # name of the project
  relativePath: docs2 # location of the project relative to root
  #prefix: prj1-tags # used to override the default prefix
`)

type projectOptions struct {
    name string
    relativePath string
    prefix string
    useV string
}

type semverOptions struct {
    prefixWithV bool
}

type options struct {
    useV bool
    giturl string
    projectsOpts []projectOptions
}

func createConfig() {
    //viper.SetDefault("semver.prefixWithV", false)
    //viper.SetDefault("semver.prefixWithV", false)
    viper.WriteConfigAs(".msc/config.yml")
}
