package models

// Bundled holds the object's store's name
//
// swagger: model
type Bundled struct {
	// Bundle tracks the name of the store containing this object.
	// This field is read-only, and cannot be changed via the API.
	//
	// read only: true
	Bundle string
}

type Bundler interface {
	SetBundle(string)
	GetBundle() string
}

// SetBundle sets the name of the content layer holding this object.
func (b *Bundled) SetBundle(name string) {
	b.Bundle = name
}

// GetBundle gets the name of the content layer holding this object.
func (b *Bundled) GetBundle() string {
	return b.Bundle
}
