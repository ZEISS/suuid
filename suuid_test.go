package suuid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   uuid.UUID
		out  string
	}{
		{
			name: "valid",
			in:   uuid.UUID([]byte("019011ec-eadd-76d0-8347-df0d7c38d19c")),
			out:  "MDE5MDExZWMtZWFkZC03Ng",
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
