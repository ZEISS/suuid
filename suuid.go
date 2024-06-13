package suuid

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// SUUID is the
type SUUID struct {
	data string
}

// String implements the fmt.Stringer interface.
func (s *SUUID) String() string {
	return s.data
}

// MarshalUUID is the representation of the UUID as a SUUID.
func (s *SUUID) MarshalUUID(u uuid.UUID) error {
	b, err := u.MarshalBinary()
	if err != nil {
		return err
	}

	s.data = base64.URLEncoding.EncodeToString(b)[:22]

	return nil
}

// UnmarshalUUID is the uuid prepresentation of the SUUID.
func (s *SUUID) UnmarshalUUID() (uuid.UUID, error) {
	id, err := base64.URLEncoding.DecodeString(s.data + "==")
	if err != nil {
		return uuid.UUID{}, err
	}

	in, err := uuid.Parse(string(id))
	if err != nil {
		return in, err
	}

	return in, nil
}
