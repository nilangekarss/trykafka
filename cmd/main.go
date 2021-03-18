package main

import (
	"context"
	"github.com/nilangekarss/trykafka"
)


func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	trykafka.Produce(ctx)
	//go trykafka.Consume(ctx)
	//trykafka.CreateTopic(ctx)
	//time.Sleep(5*time.Second)
	//trykafka.DescribeTopic(ctx)

}

