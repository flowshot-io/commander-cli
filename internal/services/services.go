package services

import (
	"fmt"

	"github.com/flowshot-io/commander-cli/internal/services/commander"
	"github.com/flowshot-io/commander-cli/internal/services/temporal"
	"github.com/flowshot-io/commander-cli/internal/services/wayfarer"
	"github.com/flowshot-io/x/pkg/manager"
)

const (
	CommanderService manager.ServiceName = "commander"
	TemporalService  manager.ServiceName = "temporalite"
	WayFarerService  manager.ServiceName = "wayfarer"
)

var Services = []manager.ServiceName{
	CommanderService,
	TemporalService,
	WayFarerService,
}

type Manager struct {
	services manager.ServiceController
}

func New(opts ...Option) (manager.Service, error) {
	so, err := Options(opts)
	if err != nil {
		return nil, err
	}

	srvs := manager.New(&manager.Options{Logger: so.logger})
	if err != nil {
		return nil, fmt.Errorf("unable to create service manager: %w", err)
	}

	if _, ok := so.serviceNames[CommanderService]; ok {
		srv, err := commander.New(commander.Options{})
		if err != nil {
			return nil, fmt.Errorf("unable to create commander service: %w", err)
		}

		srvs.Add(CommanderService, srv)
	}

	if _, ok := so.serviceNames[TemporalService]; ok {
		srv, err := temporal.New(temporal.Options{})
		if err != nil {
			return nil, fmt.Errorf("unable to create temporal service: %w", err)
		}

		srvs.Add(TemporalService, srv)
	}

	if _, ok := so.serviceNames[WayFarerService]; ok {
		srv, err := wayfarer.New(wayfarer.Options{})
		if err != nil {
			return nil, fmt.Errorf("unable to create wayfarer service: %w", err)
		}

		srvs.Add(WayFarerService, srv)
	}

	return &Manager{services: srvs}, nil
}

func Options(opts []Option) (*options, error) {
	so := newOptions(opts)

	err := so.loadAndValidate()
	if err != nil {
		return nil, err
	}

	return so, nil
}

func (m *Manager) Start() error {
	return m.services.Start()
}

func (m *Manager) Stop() error {
	return m.services.Stop()
}
