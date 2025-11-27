package put

import (
	"context"
)

type Mocked interface {
	Predicate(context.Context, any) (bool, error)
}
