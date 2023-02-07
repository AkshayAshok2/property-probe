package properties

type Getter interface {
	GetAll() []Property
}

type Adder interface {
	Add(property Property)
}

type Property struct {
	SearchTerm string `json:"search_term"`
	// Owner      string `json:"owner"`
	// Address    string `json:"address"`
}

type Repo struct {
	Properties []Property
}

func New() *Repo {
	return &Repo{
		Properties: []Property{},
	}
}

func (r *Repo) Add(property Property) {
	r.Properties = append(r.Properties, property)
}

func (r *Repo) GetAll() []Property {
	return r.Properties
}
