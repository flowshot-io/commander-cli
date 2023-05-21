package command

import (
	artifactcommand "github.com/flowshot-io/commander-cli/internal/cli/command/artifact"
	blenderfarmcommand "github.com/flowshot-io/commander-cli/internal/cli/command/blenderfarm"
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
		Use:          "commander",
		Short:        "A command-line tool for Commander.",
		Long:         `A command-line tool for performing tasks on a Commander cluster.`,
		SilenceUsage: true,
	}
	c.SetVersionTemplate("{{.Version}}\n")

	rdrv := blenderfarmcommand.NewDriver(d.clientFactory)
	adrv := artifactcommand.NewDriver(d.clientFactory)

	blenderfarmCMD := blenderfarmcommand.NewRootCommand(rdrv)
	artifactCMD := artifactcommand.NewRootCommand(adrv)

	c.AddCommand(
		blenderfarmCMD,
		artifactCMD,
	)

	return c
}
