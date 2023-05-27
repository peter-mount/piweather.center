package graph

// Common fields shared by all graph types.
//
// Note: For this to work you include it like the following, otherwise the yaml
// will not unmarshal correctly:
//
//	type MyStruct struct {
//	  Common `yaml:",inline"`
//	  // other fields here
//	}
type Common struct {
	// Min value for Y axis, default is based on data
	Min *float64 `yaml:"min,omitempty"`
	// Max value for Y axis, default is based on data
	Max *float64 `yaml:"max,omitempty"`
}
