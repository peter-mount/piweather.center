package location

type MapContainer struct {
	locations Map
}

func (c *MapContainer) GetLocation(n string) *Location {
	return c.locations.GetLocation(n)
}

func (c *MapContainer) GetLocations() []*Location {
	return c.locations.GetLocations()
}

func (c *MapContainer) SetLocation(l *Location) (Map, bool) {
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

	m, err := c.locations.Merge(b.locations)
	if err == nil && m != nil {
		c.locations = m
	}
	return err
}
