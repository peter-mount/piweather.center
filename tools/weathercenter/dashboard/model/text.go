package model

func init() {
	Register("text", func() Instance { return &Text{} })
}

type Text struct {
	Component `yaml:",inline"`
	Text      string `yaml:"text"`
}
