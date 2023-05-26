package commander

import (
	"github.com/flowshot-io/commander-cli/internal/static"
	"github.com/flowshot-io/commander/pkg/commander"
	"github.com/flowshot-io/commander/pkg/commander/config"
	"github.com/flowshot-io/x/pkg/logger"
	"github.com/flowshot-io/x/pkg/manager"
)

type (
	Options struct {
		Logger logger.Logger
	}

	Service struct {
		logger    logger.Logger
		commander manager.Service
	}
)

func New(opts Options) (manager.Service, error) {
	if opts.Logger == nil {
		opts.Logger = logger.NoOp()
	}

	config := config.Config{
		Global: config.Global{
			Temporal: config.Temporal{
				Host: static.TemporalWorkerHost,
			},
			Storage: config.Storage{
				ConnectionString: static.StorageConnection,
			},
		},
	}

	commander, err := commander.New(
		commander.WithLogger(opts.Logger),
		commander.WithConfig(&config),
		commander.WithServices([]string{"frontend", "blenderfarm", "blendernode"}),
	)
	if err != nil {
		return nil, err
	}

	return &Service{
		logger:    opts.Logger,
		commander: commander,
	}, nil
}

func (s *Service) Start() error {
	return s.commander.Start()
}

func (s *Service) Stop() error {
	return s.commander.Stop()
}
