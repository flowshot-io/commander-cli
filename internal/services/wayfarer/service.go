package wayfarer

import "github.com/flowshot-io/x/pkg/manager"

type (
	Options struct {
	}

	Service struct {
	}
)

func New(opts Options) (manager.Service, error) {
	return &Service{}, nil
}

func (s *Service) Start() error {
	return nil
}

func (s *Service) Stop() error {
	return nil
}
