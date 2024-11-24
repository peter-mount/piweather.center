package publisher

// Builder allows for a Publisher to be constructed
type Builder interface {
	// FilterEmpty indicates that the final Publisher should not process empty Reading's
	FilterEmpty() Builder

	// SetId sets the Readings ID
	SetId(string) Builder

	// Log will cause the Publisher to log readings as they are published
	Log() Builder

	DB(DBPublisher) Builder

	// Then allows for a custom Publisher to be included within the Final Publisher.
	Then(Publisher) Builder

	// Build returns the final Publisher
	Build() Publisher
}

type builder struct {
	pub         Publisher
	log         bool
	filterEmpty bool
	dbPublisher bool
}

func NewBuilder() Builder {
	return &builder{}
}

func (b *builder) SetId(id string) Builder {
	return b.Then(setId(id))
}

func (b *builder) FilterEmpty() Builder {
	b.filterEmpty = true
	return b
}

func (b *builder) Log() Builder {
	if b.log {
		return b
	}
	b.log = true
	return b.Then(logPublisher)
}

func (b *builder) DB(db DBPublisher) Builder {
	if b.dbPublisher {
		return b
	}
	b.dbPublisher = true
	return b.Then(dbPublisher(db))
}

func (b *builder) Then(p Publisher) Builder {
	b.pub = b.pub.Then(p)
	return b
}

func (b *builder) Build() Publisher {
	pub := b.pub

	if b.filterEmpty {
		pub = filterEmptyReadings(pub)
	}

	return pub
}
