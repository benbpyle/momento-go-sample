package main

import (
	"encoding/json"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

func GetSecretString() (*string, error) {
	cache, err := secretcache.New()
	if err != nil {
		return nil, err
	}

	secretString, err := cache.GetSecretString("mo-cache-token")

	if err != nil {
		return nil, err
	}

	ss := struct {
		Token string `json:"token"`
	}{}

	err = json.Unmarshal([]byte(secretString), &ss)

	if err != nil {
		return nil, err
	}

	return &ss.Token, nil
}
