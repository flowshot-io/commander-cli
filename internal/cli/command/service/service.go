package servicecommand

import (
	"github.com/spf13/cobra"
)

type Driver struct {
}

func NewDriver() *Driver {
	return &Driver{}
}

func NewRootCommand(d *Driver) *cobra.Command {
	c := &cobra.Command{
		Use:   "service",
		Short: "Service commands",
		Long:  `Service commands`,
	}

	runCMD := NewRunCommand(d)

	c.AddCommand(
		runCMD,
	)

	return c
}
