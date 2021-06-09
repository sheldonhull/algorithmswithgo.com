package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/goyek/goyek"
	"github.com/mattn/go-shellwords"
)

const (
	buildDir    = "build"
	artifactDir = "./artifacts"
)

var ()

func main() {
	if err := os.Chdir(".."); err != nil {
		log.Fatalln(err)
	}
	flow().Main()
}

func flow() *goyek.Taskflow {
	flow := &goyek.Taskflow{}
	installtools := flow.Register(TaskInstallTools())
	initdir := flow.Register(TaskInitDir())
	flow.Register(goyek.Task{
		Name:  "init",
		Usage: "setup project for use",
		Deps:  goyek.Deps{initdir, installtools},
	})
	// flow.DefaultTask = all
	return flow
}

// Exec runs the provided command line.
// It fails the task in case of any problems.
func Exec(tf *goyek.TF, workDir string, cmdLine string) {
	args, err := shellwords.Parse(cmdLine)
	if err != nil {
		tf.Fatalf("parse command line: %v", err)
	}
	cmd := tf.Cmd(args[0], args[1:]...)
	cmd.Dir = workDir
	if err := cmd.Run(); err != nil {
		tf.Fatalf("run command: %v", err)
	}
}

func TaskInitDir() goyek.Task {
	return goyek.Task{
		Name: "initdir",
		// Usage: "create project directories",
		Command: func(tf *goyek.TF) {
			// pterm.Success.Printf("directories: %v\n", artifactDirs)
			for _, i := range []string{artifactDir} {
				d, err := filepath.Abs(i)
				if err != nil {
					tf.Errorf("abs path: [%v]", i)
				}

				if err := os.Mkdir(d, 0700); err != nil {
					if strings.Contains(err.Error(), "dir exists") {
						// pterm.Success.Printf("🔃 [%s] dir already exists\n", i)
					} else {
						tf.Errorf("failed to mkdir: [%s] with error: %v", d, err)
					}
					continue
				}
				// pterm.Success.Printf("✅ [%s] dir created\n", i)
			}
		},
	}
}

func TaskInstallTools() goyek.Task {
	return goyek.Task{
		Name:  "install-tools",
		Usage: "initialize dependencies and tooling for project",
		Command: func(tf *goyek.TF) {
			Exec(tf, "", "go install github.com/gruntwork-io/fetch")
			tf.Log("✅ installed fetch cli")
			Exec(tf, "", "fetch --repo=\"https://github.com/git-town/git-town\" --release-asset=\"*linux_arm_64.tar.gz\" \""+artifactDir+"\" --progress")
			tf.Log("✅ installed git-town")
		},
	}
}
