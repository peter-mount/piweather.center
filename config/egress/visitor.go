package egress

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/amqp"
)

type EgressVisitor[T any] interface {
	Action(*Action) error
	Amqp(*amqp.Amqp) error
	Metric(*Metric) error
	Publish(*Publish) error
	Script(*Script) error
	GetData() T
	SetData(T) EgressVisitor[T]
}

type visitorCommon[T any] struct {
	action  func(EgressVisitor[T], *Action) error
	amqp    func(EgressVisitor[T], *amqp.Amqp) error
	metric  func(EgressVisitor[T], *Metric) error
	publish func(EgressVisitor[T], *Publish) error
	script  func(EgressVisitor[T], *Script) error
}

type visitor[T any] struct {
	visitorCommon[T]
	data T
}

func (v *visitor[T]) GetData() T {
	return v.data
}

func (v *visitor[T]) SetData(data T) EgressVisitor[T] {
	return &visitor[T]{
		visitorCommon: v.visitorCommon,
		data:          data,
	}
}

func (v *visitor[T]) Script(b *Script) error {
	var err error
	if b != nil {
		if v.script != nil {
			err = v.script(v, b)
		}
		if errors.IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			for _, l := range b.Amqp {
				if err == nil {
					err = v.Amqp(l)
				}
			}
		}

		if err == nil {
			for _, c := range b.Actions {
				if err == nil {
					err = v.Action(c)
				}
			}
		}
	}
	return err
}

func (v *visitor[T]) Action(b *Action) error {
	var err error
	if b != nil {
		if v.action != nil {
			err = v.action(v, b)
		}
		if errors.IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			err = v.Metric(b.Metric)
		}
	}
	return err
}

func (v *visitor[T]) Amqp(b *amqp.Amqp) error {
	var err error
	if b != nil {
		if v.amqp != nil {
			err = v.amqp(v, b)
		}
		if errors.IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor[T]) Metric(b *Metric) error {
	var err error
	if b != nil {
		if v.metric != nil {
			err = v.metric(v, b)
		}
		if errors.IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			for _, p := range b.Publish {
				err = v.Publish(p)
				if err != nil {
					break
				}
			}
		}
	}
	return err
}

func (v *visitor[T]) Publish(b *Publish) error {
	var err error
	if b != nil {
		if v.publish != nil {
			err = v.publish(v, b)
		}
		if errors.IsVisitorStop(err) {
			return nil
		}
	}
	return err
}
