package services

import (
	"github.com/flowshot-io/x/pkg/logger"
	"github.com/flowshot-io/x/pkg/manager"
)

type (
	Option interface {
		apply(*options)
	}

	applyFunc func(*options)
)

func (f applyFunc) apply(s *options) { f(s) }

// WithServices indicates which supplied services (e.g. frontend, worker) within the server to start
func WithServices(names []manager.ServiceName) Option {
	return applyFunc(func(s *options) {
		s.serviceNames = make(map[manager.ServiceName]struct{})
		for _, name := range names {
			s.serviceNames[manager.ServiceName(name)] = struct{}{}
		}
	})
}

// InterruptOn interrupts server on the signal from server. If channel is nil Start() will block forever.
func InterruptOn(interruptCh <-chan interface{}) Option {
	return applyFunc(func(s *options) {
		s.blockingStart = true
		s.interruptCh = interruptCh
	})
}

// WithLogger sets a custom logger
func WithLogger(logger logger.Logger) Option {
	return applyFunc(func(s *options) {
		s.logger = logger
	})
}