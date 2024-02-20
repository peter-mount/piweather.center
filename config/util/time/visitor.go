package time

type TimeVisitor interface {
	Duration(*Duration) error
	Time(*Time) error
}
