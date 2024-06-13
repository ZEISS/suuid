package suuid

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// SUUID is the 
type SUUID []byte

// MarshalUUID is 

// Encode return as a short 22 character slug.
func Encode(id uuid.UUID) string {
	return base64.URLEncoding.EncodeToString([]byte(id))[:22] // Drop '==' padding
}

// Decode returns the uuid.UUID as a nice slug.
func Decode(suuid string) (uuid.UUID, error) {
	id, err := base64.URLEncoding.DecodeString(suuid + "==")
	if err != nil {
		return nil, err
	}

	return id
}
