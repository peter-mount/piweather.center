package lang

import "errors"

type Visitor[T any] interface {
	At(*At) error
	Cron(*Cron) error
	Every(*Every) error
	Rule(*Rule) error
	Schedule(*Schedule) error
	Script(*Script) error
	TaskRule(*TaskRule) error
	GetData() T
	SetData(T) Visitor[T]
}

type visitorCommon[T any] struct {
	at       func(Visitor[T], *At) error
	cron     func(Visitor[T], *Cron) error
	every    func(Visitor[T], *Every) error
	rule     func(Visitor[T], *Rule) error
	schedule func(Visitor[T], *Schedule) error
	script   func(Visitor[T], *Script) error
	taskRule func(Visitor[T], *TaskRule) error
}

type visitor[T any] struct {
	visitorCommon[T]
	data T
}

func (v *visitor[T]) GetData() T {
	return v.data
}

func (v *visitor[T]) SetData(data T) Visitor[T] {
	return &visitor[T]{
		visitorCommon: v.visitorCommon,
		data:          data,
	}
}

// VisitorStop is an error which causes the current step in a Visitor to stop processing.
// It's used to enable a Visitor to handle all processing of a node within itself rather
// than the Visitor proceeding to any child nodes of that node.
var VisitorStop = errors.New("visitor stop")

func IsVisitorStop(err error) bool {
	return err != nil && errors.Is(err, VisitorStop)
}

func (v *visitor[T]) Script(b *Script) error {
	var err error
	if b != nil {
		if v.script != nil {
			err = v.script(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			for _, l := range b.Rules {
				if err == nil {
					err = v.Rule(l)
				}
			}
		}
	}
	return err
}

func (v *visitor[T]) At(b *At) error {
	var err error
	if b != nil {
		if v.cron != nil {
			err = v.at(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor[T]) Cron(b *Cron) error {
	var err error
	if b != nil {
		if v.cron != nil {
			err = v.cron(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor[T]) Every(b *Every) error {
	var err error
	if b != nil {
		if v.every != nil {
			err = v.every(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor[T]) Rule(b *Rule) error {
	var err error
	if b != nil {
		if v.rule != nil {
			err = v.rule(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor[T]) Schedule(b *Schedule) error {
	var err error
	if b != nil {
		if v.schedule != nil {
			err = v.schedule(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			err = v.At(b.At)
		}

		if err == nil {
			err = v.Every(b.Every)
		}

		if err == nil {
			err = v.Cron(b.Cron)
		}
	}
	return err
}

func (v *visitor[T]) TaskRule(b *TaskRule) error {
	var err error
	if b != nil {
		if v.taskRule != nil {
			err = v.taskRule(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}
