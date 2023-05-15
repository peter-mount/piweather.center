package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Task func(amqp.Delivery) error

func (a Task) Do(msg amqp.Delivery) error {
	if a != nil {
		return a(msg)
	}
	return nil
}

func (a Task) Then(b Task) Task {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(delivery amqp.Delivery) error {
		err := a.Do(delivery)
		if err == nil {
			err = b.Do(delivery)
		}
		return err
	}
}

type Predicate func(amqp.Delivery) bool

func (a Task) IfRoute(t Predicate, b Task) Task {
	return a.Then(func(delivery amqp.Delivery) error {
		if t(delivery) {
			return b.Do(delivery)
		}
		return nil
	})
}

func (a Task) IfRouteEquals(r string, b Task) Task {
	return a.IfRoute(func(delivery amqp.Delivery) bool {
		return delivery.RoutingKey == r
	}, b)
}

func (a Task) IfRouteNotEquals(r string, b Task) Task {
	return a.IfRoute(func(delivery amqp.Delivery) bool {
		return delivery.RoutingKey != r
	}, b)
}

func (q *Queue) nakTask(f Task) Task {
	return func(msg amqp.Delivery) error {
		if err := f(msg); err != nil {
			q.logError(err)
			q.logError(msg.Nack(false, true))
			return err
		}

		q.logError(msg.Ack(false))
		return nil
	}
}
