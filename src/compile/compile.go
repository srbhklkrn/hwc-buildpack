package main

import (
	"archive/zip"
	. "common"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "github.com/go-yaml/yaml"
	"github.com/sesmith177/go-ce-test/buildpack"
)

func main() {
	buildDir, _, err := parseArgs(os.Args[1:])
	CheckErr(err)

	var m buildpack.Manifest

	buildpackDir, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), ".."))
	CheckErr(err)

	manifestData, err := ioutil.ReadFile(filepath.Join(buildpackDir, "manifest.yml"))
	CheckErr(err)

	err = yaml.Unmarshal(manifestData, &m)
	CheckErr(err)

	CheckWebConfig(buildDir)

	hwcDir := filepath.Join(buildDir, ".cloudfoundry")
	err = os.MkdirAll(hwcDir, 0700)
	CheckErr(err)

	downloader := buildpack.NewDownloader(hwcDir, &m)
	dest, err := downloader.Fetch(buildpack.Dependency{Name: "hwc", Version: "1.0.0"}, "hwc.zip")
	CheckErr(err)

	err = unzipHWC(dest, filepath.Join(hwcDir, "hwc.exe"))

	CheckErr(err)

	os.Exit(0)
}

func unzipHWC(src string, output string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		f2, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		defer f2.Close()

		_, err = io.Copy(f2, rc)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseArgs(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", errors.New("Invalid usage. Expected: compile.exe <build_dir> <cache_dir>")
	}

	return args[0], args[1], nil
}
