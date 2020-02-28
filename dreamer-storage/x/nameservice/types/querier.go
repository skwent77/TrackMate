package types

import "strings"

// Query Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
	//package json impolements encoding and decoding of JSON as defined in RFC 
	//mapping between JSON and Go values is described in documentation
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value
}

// Query Result Payload for a names query
type QueryResNames []string

// implement fmt.Stringer
func (n QueryResNames) String() string {
	return strings.Join(n[:], "\n")
}
