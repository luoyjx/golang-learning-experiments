package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	p         *kafka.Producer
	closeChan chan struct{}
}

func NewProducer() *Producer {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
		},
	)

	if err != nil {
		panic(err)
	}

	closeChan := make(chan struct{}, 1)
	go func() {
		defer close(closeChan)
		for {
			select {
			case msg := <-p.Events():
				switch ev := msg.(type) {
				case *kafka.Message:
					m := ev
					if m.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
					} else {
						fmt.Printf(
							"Delivered message to topic %s [%d] at offset %v\n",
							*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset,
						)
					}
				default:
					// fmt.Printf("Ignored event: %s\n", ev)
				}
			case <-closeChan:
				fmt.Println("closing ...")
				break
			}
		}
	}()

	return &Producer{
		p:         p,
		closeChan: closeChan,
	}
}

func (p *Producer) Send(data []byte) error {
	fmt.Println("sending message ", string(data))
	topic := "test"

	p.p.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: data,
	}

	fmt.Println("sended ", string(data))

	return nil
}

func (p *Producer) Close() {
	p.p.Close()
	p.closeChan <- struct{}{}
}

func main() {
	p := NewProducer()

	p.Send([]byte("1"))
	p.Send([]byte("2"))
	p.Send([]byte("3"))
	p.Send([]byte("4"))

	time.Sleep(time.Second * 10)
	p.Close()
	time.Sleep(time.Second * 2)
}
