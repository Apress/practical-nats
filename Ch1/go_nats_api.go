package main


import (
        "fmt"
        "github.com/nats-io/go-nats"
)

func main(){
        nc, _ := nats.Connect(nats.DefaultURL)

        // Simple Publisher
        nc.Publish("foo", []byte("Hello World"))

        // Simple Async Subscriber
        nc.Subscribe("foo", func(m *nats.Msg) {
                fmt.Printf("Received a message: %s\n", string(m.Data))
        })

        // Unsubscribe
        sub.Unsubscribe()

        // Requests
        msg, _ := nc.Request("help", []byte("help me"), 10*time.Millisecond)
}
