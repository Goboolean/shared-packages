package resolver

import "context"


// An interface form that every infrastructure package should met
type Transactioner interface {
	Commit() error
	Rollback() error
	Context() context.Context
	Transaction() interface{}
}