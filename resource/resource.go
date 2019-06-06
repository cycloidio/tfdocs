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
}

type Attribute struct {
	Name        string
	Description string
	// It'll be on the Documentation too
	//Required      bool
}
