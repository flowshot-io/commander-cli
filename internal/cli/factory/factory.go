package factory

import (
	"fmt"

	"github.com/flowshot-io/commander-client-go/client"
	"github.com/flowshot-io/commander-client-go/commanderservice/v1"
	"github.com/spf13/cobra"
)

type ClientFactory interface {
	CommanderClient(c *cobra.Command) (commanderservice.CommanderServiceClient, error)
}

type clientFactory struct {
}

func NewClientFactory() ClientFactory {
	return &clientFactory{}
}

// CommanderClient returns a CommanderServiceClient
func (f *clientFactory) CommanderClient(c *cobra.Command) (commanderservice.CommanderServiceClient, error) {
	nopts := client.Options{Host: "localhost:50051"}

	commander, err := client.Dial(nopts)
	if err != nil {
		return nil, fmt.Errorf("unable to create Navigator client: %w", err)
	}

	return commander, nil
}
