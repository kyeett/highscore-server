package client

import (
	"crypto/md5"
	"fmt"

	"github.com/google/uuid"
)

func encodedUUIDFromBytes(input []byte) (uuid.UUID, error) {
	s := fmt.Sprintf("%x", md5.Sum(input))

	id := fmt.Sprintf("%s-%s-%s-%s-%s",
		s[0:8], s[8:12], s[12:16], s[16:20], s[20:])

	return uuid.Parse(id)
}
