package suuid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestMarshalUUID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   uuid.UUID
		out  string
	}{
		{
			name: "valid v7",
			in:   uuid.MustParse("019011ec-eadd-76d0-8347-df0d7c38d19c"),
			out:  "MDE5MDExZWMtZWFkZC03NmQwLTgzNDctZGYwZDdjMzhkMTlj",
		},
		{
			name: "valid v4",
			in:   uuid.MustParse("540f9987-b0b2-421b-ac56-e3942ccea1c2"),
			out:  "NTQwZjk5ODctYjBiMi00MjFiLWFjNTYtZTM5NDJjY2VhMWMy",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SUUID
			require.NoError(t, s.MarshalUUID(tt.in))
			require.Equal(t, tt.out, s.String())
		})
	}
}

func TestUnmarshalUUID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  uuid.UUID
	}{
		{
			name: "valid v7",
			in:   "MDE5MDExZWMtZWFkZC03NmQwLTgzNDctZGYwZDdjMzhkMTlj",
			out:  uuid.MustParse("019011ec-eadd-76d0-8347-df0d7c38d19c"),
		},
		{
			name: "valid v4",
			in:   "NTQwZjk5ODctYjBiMi00MjFiLWFjNTYtZTM5NDJjY2VhMWMy",
			out:  uuid.MustParse("540f9987-b0b2-421b-ac56-e3942ccea1c2"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Parse(tt.in)
			u, err := s.UnmarshalUUID()
			require.NoError(t, err)
			require.Equal(t, tt.out, u)
		})
	}
}
