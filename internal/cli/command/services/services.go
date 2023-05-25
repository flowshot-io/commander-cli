package servicescommand

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
		Use:   "services",
		Short: "Services commands",
		Long:  `Services commands`,
	}

	runCMD := NewRunCommand(d)

	c.AddCommand(
		runCMD,
	)

	return c
}
