package integration_test

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var bpDir, simpleBuildpack, simpleBuildpackFile string

type passToRunAll struct {
	BpDir               string
	SimpleBuildpack     string
	SimpleBuildpackFile string
}

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

	tmpfile, err := ioutil.TempFile("", "simple-bp-")
	Expect(err).NotTo(HaveOccurred())
	Expect(tmpfile.Close()).To(Succeed())
	data := passToRunAll{
		BpDir:               bpDir,
		SimpleBuildpack:     "simple_bp_" + cutlass.RandStringRunes(6),
		SimpleBuildpackFile: tmpfile.Name() + ".zip",
	}

	cmd := exec.Command("go", "build", "-o", "bp-supply", "main.go")
	cmd.Dir = bpDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	Expect(cmd.Run()).To(Succeed())

	cmd := exec.Command("cf", "install-plugin", "-f", "bp-supply")
	cmd.Dir = bpDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	Expect(cmd.Run()).To(Succeed())

	cmd := exec.Command("cf", "bp-supply", "--path=fixtures/simple", "--version=1.2.3", "Simple", data.SimpleBuildpackFile)
	cmd.Dir = bpDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	Expect(cmd.Run()).To(Succeed())

	Expect(cutlass.CreateOrUpdateBuildpack(data.SimpleBuildpack, data.SimpleBuildpackFile)).To(Succeed())
	data.SimpleBuildpack = data.SimpleBuildpack + "_buildpack"

	dataBytes, err := json.Marshal(data)
	Expect(err).NotTo(HaveOccurred())
	return dataBytes
}, func(dataBytes []byte) {
	// Run on all nodes
	data := passToRunAll{}
	err := json.Unmarshal(dataBytes, &data)
	Expect(err).NotTo(HaveOccurred())

	bpDir = data.BpDir
	simpleBuildpack = data.SimpleBuildpack
	simpleBuildpackFile = data.SimpleBuildpackFile

	cutlass.SeedRandom()
	cutlass.DefaultStdoutStderr = GinkgoWriter
})

var _ = SynchronizedAfterSuite(func() {
	// Run on all nodes
}, func() {
	// Run once
	RemoveBuildpack(simpleBuildpack)
	Expect(os.RemoveAll(simpleBuildpackFile)).To(Succeed())
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

func RemoveBuildpack(buildpack string) {
	Expect(exec.Command("cf", "delete-buildpack", "-f", buildpack).Run()).To(Succeed())
}
