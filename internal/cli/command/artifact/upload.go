package artifactcommand

import (
	"fmt"

	"github.com/flowshot-io/commander-cli/internal/artifact"
	"github.com/spf13/cobra"
)

func NewUploadCommand(d *Driver) *cobra.Command {
	var artifactName string
	var tarGzFile string
	var paths []string

	c := &cobra.Command{
		Use:   "upload",
		Short: "Upload an artifact to storage",
		Long:  `Upload an artifact to storage.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if tarGzFile != "" && len(paths) > 0 {
				return fmt.Errorf("cannot specify both tar.gz file and paths")
			}

			var a *artifact.Artifact
			var err error

			if tarGzFile != "" {
				a = artifact.New(artifactName)
				err = a.LoadFromTarGzFile(tarGzFile)
				if err != nil {
					return fmt.Errorf("error loading artifact from tar.gz file: %w", err)
				}
			} else {
				a, err = artifact.NewWithPaths(artifactName, paths)
				if err != nil {
					return fmt.Errorf("error creating artifact with paths: %w", err)
				}
			}

			fmt.Printf("Artifact '%s' loaded successfully\n", artifactName)

			client, err := d.clientFactory.ArtifactClient(cmd)
			if err != nil {
				return fmt.Errorf("error creating artifact client: %w", err)
			}

			err = client.UploadArtifact(cmd.Context(), a)
			if err != nil {
				return fmt.Errorf("error uploading artifact: %w", err)
			}

			fmt.Printf("Artifact '%s' uploaded successfully\n", artifactName)

			return nil
		},
	}

	c.Flags().StringVarP(&artifactName, "artifact", "a", "", "Artifact to upload (required)")
	c.MarkFlagRequired("artifact")
	c.Flags().StringVarP(&tarGzFile, "file", "f", "", "Path to the tar.gz file")
	c.Flags().StringSliceVarP(&paths, "paths", "p", []string{}, "List of paths or directories to create an artifact from")

	return c
}
