package gopy

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func promptName() (nameOfProject string, err error) {
	validate := func(input string) error {
		var minimalInput = 8
		if len(input) < minimalInput {
			return errors.New("Very short")
		}
		return nil
	}

	promptNameOfProject := promptui.Prompt{
		Label:    "Change name of your project",
		Validate: validate,
	}

	nameOfProject, err = promptNameOfProject.Run()
	return
}
