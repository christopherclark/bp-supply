package integration_test

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create a supply buildpack", func() {
	var app *cutlass.App
	AfterEach(func() { app = DestroyApp(app) })

	Context("as a supply buildpack", func() {
		BeforeEach(func() {
			app = cutlass.New(filepath.Join(bpDir, "fixtures", "app"))
			app.Buildpacks = []string{simpleBuildpack, "binary_buildpack"}
			app.SetEnv("BP_DEBUG", "1")
		})

		It("supplies fixtures/simple files to later buildpacks", func() {
			PushApp(app)

			Expect(app.Stdout.String()).To(ContainSubstring("-----> Simple Buildpack version 1.2.3"))
			Expect(app.GetBody("/")).To(ContainSubstring("Jane's Contents"))

			By("testing the config.yml for our supply buildpack", func() {
				Expect(app.GetBody("/supply_config")).To(MatchYAML(`{"name": "Simple", "config":{}}`))
			})
		})
	})
})
