package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func hashPassword(t *testing.T) (string, string, error) {
	password := RandomString(6)
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	return password, hashedPassword, err
}

func TestCheckPassword(t *testing.T) {
	password, hashedPassword, err := hashPassword(t)
	require.NoError(t, err)
	require.NotEmpty(t, password)
	require.NotEmpty(t, hashedPassword)
	require.NotEqual(t, password, hashedPassword)

	err = CheckPassword(hashedPassword, password)
	require.NoError(t, err)
}
