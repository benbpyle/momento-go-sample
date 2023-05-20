package main

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type Route struct {
	Key      string `dynamodbav:"pk"`
	QueueUrl string `dynamodbav:"QueueUrl"`
}

func determineRoute(ctx context.Context, e SampleEvent) *Route {
	r, err := ReadRoute(ctx, e.Name)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Error fetching route from cache")
		return nil
	}

	if r != nil {
		log.Debugf("Route was in Cache")
		return r
	}

	route, err := routeRepository.GetRoute(ctx, e.Name)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Error fetching route from DDB")
		return nil
	}

	if route != nil {
		log.Info("Setting the Cache")
		err = SetRoute(ctx, route)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Error setting route from cache")
		}
	}

	return route
}
