package gopy

import (
	"github.com/manifoldco/promptui"
)

func promptEnv() (wantEnvs string, err error) {
	promptEnvs := promptui.Select{
		Label: "You would like to generate the enviroments?",
		Items: []string{"Yes", "No"},
	}

	_, wantEnvs, err = promptEnvs.Run()
	return
}
