package configtest

import (
	"github.com/ezeslucky/monitrix/pkg/config"
	"github.com/ezeslucky/monitrix/pkg/config/envprovider"
)

func NewResolverConfig() config.ResolverConfig {
	return config.ResolverConfig{
		Uris:              []string{"env:"},
		ProviderFactories: []config.ProviderFactory{envprovider.NewFactory()},
	}
}
