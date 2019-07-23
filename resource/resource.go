package resource

type Resource struct {
	Name             string
	Type             string
	Category         string
	ShortDescription string
	Description      string
	Arguments        []Attribute
	Attributes       []Attribute
	Keywords         []string
	// Metadata is used to fill any random data that's needed
	// on this resource
	Metadata map[string]interface{}
}

type Attribute struct {
	Name        string
	Description string
	// It'll be on the Documentation too
	//Required      bool
}
