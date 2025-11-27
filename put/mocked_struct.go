package put

import (
	"context"
)

type MockedStruct struct {
}

func (*MockedStruct) Predicate(ctx context.Context, x any) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	switch y := x.(type) {
	case bool:
		return y, nil

	case uint, uint8, uint16, uint32, uint64,
		int, int8, int16, int32, int64,
		float64, float32:
		return y != 0, nil

	case string:
		return len(y) != 0, nil

	default:
		return false, nil
	}
}
