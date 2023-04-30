package command

import (
	"github.com/flowshot-io/commander-cli/internal/cli/command/blenderfarm"
	"github.com/flowshot-io/commander-cli/internal/cli/factory"
	"github.com/spf13/cobra"
)

type Driver struct {
	clientFactory factory.ClientFactory
}

func NewDriver() *Driver {
	clientFactory := factory.NewClientFactory()

	return &Driver{
		clientFactory: clientFactory,
	}
}

func NewCommand(d *Driver) *cobra.Command {
	c := &cobra.Command{
		Use:   "cli",
		Short: "A command-line tool for Commander.",
		Long:  `A command-line tool for performing tasks on a Commander cluster.`,
	}
	c.SetVersionTemplate("{{.Version}}\n")

	rdrv := blenderfarm.NewDriver(d.clientFactory)

	blenderfarmCMD := blenderfarm.NewRootCommand(rdrv)

	c.AddCommand(
		blenderfarmCMD,
	)

	return c
}
