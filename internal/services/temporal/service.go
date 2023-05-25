package temporal

import (
	"github.com/flowshot-io/x/pkg/logger"
	"github.com/flowshot-io/x/pkg/manager"
	"github.com/temporalio/cli/server"
)

type (
	Options struct {
		Logger logger.Logger
		DBPath string
	}

	Service struct {
		logger   logger.Logger
		temporal *server.Server
	}
)

func New(opts Options) (manager.Service, error) {
	if opts.DBPath == "" {
		opts.DBPath = "./temporal.db"
	}

	if opts.Logger == nil {
		opts.Logger = logger.NoOp()
	}

	topts := []server.ServerOption{
		server.WithDatabaseFilePath(opts.DBPath),
		server.WithLogger(NewTemporalLogger(opts.Logger)),
		server.WithNamespaces("commander"),
	}

	temporal, err := server.NewServer(topts...)
	if err != nil {
		return nil, err
	}

	return &Service{
		logger:   opts.Logger,
		temporal: temporal,
	}, nil
}

func (s *Service) Start() error {
	return s.temporal.Start()
}

func (s *Service) Stop() error {
	s.temporal.Stop()
	return nil
}
