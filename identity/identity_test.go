package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentity(t *testing.T) {
	cases := []struct {
		description string
		value       int
		expected    int
	}{
		{
			description: "sample",
			value:       10,
			expected:    10,
		},
	}
	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			actual := Identity(tt.value)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
