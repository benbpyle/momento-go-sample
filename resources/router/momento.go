package main

import (
	"context"
	"time"

	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
	"github.com/momentohq/client-sdk-go/responses"
	log "github.com/sirupsen/logrus"
)

func NewMomentoClient(token string) (momento.CacheClient, error) {
	// Initializes Momento
	credentialProvider, err := auth.FromString(token)

	if err != nil {
		return nil, err
	}

	// Initializes Momento
	client, err := momento.NewCacheClient(
		config.InRegionLatest(),
		credentialProvider,
		600*time.Second)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func SetRoute(ctx context.Context, route *Route) error {
	v := momento.String(route.QueueUrl)
	request := momento.SetRequest{
		CacheName: cacheName,
		Key:       momento.String(route.Key),
		Value:     v,
	}

	_, err := cacheClient.Set(ctx, &request)

	if err == nil {
		log.WithFields(log.Fields{
			"key": route.Key,
		}).Info("Cache set")
	}

	return err
}

func ReadRoute(ctx context.Context, key string) (*Route, error) {
	request := momento.GetRequest{
		CacheName: cacheName,
		Key:       momento.String(key),
	}

	resp, err := cacheClient.Get(ctx, &request)

	if err != nil {
		return nil, err
	}

	if v, ok := resp.(*responses.GetHit); ok {
		log.WithFields(log.Fields{
			"key": key,
		}).Info("Cache hit")
		return &Route{
			QueueUrl: v.ValueString(),
		}, nil
	}

	return nil, nil
}
