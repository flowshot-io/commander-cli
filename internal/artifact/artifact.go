package artifact

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

type Artifact struct {
	Name      string
	Vfs       afero.Fs
	readIndex int
}

func New(artifactName string) *Artifact {
	if !strings.HasSuffix(artifactName, ".tar.gz") {
		artifactName = artifactName + ".tar.gz"
	}

	return &Artifact{
		Name:      artifactName,
		Vfs:       afero.NewMemMapFs(),
		readIndex: 0,
	}
}

func NewWithPaths(artifactName string, paths []string) (*Artifact, error) {
	artifact := New(artifactName)

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("error stating path: %s, error: %w", path, err)
		}

		if info.IsDir() {
			err = filepath.Walk(path, func(subPath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					content, err := os.ReadFile(subPath)
					if err != nil {
						return fmt.Errorf("error reading file: %s, error: %w", subPath, err)
					}
					artifact.AddFile(subPath, content)
				}
				return nil
			})
			if err != nil {
				return nil, fmt.Errorf("error walking directory: %s, error: %w", path, err)
			}
		} else {
			content, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("error reading file: %s, error: %w", path, err)
			}
			artifact.AddFile(path, content)
		}
	}

	return artifact, nil
}

func (a *Artifact) AddFile(filePath string, content []byte) error {
	file, err := a.Vfs.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	return err
}

func (a *Artifact) CreateTarGz() error {
	tarGzFile, err := os.Create(a.Name)
	if err != nil {
		return err
	}
	defer tarGzFile.Close()

	gzWriter := gzip.NewWriter(tarGzFile)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	err = afero.Walk(a.Vfs, "/", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(path)

		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}

		file, err := a.Vfs.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(tarWriter, file)
		return err
	})
	return err
}

func (a *Artifact) LoadFromTarGzFile(tarGzFilePath string) error {
	tarGzFile, err := os.Open(tarGzFilePath)
	if err != nil {
		return fmt.Errorf("error opening tar.gz file: %w", err)
	}
	defer tarGzFile.Close()

	err = a.LoadFromReader(tarGzFile)
	if err != nil {
		return fmt.Errorf("error loading artifact from tar.gz file: %w", err)
	}

	return nil
}

func (a *Artifact) SaveToDirectory(outputDir string) error {
	return afero.Walk(a.Vfs, "/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		outPath := filepath.Join(outputDir, path)
		if info.IsDir() {
			return os.MkdirAll(outPath, info.Mode())
		}

		inFile, err := a.Vfs.Open(path)
		if err != nil {
			return err
		}
		defer inFile.Close()

		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, inFile)
		return err
	})
}

// Implementing io.Writer interface
func (a *Artifact) Write(p []byte) (n int, err error) {
	err = a.LoadFromReader(bytes.NewReader(p))
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// Implementing io.Reader interface
func (a *Artifact) Read(p []byte) (n int, err error) {
	var buf bytes.Buffer
	err = a.SaveToWriter(&buf)
	if err != nil {
		return 0, err
	}
	data := buf.Bytes()
	bytesRead := copy(p, data[a.readIndex:])
	a.readIndex += bytesRead

	if bytesRead == 0 {
		return 0, io.EOF
	}
	return bytesRead, nil
}

func (a *Artifact) SaveToWriter(writer io.Writer) error {
	gzWriter := gzip.NewWriter(writer)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	err := afero.Walk(a.Vfs, "/", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(path)

		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}

		file, err := a.Vfs.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(tarWriter, file)
		return err
	})
	return err
}

func (a *Artifact) LoadFromReader(reader io.Reader) error {
	gzReader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			err = a.Vfs.MkdirAll(header.Name, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := a.Vfs.Create(header.Name)
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, tarReader)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *Artifact) Size() (int64, error) {
	var totalSize int64 = 0

	err := afero.Walk(a.Vfs, "/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			totalSize += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return totalSize, nil
}
