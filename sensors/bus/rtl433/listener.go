package rtl433

type Listener struct {
	Model   string
	Id      string
	SubType string // Optional subtype
	Handler func(*Message)
}

// optionalEqual returns true if a=="" or a==b
func optionalEqual(a, b string) bool {
	return a == "" || a == b
}

func (l *Listener) Accept(m *Message) {
	if l.Handler != nil &&
		l.Model == m.Model &&
		optionalEqual(l.Id, m.ID) &&
		optionalEqual(l.SubType, m.SubType) {
		l.Handler(m)
	}
}
