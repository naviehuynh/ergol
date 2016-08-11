package types

// Filter modifies Log
type Filter interface {
	Apply() Log
}
