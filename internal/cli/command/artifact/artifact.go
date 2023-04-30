package artifactcommand

import (
	"github.com/flowshot-io/commander-cli/internal/cli/factory"
	"github.com/spf13/cobra"
)

type Driver struct {
	clientFactory factory.ClientFactory
}

func NewDriver(clientFactory factory.ClientFactory) *Driver {
	return &Driver{
		clientFactory: clientFactory,
	}
}

func NewRootCommand(d *Driver) *cobra.Command {
	c := &cobra.Command{
		Use:   "artifact",
		Short: "Artifact commands",
		Long:  `Artifact commands`,
	}

	uploadCMD := NewUploadCommand(d)

	c.AddCommand(
		uploadCMD,
	)

	return c
}
