package main

import (
	"os"
	"time"

	"github.com/cloudfoundry/hwc-buildpack/src/hwc/finalize"
	_ "github.com/cloudfoundry/hwc-buildpack/src/hwc/hooks"

	"github.com/cloudfoundry/libbuildpack"
)

func init() {
	os.Setenv("TZ", "Africa/Cairo")
}

func main() {
	logger := libbuildpack.NewLogger(os.Stdout)

	buildpackDir, err := libbuildpack.GetBuildpackDir()
	if err != nil {
		logger.Error("Unable to determine buildpack directory: %s", err.Error())
		os.Exit(9)
	}

	manifest, err := libbuildpack.NewManifest(buildpackDir, logger, time.Now())
	if err != nil {
		logger.Error("Unable to load buildpack manifest: %s", err.Error())
		os.Exit(10)
	}

	stager := libbuildpack.NewStager(os.Args[1:], logger, manifest)

	if err = manifest.ApplyOverride(stager.DepsDir()); err != nil {
		logger.Error("Unable to apply override.yml files: %s", err)
		os.Exit(17)
	}

	if err := stager.SetStagingEnvironment(); err != nil {
		logger.Error("Unable to setup environment variables: %s", err.Error())
		os.Exit(11)
	}

	harmonizer := finalize.NewHarmonizer(logger, stager.BuildDir(), stager.DepDir())

	f := finalize.Finalizer{
		Manifest:   manifest,
		Stager:     stager,
		Command:    &libbuildpack.Command{},
		Log:        logger,
		Harmonizer: harmonizer,
	}

	if err := f.Run(); err != nil {
		os.Exit(12)
	}

	if err := libbuildpack.RunAfterCompile(stager); err != nil {
		logger.Error("After Compile: %s", err.Error())
		os.Exit(13)
	}

	if err := stager.SetLaunchEnvironment(); err != nil {
		logger.Error("Unable to setup launch environment: %s", err.Error())
		os.Exit(14)
	}

	stager.StagingComplete()
}
