package docker

// Repository represents a docker image inside repository
type Repository struct {
	Name   string
	Images []*Image
}

// NewRepository returns a new instance of Repository
func NewRepository(name string) *Repository {
	return &Repository{name, nil}
}
