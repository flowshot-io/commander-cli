package factory

import (
	"fmt"

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
	nopts := client.Options{Host: "localhost:50051"}

	commander, err := client.Dial(nopts)
	if err != nil {
		return nil, fmt.Errorf("unable to create Navigator client: %w", err)
	}

	return commander, nil
}

// ArtifactClient returns a ArtifactServiceClient
func (f *clientFactory) ArtifactClient(c *cobra.Command) (artifactservice.ArtifactServiceClient, error) {
	connectionString := "s3://commander/workflows/?accessKey=5kpWVH8bjA3ak8Kv&secretKey=ipvdKs21pyp3aFmKwNbU9iAJJTkH3c9Q&endpoint=http://localhost:9099"

	store, err := services.NewStorageFromString(connectionString)
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
