package main

import (
	"fmt"
	"github.com/cloudfoundry/switchblade"
	"os"
	"path/filepath"
)

func main() {

	token := os.Getenv("GITHUB_TOKEN")
	platform, err := switchblade.NewPlatform(switchblade.Docker, token, "cflinuxfs3")
	if err != nil {
		panic(err)
	}

	currentDir, err := os.Getwd()

	_, logs, err := platform.Deploy.
		WithBuildpacks("python_buildpack").
		WithEnv(map[string]string{"APPD_AGENT": ""}).
		WithServices(map[string]switchblade.Service{
			"appdynamics-service": {
				"account-access-key": "test-key", "account-name": "test-account", "host-name": "test-ups-host", "port": "1234", "ssl-enabled": "true",
			},
		}).
		Execute("app", filepath.Join(currentDir, "app"))

	if err != nil {
		panic(err)
	}

	fmt.Println(logs.String())

}
