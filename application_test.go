package kanthorsdk_test

import (
	"testing"

	"github.com/jaswdr/faker"
	kanthorsdk "github.com/scrapnode/kanthor-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestApplication(t *testing.T) {
	sdk, err := New()
	assert.Nil(t, err)

	f := faker.New()

	t.Run("success", func(st *testing.T) {
		createCtx, cancel := Context()
		defer cancel()
		createReq := &kanthorsdk.ApplicationCreateReq{
			Name: f.App().Name(),
		}
		createRes, err := sdk.Application.Create(createCtx, createReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, createRes.Id)
		assert.NotEmpty(st, createRes.WsId)
		assert.NotEmpty(st, createRes.Name)
		assert.Greater(st, createRes.CreatedAt, int64(0))
		assert.Greater(st, createRes.UpdatedAt, int64(0))

		updateCtx, cancel := Context()
		defer cancel()
		updateReq := &kanthorsdk.ApplicationUpdateReq{
			Name: f.App().Name(),
		}
		updateRes, err := sdk.Application.Update(updateCtx, createRes.Id, updateReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, updateRes.Id)
		assert.NotEmpty(st, updateRes.WsId)
		assert.NotEqual(st, updateRes.Name, createRes.Name)
		assert.Greater(st, updateRes.UpdatedAt, createRes.UpdatedAt)
		assert.Equal(st, createRes.CreatedAt, updateRes.CreatedAt)

		getCtx, cancel := Context()
		defer cancel()
		getRes, err := sdk.Application.Get(getCtx, createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, getRes.Id)
		assert.NotEmpty(st, getRes.WsId)
		assert.Equal(st, getRes.Name, updateRes.Name)
		assert.Greater(st, getRes.UpdatedAt, createRes.UpdatedAt)
		assert.Equal(st, createRes.CreatedAt, getRes.CreatedAt)

		ctx, cancel := Context()
		defer cancel()
		listRes, err := sdk.Application.List(ctx, kanthorsdk.WithIds([]string{getRes.Id}))
		assert.Nil(st, err)
		assert.NotNil(st, listRes.Data)
		assert.GreaterOrEqual(st, listRes.Count, int64(1))
		assert.Equal(st, createRes.Id, listRes.Data[0].Id)

		deleteCtx, cancel := Context()
		defer cancel()
		deleteRes, err := sdk.Application.Delete(deleteCtx, createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, deleteRes.Id)
	})
}
