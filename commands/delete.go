package commands

import (
	"fmt"

	"github.com/pivotal-cf/credhub-cli/actions"
	"github.com/pivotal-cf/credhub-cli/client"
	"github.com/pivotal-cf/credhub-cli/config"
	"github.com/pivotal-cf/credhub-cli/repositories"
)

type DeleteCommand struct {
	SecretIdentifier string `short:"n" long:"name" required:"yes" description:"Selects the secret to delete"`
}

func (cmd DeleteCommand) Execute([]string) error {
	cfg := config.ReadConfig()
	repository := repositories.NewSecretRepository(client.NewHttpClient(cfg))
	action := actions.NewAction(repository, cfg)

	_, err := action.DoAction(client.NewDeleteSecretRequest(cfg, cmd.SecretIdentifier), cmd.SecretIdentifier)

	if err == nil {
		fmt.Println("Secret successfully deleted")
	}

	return err
}
