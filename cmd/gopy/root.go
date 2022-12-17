package gopy

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

func Execute() {
	nameOfProject, err := promptName()
	wantEnvs, err := promptEnv()
	wantKafka, err := promptKafka()

	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond) // Build our new spinner
	s.Start()

	defer s.Stop()

	cmdClone := exec.Command("git", "clone", "git@github.com:solfacil/python-microservice-template.git")

	err = cmdClone.Run()
	if err != nil {
		// something went wrong
	}

	err = os.Rename("python-microservice-template", nameOfProject)
	if err != nil {
		log.Fatal(err)
	}

	err = os.RemoveAll("./" + nameOfProject + "/.git")
	if err != nil {
		log.Fatal(err)
	}

	if wantKafka != "Yes" {
		src := nameOfProject + "/app/internal/kafka"
		e := os.RemoveAll(src)
		if e != nil {
			log.Fatal(e)
			os.RemoveAll(nameOfProject)
		}
	}
	if wantEnvs == "Yes" {
		src := nameOfProject + "/.env-sample"
		dest := nameOfProject + "/.env"

		bytesRead, err := ioutil.ReadFile(src)

		if err != nil {
			log.Fatal(err)
			os.RemoveAll(nameOfProject)
		}

		err = ioutil.WriteFile(dest, bytesRead, 0644)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("\n\nAll Done! Kafka?:%q Envs?:%q\n\n", wantKafka, wantEnvs)
	fmt.Printf("type - cd %q - and start coding", nameOfProject)

}
