package log

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefault(t *testing.T) {
	logger := SetDefault(nil)
	assert.NotNil(t, logger)

	logger = logrus.New()
	assert.Equal(t, logger, SetDefault(logger))
}

func TestNewWLLogger(t *testing.T) {
	l := NewWLLogger("ddd")
	l.Infof("dddd")
}
