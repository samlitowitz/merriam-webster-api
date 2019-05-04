package dictionary

import "context"

type Client interface {
	Lookup(ctx context.Context, word string) ([]*Response, error)
}
