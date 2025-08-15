package profiles

import "math/rand"

func randSliceItem[T any](slice *[]T) T {
	return (*slice)[rand.Intn(len(*slice))]
}
