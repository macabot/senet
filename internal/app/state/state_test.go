package state_test

import (
	"encoding/json"
	"testing"

	"github.com/macabot/senet/internal/app/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPageMarshalUnmarshal(t *testing.T) {
	b, err := json.Marshal(state.StartPage)
	require.NoError(t, err)
	assert.Equal(t, `"Start"`, string(b))

	var page state.Page
	require.NoError(t, json.Unmarshal(b, &page))
	assert.Equal(t, state.StartPage, page)
}
