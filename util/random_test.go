package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()
	require.Equal(t, 6, len(owner))
}

func TestRandomCurrency(t *testing.T) {
	currency := RandomCurrency()
	require.Contains(t, []string{"EUR", "USD", "CAD"}, currency)
}
