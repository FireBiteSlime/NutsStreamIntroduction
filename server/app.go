package main

import (
	"fmt"
	"math/rand"
	"nstream/deserializer"
	"nstream/randomizer"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sc, _ := stan.Connect("test-cluster", "script", stan.NatsURL("nats://localhost:4223"))
	defer sc.Close()
	for i := 1; ; i++ {
		order := randomizer.RandomOrder()
		bOrder, err := deserializer.MarshalJson(order)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bOrder))
		sc.Publish("foo", bOrder)
		time.Sleep(3 * time.Second)
	}

}
