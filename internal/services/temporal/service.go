package temporal

import (
	"github.com/flowshot-io/x/pkg/logger"
	"github.com/flowshot-io/x/pkg/manager"
)

type (
	Options struct {
		Logger logger.Logger
	}

	Service struct {
		logger logger.Logger
	}
)

func New(opts Options) (manager.Service, error) {
	if opts.Logger == nil {
		opts.Logger = logger.NoOp()
	}

	return &Service{
		logger: opts.Logger,
	}, nil
}

func (s *Service) Start() error {
	return nil
}

func (s *Service) Stop() error {
	return nil
}
