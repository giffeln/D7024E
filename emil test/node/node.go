package node

import (
	"crypto/rand"
)

const IDLength = 256
const IDByteLength = IDLength / 8

var MyId = newId()

type ID [IDByteLength]byte

type randRead func([]byte) (int, error)

var rng randRead

func init() {
	rng = rand.Read
}

func newId() (id ID) {
	_, err := rng(id[:])
	if err != nil {
		panic(err)
	}
	return id
}
