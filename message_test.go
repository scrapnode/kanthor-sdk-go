package kanthorsdk_test

import (
	"testing"
	"time"

	kanthorsdk "github.com/scrapnode/kanthor-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestMessageCreate(t *testing.T) {
	sdk, err := New()
	assert.Nil(t, err)

	appId, err := AppId(sdk)
	assert.Nil(t, err)

	t.Run("success", func(st *testing.T) {
		ctx, cancel := Context()
		defer cancel()

		req := &kanthorsdk.MessageCreateReq{
			AppId:   appId,
			Type:    "testing.sdk",
			Headers: map[string]string{"X-Client": "Kanthor SDK Go"},
			Body:    map[string]any{"timestamp": time.Now().UnixMilli()},
		}
		res, err := sdk.Message.Create(ctx, req)
		assert.Nil(t, err)

		assert.NotEmpty(t, res.Id)
	})

	t.Run("success with default value", func(st *testing.T) {
		ctx, cancel := Context()
		defer cancel()

		req := &kanthorsdk.MessageCreateReq{
			AppId: appId,
			Type:  "testing.sdk",
			Body:  map[string]any{"timestamp": time.Now().UnixMilli()},
		}
		res, err := sdk.Message.Create(ctx, req)
		assert.Nil(t, err)

		assert.NotEmpty(t, res.Id)
	})
}
