package blenderfarm

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
		Use:   "blenderfarm",
		Short: "Blenderfarm workflow commands",
		Long:  `Blenderfarm workflow commands`,
	}

	startCMD := NewStartCommand(d)

	c.AddCommand(
		startCMD,
	)

	return c
}
