package temporal

import (
	"github.com/flowshot-io/x/pkg/logger"
	"go.temporal.io/server/common/log/tag"
)

type TemporalTagLoggerAdapter struct {
	logger logger.Logger
}

func NewTemporalLogger(logger logger.Logger) *TemporalTagLoggerAdapter {
	return &TemporalTagLoggerAdapter{
		logger: logger,
	}
}

func (t *TemporalTagLoggerAdapter) Debug(msg string, tags ...tag.Tag) {
	t.logger.Debug(msg, tagsToFields(tags...))
}

func (t *TemporalTagLoggerAdapter) Info(msg string, tags ...tag.Tag) {
	t.logger.Info(msg, tagsToFields(tags...))
}

func (t *TemporalTagLoggerAdapter) Warn(msg string, tags ...tag.Tag) {
	t.logger.Warn(msg, tagsToFields(tags...))
}

func (t *TemporalTagLoggerAdapter) Error(msg string, tags ...tag.Tag) {
	t.logger.Error(msg, tagsToFields(tags...))
}

func (t *TemporalTagLoggerAdapter) DPanic(msg string, tags ...tag.Tag) {
	t.logger.Error("DPanic: "+msg, tagsToFields(tags...))
}

func (t *TemporalTagLoggerAdapter) Panic(msg string, tags ...tag.Tag) {
	t.logger.Error("Panic: "+msg, tagsToFields(tags...))
}

func (t *TemporalTagLoggerAdapter) Fatal(msg string, tags ...tag.Tag) {
	t.logger.Error("Fatal: "+msg, tagsToFields(tags...))
}

// Convert tags to a map[string]interface{}
func tagsToFields(tags ...tag.Tag) map[string]interface{} {
	fields := make(map[string]interface{})
	for _, t := range tags {
		fields[t.Key()] = t.Value()
	}
	return fields
}
