package blenderfarmcommand

import (
	"fmt"

	"github.com/flowshot-io/commander-client-go/commanderservice/v1"
	"github.com/spf13/cobra"
)

func NewStartCommand(d *Driver) *cobra.Command {
	var artifactName string
	var startFrame int32
	var endFrame int32
	var batchSize int32

	c := &cobra.Command{
		Use:   "start",
		Short: "Start workflow",
		Long:  `Start workflow`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Starting workflow")

			client, err := d.clientFactory.CommanderClient(cmd)
			if err != nil {
				return fmt.Errorf("error creating commander client: %w", err)
			}

			resp, err := client.CreateBlenderFarmWorkflow(cmd.Context(), &commanderservice.CreateBlenderFarmWorkflowRequest{
				File:       artifactName,
				StartFrame: startFrame,
				EndFrame:   endFrame,
				BatchSize:  batchSize,
			})
			if err != nil {
				return fmt.Errorf("error starting workflow: %w", err)
			}

			cmd.Println("Workflow started", resp)

			return nil
		},
	}

	c.Flags().StringVarP(&artifactName, "artifact", "a", "", "Artifact name (required)")
	c.Flags().Int32VarP(&startFrame, "start", "s", 0, "Start frame")
	c.Flags().Int32VarP(&endFrame, "end", "e", 0, "End frame")
	c.Flags().Int32VarP(&batchSize, "batch", "b", 0, "Batch size")

	c.MarkFlagRequired("artifact")

	return c
}
