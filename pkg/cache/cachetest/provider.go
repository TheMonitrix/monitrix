package cachetest

import (
	"context"

	"github.com/ezeslucky/monitrix/pkg/cache"
	"github.com/ezeslucky/monitrix/pkg/cache/memorycache"
	"github.com/ezeslucky/monitrix/pkg/factory/factorytest"
)

type provider struct{}

func New(config cache.Config) (cache.Cache, error) {
	cache, err := memorycache.New(context.TODO(), factorytest.NewSettings(), config)
	if err != nil {
		return nil, err
	}

	return cache, nil
}
