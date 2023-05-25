package services

import (
	"fmt"

	"github.com/flowshot-io/x/pkg/logger"
	"github.com/flowshot-io/x/pkg/manager"
	"golang.org/x/exp/slices"
)

type options struct {
	serviceNames map[manager.ServiceName]struct{}

	interruptCh   <-chan interface{}
	blockingStart bool

	logger logger.Logger
}

func newOptions(opts []Option) *options {
	so := &options{
		logger: logger.New(&logger.Options{Pretty: true}),
	}
	for _, opt := range opts {
		opt.apply(so)
	}

	return so
}

func (so *options) loadAndValidate() error {
	for serviceName := range so.serviceNames {
		if !slices.Contains(Services, serviceName) {
			return fmt.Errorf("invalid service %q in service list %v", serviceName, so.serviceNames)
		}
	}

	return nil
}
