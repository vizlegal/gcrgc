package docker

// Registry represents a docker registry
type Registry struct {
	Name         string
	Repositories []Repository
}

// NewRegistry returns a new instance of Registry
func NewRegistry(name string, repos []Repository) *Registry {
	return &Registry{name, repos}
}

// ContainsRepository checks if a repository is present in the Registry
func (r Registry) ContainsRepository(name string) bool {
	for _, item := range r.Repositories {
		if name == item.Name {
			return true
		}
	}

	return false
}
