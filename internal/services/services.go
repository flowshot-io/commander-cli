package services

import (
	"fmt"

	"github.com/flowshot-io/x/pkg/manager"
)

const (
	CommanderService manager.ServiceName = "commander"
	TemporalService  manager.ServiceName = "temporalite"
	WayFarerService  manager.ServiceName = "wayfarer"
)

var Services = []string{
	string(CommanderService),
	string(TemporalService),
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
