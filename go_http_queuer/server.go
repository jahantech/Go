package main

import (
	"io"
	"net/http"
	"github.com/streadway/amqp"
)

func messagePoster(w http.ResponseWriter, r *http.Request) {
	publishMessage("test")
	io.WriteString(w,"Msg uploaded")
}

func main() {
	http.HandleFunc("/", messagePoster)
	http.ListenAndServe(":8000", nil)
}

func publishMessage(a string) int {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil{
		return 1
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{
		return 2
	}
	defer ch.Close()


	q, err := ch.QueueDeclare(
	  "msg", // name
	  false,   // durable
	  false,   // delete when unused
	  false,   // exclusive
	  false,   // no-wait
	  nil,     // arguments
	)
	
	if err !=nil {
		return 3
	}

	body := "Web traffic"
	err = ch.Publish(
	  "",     // exchange
	  q.Name, // routing key
	  false,  // mandatory
	  false,  // immediate
	  amqp.Publishing {
	    ContentType: "text/plain",
	    Body:        []byte(body),
	  })
	if err != nil {
		return 5
	}
	return 0
}
