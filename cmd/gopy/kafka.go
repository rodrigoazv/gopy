package gopy

import (
	"github.com/manifoldco/promptui"
)

func promptKafka() (wantKafka string, err error) {
	promptKafka := promptui.Select{
		Label: "You would like to use Kafka?",
		Items: []string{"Yes", "No"},
	}

	_, wantKafka, err = promptKafka.Run()
	return
}
