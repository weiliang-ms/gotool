package log

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func SetDefault(logger *logrus.Logger) *logrus.Logger {
	if logger == nil {
		return logrus.New()
	}
	return logger
}

func NewWLLogger(banner string) *logrus.Logger {
	logger := logrus.New()
	fieldMap := FieldMap{}
	fieldMap["banner"] = banner
	formatter := &CustomFormatter{
		EnvironmentOverrideColors: true, TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: fieldMap,
	}

	//formatter.Format(entry)
	logger.SetFormatter(formatter)
	logger.SetLevel(DebugLevel)
	return logger
}
