package factory

import (
	"fmt"

	"github.com/flowshot-io/commander-cli/internal/static"
	_ "github.com/flowshot-io/polystore/pkg/services/fs"
	_ "github.com/flowshot-io/polystore/pkg/services/s3"

	"github.com/flowshot-io/commander-client-go/client"
	"github.com/flowshot-io/commander-client-go/commanderservice/v1"
	"github.com/flowshot-io/polystore/pkg/services"
	"github.com/flowshot-io/x/pkg/artifactservice"
	"github.com/spf13/cobra"
)

type ClientFactory interface {
	CommanderClient(c *cobra.Command) (commanderservice.CommanderServiceClient, error)
	ArtifactClient(c *cobra.Command) (artifactservice.ArtifactServiceClient, error)
}

type clientFactory struct {
}

func NewClientFactory() ClientFactory {
	return &clientFactory{}
}

// CommanderClient returns a CommanderServiceClient
func (f *clientFactory) CommanderClient(c *cobra.Command) (commanderservice.CommanderServiceClient, error) {
	nopts := client.Options{Host: static.TemporalFrontEndHost}

	commander, err := client.Dial(nopts)
	if err != nil {
		return nil, fmt.Errorf("unable to create Navigator client: %w", err)
	}

	return commander, nil
}

// ArtifactClient returns a ArtifactServiceClient
func (f *clientFactory) ArtifactClient(c *cobra.Command) (artifactservice.ArtifactServiceClient, error) {
	store, err := services.NewStorageFromString(static.StorageConnection)
	if err != nil {
		return nil, fmt.Errorf("unable to create Storage client: %w", err)
	}

	service, err := artifactservice.New(artifactservice.Options{
		Store: store,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create Artifact client: %w", err)
	}

	return service, nil
}
