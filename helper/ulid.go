package helper

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func NewID() (string, error) {
	defaultEntropySource := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	ul, err := ulid.New(ulid.Timestamp(time.Now()), defaultEntropySource)
	if err != nil {
		return "", err
	}
	return ul.String(), nil
}
