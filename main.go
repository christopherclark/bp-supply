package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/pborman/getopt/v2"
)

type SupplyPlugin struct{}

func (p *SupplyPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "bp-supply",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 9,
		},
		Commands: []plugin.Command{
			{
				Name:     "bp-supply",
				HelpText: "Build a supply buildpack",
				UsageDetails: plugin.Usage{
					Usage: "cf bp-supply [-p PATH] [-v VERSION] [NAME] [BUILDPACK_FILE]",
					Options: map[string]string{
						"-p": "Path to create buildpack from",
						"-v": "Version of buildpack",
					},
				},
			},
		},
	}
}

func (p *SupplyPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	path := getopt.StringLong("path", 'p', ".", "Path to create buildpack from")
	version := getopt.StringLong("version", 'v', "0.0.0", "Version of buildpack")
	var opts = getopt.CommandLine
	if err := opts.Getopt(args, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		// s.usage()
		os.Exit(1)
	}
	args = opts.Args()

	if len(args) == 0 {
		// Installing, do nothing
	} else if len(args) == 2 {
		p.CreateBuildpack(*path, args[0], *version, args[1])
	} else {
		fmt.Println("Args:", args)
		panic("FIXME: usage")
	}
}

func main() {
	plugin.Start(&SupplyPlugin{})
}

func (p *SupplyPlugin) CreateBuildpack(path, bpName, bpVersion, zipPath string) error {
	newfile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	if err := addFileToZip(zipWriter, "bin/compile", binCompile, 0755); err != nil {
		return err
	}
	if err := addFileToZip(zipWriter, "bin/supply", binSupply(bpName, bpVersion), 0755); err != nil {
		return err
	}

	if err := filepath.Walk(path, func(filename string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		zipfile, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		filename, err = filepath.Rel(path, filename)
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name = filepath.Join("content", filename)

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(filepath.Join(path, filename))
			if err != nil {
				return err
			}
			if _, err = writer.Write([]byte(target)); err != nil {
				return err
			}
		} else {
			if _, err = io.Copy(writer, zipfile); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

const binCompile = `#!/bin/bash

echo 'Warning: this buildpack can only be run as a supply buildpack, it can not be run alone' >&2
exit 1
`

const binSupplyTemplate = `#!/bin/bash
set -euo pipefail

BUILDPACK_DIR=$(dirname $(readlink -f ${BASH_SOURCE%/*}))
BUILD_DIR=$1
CACHE_DIR=$2
DEPS_DIR=$3
DEPS_IDX=$4

echo "-----> {{.Name}} Buildpack version {{.Version}}"

rsync -a $BUILDPACK_DIR/content/ $DEPS_DIR/$DEPS_IDX/

echo "config: {}" > $DEPS_DIR/$DEPS_IDX/config.yml
echo "name: {{.Name}}" >> $DEPS_DIR/$DEPS_IDX/config.yml
`

func binSupply(bpName, bpVersion string) string {
	tmpl, err := template.New("binSupply").Parse(binSupplyTemplate)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, map[string]string{"Name": bpName, "Version": bpVersion}); err != nil {
		panic(err)
	}
	return tpl.String()
}

func addFileToZip(zipWriter *zip.Writer, path, contents string, mode int) error {
	header := &zip.FileHeader{Name: path, Method: zip.Deflate}
	header.SetMode(os.FileMode(mode))
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	if _, err := writer.Write([]byte(contents)); err != nil {
		return err
	}
	return nil
}
