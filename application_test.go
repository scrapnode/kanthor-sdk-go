package kanthorsdk_test

import (
	"testing"

	"github.com/jaswdr/faker"
	kanthorsdk "github.com/scrapnode/kanthor-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestApplication(t *testing.T) {
	f := faker.New()

	t.Run("success", func(st *testing.T) {
		sdk, err := New()
		assert.Nil(st, err)

		createCtx, cancel := Context()
		defer cancel()
		createReq := &kanthorsdk.ApplicationCreateReq{}
		createReq.SetName(f.App().Name())
		createRes, err := sdk.Application.Create(createCtx, createReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, createRes.Id)
		assert.NotEmpty(st, createRes.WsId)
		assert.NotEmpty(st, createRes.Name)
		assert.Greater(st, *createRes.CreatedAt, int64(0))
		assert.Greater(st, *createRes.UpdatedAt, int64(0))

		getCtx, cancel := Context()
		defer cancel()
		getRes, err := sdk.Application.Get(getCtx, *createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, getRes.Id)
		assert.NotEmpty(st, getRes.WsId)
		assert.NotEmpty(st, getRes.Name)
		assert.Greater(st, *getRes.CreatedAt, int64(0))
		assert.Greater(st, *getRes.UpdatedAt, int64(0))

		updateCtx, cancel := Context()
		defer cancel()
		updateReq := &kanthorsdk.ApplicationUpdateReq{}
		updateReq.SetName(f.App().Name())
		updateRes, err := sdk.Application.Update(updateCtx, *createRes.Id, updateReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, updateRes.Id)
		assert.NotEmpty(st, updateRes.WsId)
		assert.NotEmpty(st, updateRes.Name)
		assert.Greater(st, *updateRes.CreatedAt, int64(0))
		assert.Greater(st, *updateRes.UpdatedAt, int64(0))

		ctx, cancel := Context()
		defer cancel()
		listRes, err := sdk.Application.List(ctx)
		assert.Nil(st, err)
		assert.NotNil(st, listRes.Data)
		assert.GreaterOrEqual(st, *listRes.Count, int64(1))
		assert.Equal(st, listRes.Data[0].Id, createRes.Id)

		deleteCtx, cancel := Context()
		defer cancel()
		deleteRes, err := sdk.Application.Delete(deleteCtx, *createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, deleteRes.Id)
		assert.NotEmpty(st, deleteRes.WsId)
		assert.NotEmpty(st, deleteRes.Name)
		assert.Greater(st, *deleteRes.CreatedAt, int64(0))
		assert.Greater(st, *deleteRes.UpdatedAt, int64(0))
	})
}
