package log

import "go.uber.org/zap"

var Instance *zap.Logger

func init() {
	var err error

	Instance, err = zap.NewProduction()
	if err != nil {
		Instance = zap.NewNop()
	}
}
