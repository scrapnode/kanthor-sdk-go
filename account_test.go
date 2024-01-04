package kanthorsdk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountGet(t *testing.T) {
	t.Run("success", func(st *testing.T) {
		sdk, err := New()
		assert.Nil(t, err)

		ctx, cancel := Context()
		defer cancel()
		res, err := sdk.Account.Get(ctx)
		assert.Nil(t, err)

		assert.NotNil(t, res.Account)
		assert.NotEmpty(t, res.Account.Sub)
		assert.NotEmpty(t, res.Account.Name)
		assert.NotNil(t, res.Account.Metadata)
		assert.Greater(t, len(res.Account.Metadata), 0)
	})
}
