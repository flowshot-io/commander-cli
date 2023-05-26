package servicecommand

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

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

			var wg sync.WaitGroup
			wg.Add(1)

			// Start the commander in a separate Goroutine
			go func() {
				defer wg.Done()
				err = services.Start()
				if err != nil {
					cmd.PrintErr(err)
				}
			}()

			// Listen for os.Interrupt signals (e.g., Ctrl+C)
			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

			<-signalChan
			cmd.Println("Received interrupt signal, stopping all services...")
			services.Stop()

			wg.Wait()

			return nil
		},
	}

	return c
}
