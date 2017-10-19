package integration_test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var bpDir string

func init() {
	flag.StringVar(&cutlass.DefaultMemory, "memory", "128M", "default memory for pushed apps")
	flag.StringVar(&cutlass.DefaultDisk, "disk", "256M", "default disk for pushed apps")
	flag.Parse()
}

var _ = SynchronizedBeforeSuite(func() []byte {
	// Run once
	var err error
	bpDir, err = cutlass.FindRoot()
	Expect(err).NotTo(HaveOccurred())

	cmd := exec.Command("go", "build", "-o", "bp-supply", "main.go")
	cmd.Dir = bpDir
	Expect(cmd.Run()).To(Succeed())

	cmd = exec.Command("cf", "install-plugin", "-f", "bp-supply")
	cmd.Dir = bpDir
	Expect(cmd.Run()).To(Succeed())

	return nil
}, func(_ []byte) {
	// Run on all nodes
	var err error
	bpDir, err = cutlass.FindRoot()
	Expect(err).NotTo(HaveOccurred())

	cutlass.SeedRandom()
	cutlass.DefaultStdoutStderr = GinkgoWriter
})

var _ = SynchronizedAfterSuite(func() {
	// Run on all nodes
}, func() {
	// Run once
	Expect(cutlass.DeleteOrphanedRoutes()).To(Succeed())
})

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

func PushApp(app *cutlass.App) {
	Expect(app.Push()).To(Succeed())
	Eventually(func() ([]string, error) { return app.InstanceStates() }, 20*time.Second).Should(Equal([]string{"RUNNING"}))
}

func DestroyApp(app *cutlass.App) *cutlass.App {
	if app != nil {
		app.Destroy()
	}
	return nil
}

func CreateSupplyBuildpack(fixture, bpName, bpVersion string) string {
	tmpfile, err := ioutil.TempFile("", "bp-supply-")
	Expect(err).NotTo(HaveOccurred())
	Expect(tmpfile.Close()).To(Succeed())

	simpleBuildpack := "bp-supply-" + fixture + "-" + cutlass.RandStringRunes(6)
	simpleBuildpackFile := tmpfile.Name() + ".zip"

	cmd := exec.Command("cf", "bp-supply", fmt.Sprintf("--path=fixtures/%s", fixture), fmt.Sprintf("--version=%s", bpVersion), bpName, simpleBuildpackFile)
	cmd.Dir = bpDir
	Expect(cmd.Run()).To(Succeed())

	Expect(cutlass.CreateOrUpdateBuildpack(simpleBuildpack, simpleBuildpackFile)).To(Succeed())

	Expect(os.Remove(simpleBuildpackFile)).To(Succeed())

	return simpleBuildpack + "_buildpack"
}

func RemoveBuildpack(buildpack string) {
	Expect(exec.Command("cf", "delete-buildpack", "-f", buildpack).Run()).To(Succeed())
}
