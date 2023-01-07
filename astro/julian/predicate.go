package julian

type Predicate func(Day, Day) bool

func After(a, b Day) bool { return a.After(b) }

func Before(a, b Day) bool { return a.Before(b) }

func Equal(a, b Day) bool { return a == b }

func False(_, _ Day) bool { return false }

func True(_, _ Day) bool { return true }

func AfterEqual(a, b Day) bool { return a.After(b) || a == b }

func BeforeEqual(a, b Day) bool { return a.Before(b) || a == b }
