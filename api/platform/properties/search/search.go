package search

type Getter interface {
	GetAll() []Search
}

type Adder interface {
	Add(search Search)
}

type Search struct {
	SearchTerm string `json:"search_term"`
}

type Repo struct {
	Searches []Search
}

func New() *Repo {
	return &Repo{
		Searches: []Search{},
	}
}

func (r *Repo) Add(search Search) {
	r.Searches = append(r.Searches, search)
}

func (r *Repo) GetAll() []Search {
	return r.Searches
}
