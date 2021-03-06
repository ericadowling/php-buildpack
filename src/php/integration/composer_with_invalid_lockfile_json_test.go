package integration_test

import (
	"os"
	"path/filepath"
	"time"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("When composer.lock is invalid JSON", func() {
	var app *cutlass.App
	AfterEach(func() { app = DestroyApp(app) })

	It("fails to stage", func() {
		app = cutlass.New(filepath.Join(bpDir, "fixtures", "composer_lock_invalid_json"))
		app.SetEnv("COMPOSER_GITHUB_OAUTH_TOKEN", os.Getenv("COMPOSER_GITHUB_OAUTH_TOKEN"))
		Expect(app.Push()).ToNot(Succeed())
		Expect(app.ConfirmBuildpack(buildpackVersion)).To(Succeed())

		Eventually(app.Stdout.String, 10*time.Second).Should(ContainSubstring("Invalid JSON present in composer.lock. Parser said"))
	})
})
