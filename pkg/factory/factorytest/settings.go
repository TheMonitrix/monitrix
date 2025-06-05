package factorytest

import (
	"github.com/ezeslucky/monitrix/pkg/factory"
	"github.com/ezeslucky/monitrix/pkg/instrumentation/instrumentationtest"
)

func NewSettings() factory.ProviderSettings {
	return instrumentationtest.New().ToProviderSettings()
}
