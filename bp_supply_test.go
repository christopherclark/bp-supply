package main_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/cloudfoundry/bp-supply"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BpSupply", func() {
	var (
		tmpDir  string
		subject *SupplyPlugin
	)
	BeforeEach(func() {
		var err error
		tmpDir, err = ioutil.TempDir("", "bp-supply-zip-")
		Expect(err).ToNot(HaveOccurred())

		subject = &SupplyPlugin{}
	})
	AfterEach(func() {
		Expect(os.RemoveAll(tmpDir)).To(Succeed())
	})

	Describe("zipfile", func() {
		var zipFile string
		BeforeEach(func() {
			zipFile = filepath.Join(tmpDir, "bp.zip")
			Expect(subject.CreateBuildpack("fixtures/simple", "Simple", "0.1.2", zipFile)).To(Succeed())
		})
		It("is created", func() {
			Expect(zipFile).To(BeARegularFile())
		})

		It("has bin/compile which refuses to run", func() {
			Expect(ZipFileContent(zipFile, "bin/compile")).To(ContainSubstring("Warning: this buildpack can only be run as a supply buildpack, it can not be run alone"))
		})

		It("has bin/supply which copies content/ to DEPS_DIR/DEPS_IDX/", func() {
			content, err := ZipFileContent(zipFile, "bin/supply")
			Expect(err).ToNot(HaveOccurred())
			Expect(content).To(ContainSubstring("---> Simple Buildpack version 0.1.2"))
			Expect(content).To(ContainSubstring("rsync -a $BUILDPACK_DIR/content/ $DEPS_DIR/$DEPS_IDX/"))
		})

		It("has bin/jane copied to content/bin/jane", func() {
			Expect(ZipFileContent(zipFile, "content/bin/jane")).To(ContainSubstring("Jane's Contents"))
		})

		It("has lib/fred.6.so copied to content/lib/fred.6.so", func() {
			Expect(ZipFileContent(zipFile, "content/lib/fred.6.so")).To(Equal("Contents of lib/fred\n"))
		})

		It("has lib/fred.so as a symlink at content/lib/fred.so", func() {
			Expect(ZipFileContent(zipFile, "content/lib/fred.so")).To(Equal("fred.6.so"))
		})
	})
})
