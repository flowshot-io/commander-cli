package artifactcommand

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func NewDownloadCommand(d *Driver) *cobra.Command {
	var artifactName string
	var destinationPath string

	c := &cobra.Command{
		Use:   "download",
		Short: "download an artifact from storage",
		Long:  `download an artifact from storage.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := d.clientFactory.ArtifactClient(cmd)
			if err != nil {
				return fmt.Errorf("error creating artifact client: %w", err)
			}

			artifact, err := client.DownloadArtifact(cmd.Context(), artifactName)
			if err != nil {
				return fmt.Errorf("error getting artifact: %w", err)
			}

			// fmt.Printf("Artifact '%s' downloaded successfully\n", artifact.Name)

			extractionPath, err := makeAbsolutePath(filepath.Join(destinationPath, artifactName))
			if err != nil {
				return fmt.Errorf("error making destination path absolute: %w", err)
			}

			err = artifact.ExtractToDirectory(extractionPath)
			if err != nil {
				return fmt.Errorf("error extracting artifact file: %w", err)
			}

			fmt.Printf("Artifact '%s' extracted successfully to '%s'\n", artifact.GetName(), extractionPath)

			return nil
		},
	}

	c.Flags().StringVarP(&artifactName, "name", "n", "", "Name for artifact to upload (required)")
	c.MarkFlagRequired("name")

	c.Flags().StringVarP(&destinationPath, "destination", "d", ".", "Destination path")

	return c
}

func makeAbsolutePath(relPath string) (string, error) {
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		return "", err
	}
	return absPath, nil
}
