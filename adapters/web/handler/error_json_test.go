package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json Error"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello Json Error"}`), result)
}
