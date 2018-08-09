package supply

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack"
)

type Supplier struct {
	installer *libbuildpack.Installer
	stager    *libbuildpack.Stager
}

func New(stager *libbuildpack.Stager, manifest *libbuildpack.Manifest, installer *libbuildpack.Installer, logger *libbuildpack.Logger, command *libbuildpack.Command) *Supplier {
	return &Supplier{installer: installer, stager: stager}
}

func (s *Supplier) Run() error {
	dep := libbuildpack.Dependency{Name: "hwc", Version: "12.0.0"}
	dir := filepath.Join(s.stager.DepDir(), "hwc")
	if err := s.installer.InstallDependency(dep, dir); err != nil {
		return err
	}

	if err := s.stager.AddBinDependencyLink(filepath.Join(dir, "hwc.exe"), "hwc.exe"); err != nil {
		return err
	}

	return nil
}
