package location

type MapContainer struct {
	locations Map
}

func (c *MapContainer) GetLocation(n string) *Location {
	if c.locations == nil {
		return nil
	}
	return c.locations.GetLocation(n)
}

func (c *MapContainer) GetLocations() []*Location {
	if c.locations == nil {
		return nil
	}
	return c.locations.GetLocations()
}

func (c *MapContainer) SetLocation(l *Location) (Map, bool) {
	if c.locations == nil {
		c.locations = NewMap()
	}
	m, added := c.locations.SetLocation(l)
	if added && m != nil {
		c.locations = m
	}
	return m, added
}

func (c *MapContainer) MergeLocations(b MapContainer) error {
	if b.locations == nil {
		return nil
	}

	if c.locations == nil {
		c.locations = NewMap()
	}

	return c.locations.Merge(b.locations)
}
