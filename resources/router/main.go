package main

import (
	"context"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/momentohq/client-sdk-go/momento"
	log "github.com/sirupsen/logrus"
)

var (
	routeRepository RouterRepository
	tableName       string
	cacheClient     momento.CacheClient
	cacheName       string
)

func handler(ctx context.Context, e SampleEvent) error {
	log.WithFields(log.Fields{
		"event": e,
	}).Debug("Printing out the event")

	route := determineRoute(ctx, e)
	if route != nil {
		log.WithFields(log.Fields{
			"route": route,
		}).Debug("Printing out the route")
	}

	return nil
}

func init() {
	isLocal, _ := strconv.ParseBool(os.Getenv("IS_LOCAL"))
	token, err := GetSecretString()

	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: isLocal,
	})

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Fetching token failed, now I have to go away")
	}

	cacheClient, err = NewMomentoClient(*token)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Creating cache client, now I have to go away")
	}

	dbClient := NewDynamoDBClient(isLocal)
	routeRepository = &RouterDynamoRepository{db: dbClient}
	tableName = os.Getenv("TABLE_NAME")
	cacheName = os.Getenv("CACHE_NAME")
	SetLevel(os.Getenv("LOG_LEVEL"))
}

func main() {
	lambda.Start(handler)
}
