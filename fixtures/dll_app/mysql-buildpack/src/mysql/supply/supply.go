package supply

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack"
)

type Stager interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/stager.go
	BuildDir() string
	DepDir() string
	DepsIdx() string
	DepsDir() string
	AddBinDependencyLink(string, string) error
}

type Manifest interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/manifest.go
	AllDependencyVersions(string) []string
	DefaultVersion(string) (libbuildpack.Dependency, error)
}

type Installer interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/installer.go
	InstallDependency(libbuildpack.Dependency, string) error
	InstallOnlyVersion(string, string) error
}

type Command interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/command.go
	Execute(string, io.Writer, io.Writer, string, ...string) error
	Output(dir string, program string, args ...string) (string, error)
}

type Supplier struct {
	Manifest  Manifest
	Installer Installer
	Stager    Stager
	Command   Command
	Log       *libbuildpack.Logger
}

func (s *Supplier) Run() error {
	s.Log.BeginStep("Supplying mysql")
	fmt.Println("Hi I am the supplier and I am being run!")

	dep := libbuildpack.Dependency{Name: "mysql", Version: "0.0.1"}
	if err := s.Installer.InstallDependency(dep, s.Stager.DepDir()); err != nil {
		return err
	}

	// /tmp/deps/0/hwc/hwc.exe
	if err := s.Stager.AddBinDependencyLink(filepath.Join(s.Stager.DepDir(), "MySql.Data.dll"), "MySql.Data.dll"); err != nil {
		fmt.Printf("SYMLINK ERROR: %s", err.Error())
		return err
	}
	return nil
}
