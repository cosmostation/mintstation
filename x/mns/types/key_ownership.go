package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OwnershipKeyPrefix is the prefix to retrieve all Ownership
	OwnershipKeyPrefix = "Ownership/value/"
)

// OwnershipKey returns the store key to retrieve a Ownership from the index fields
func OwnershipKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
