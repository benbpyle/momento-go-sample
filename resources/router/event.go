package main

type SampleEvent struct {
	Name          string `json:"name"`
	CorrelationId string `json:"correlationId"`
}
