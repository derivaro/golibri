package golibri

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitCh struct {
	Channel *amqp.Channel
}

func RabbitSendMsg(rchan RabbitCh, queue string, payload string) string {
	er0 := rchan.RabbitPush(queue, []byte(payload))
	if er0 != nil {
		fmt.Println("--> Rabbit error with payload: " + payload)
		return "ko"
	} else {
		return "ok"
	}
}

// GetConn -
func RabbitGetConn(rabbitURL string) (RabbitCh, error) {
	connr, err := amqp.Dial(rabbitURL)
	if err != nil {

		return RabbitCh{}, err
	}
	ch, err := connr.Channel()
	return RabbitCh{
		Channel: ch,
	}, err
}

func (conn RabbitCh) RabbitPush(routingKey string, data []byte) error {
	return conn.Channel.Publish(
		"",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}
