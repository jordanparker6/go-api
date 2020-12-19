package core

import (
	"encoding/base64"
	"strings"

	"github.com/google/uuid"
)

// Newid returns a base64 encoded uuid.
func NewId() string {
	id := uuid.New()
	bytes, err := id.MarshalBinary()
	if err != nil {
		panic("UUID encoding error.")
	}

	return strings.TrimRight(base64.URLEncoding.EncodeToString(bytes), "=")
}
