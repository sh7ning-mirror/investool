package eniu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathCode(t *testing.T) {
	code := _e.GetPathCode(_ctx, "002459.SZ")
	require.Equal(t, "sz002459", code)
}

func TestQueryHistoryStockPrice(t *testing.T) {
	data, err := _e.QueryHistoryStockPrice(_ctx, "002459.SZ")
	require.Nil(t, err)
	require.NotEmpty(t, data.Date)
	t.Log(data)
}
