package trykafka

import (
	"context"
	"fmt"
	"time"

	//      "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(ctx context.Context) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})

	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("msg key and value is %s  %s\n", string(ev.Key), string(ev.Value))
					fmt.Printf("Delivered message (%s, %s) to partition%d and offset %d\n",
						string(ev.Key), string(ev.Value), ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	/*
	// Produce messages to topic (asynchronously)
	topic := "myTopic"
	i := 0
	for i < 6 {
		for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(word),
			}, nil)
		}
		time.Sleep(time.Second * 5)
		i++
	}
	*/

	// Wait for message deliveries before shutting down

	topicWithMulPart := "topic3"
	//key1 := "abc"
	//key2 := "pqr"
	//key3 := "xyz"
	//key4 := "clusterA"

	count := 0
	for count < 10 {


	p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{
		Topic:     &topicWithMulPart,
		Partition: kafka.PartitionAny,
	},
		Value: []byte("value1-" + string(count)),
		//Key:   []byte(key1),
		Timestamp:      time.Time{},
		TimestampType:  0,
		Opaque:         nil,
		Headers:        nil,
	}, nil)

	p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{
		Topic:     &topicWithMulPart,
		Partition: kafka.PartitionAny,
	},
		Value: []byte("value2-"+ string(count)),
		//Key:   []byte(key2),
		Timestamp:      time.Time{},
		TimestampType:  0,
		Opaque:         nil,
		Headers:        nil,
	}, nil)

	p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{
		Topic:     &topicWithMulPart,
		Partition: kafka.PartitionAny,
	},
		Value: []byte("value3-"+ string(count)),
		//Key:   []byte(key3),
		Timestamp:      time.Time{},
		TimestampType:  0,
		Opaque:         nil,
		Headers:        nil,
	}, nil)

	p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{
		Topic:     &topicWithMulPart,
		Partition: kafka.PartitionAny,
	},
		Value: []byte("value4-"+ string(count)),
		//Key:   []byte(key4),
		Timestamp:      time.Time{},
		TimestampType:  0,
		Opaque:         nil,
		Headers:        nil,
	}, nil)


	count++
	time.Sleep(2 * time.Second)
	}
	/*
	j := 0
	for j < 3 {
		keyfield := "key"+ string(j)
		valuedata := "value" + string(j)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topicWithMulPart,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(valuedata),
			Key:   []byte(keyfield),
			Timestamp:      time.Time{},
			TimestampType:  0,
			Opaque:         nil,
			Headers:        nil,
		},nil)
		time.Sleep(3*time.Second)
		j++
	}
	*/

	p.Flush(15 * 1000)


}

/*
import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strconv"
	"time"
)

func Produce(ctx context.Context) {
	// initialize a counter
	i := 0

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
	})

	for {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}
*/
