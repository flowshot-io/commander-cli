package servicescommand

import (
	"github.com/flowshot-io/commander-cli/internal/services"
	"github.com/flowshot-io/x/pkg/manager"
	"github.com/spf13/cobra"
)

func NewRunCommand(d *Driver) *cobra.Command {
	c := &cobra.Command{
		Use:   "run",
		Short: "Run a service",
		Long:  `Run a service`,
		RunE: func(cmd *cobra.Command, args []string) error {
			services, err := services.New(
				services.WithServices([]manager.ServiceName{
					services.CommanderService,
					services.TemporalService,
					services.WayFarerService,
				}),
			)
			if err != nil {
				return err
			}

			return services.Start()
		},
	}

	return c
}
