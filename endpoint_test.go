package kanthorsdk_test

import (
	"net/http"
	"testing"

	"github.com/jaswdr/faker"
	kanthorsdk "github.com/scrapnode/kanthor-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestEndpoint(t *testing.T) {
	sdk, err := New()
	assert.Nil(t, err)

	appId, err := AppId(sdk)
	assert.Nil(t, err)

	f := faker.New()

	t.Run("success", func(st *testing.T) {
		createCtx, cancel := Context()
		defer cancel()
		createReq := &kanthorsdk.EndpointCreateReq{
			AppId:  appId,
			Name:   f.App().Name(),
			Method: http.MethodPost,
			Uri:    f.Internet().URL(),
		}
		createRes, err := sdk.Endpoint.Create(createCtx, createReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, createRes.Id)
		assert.NotEmpty(st, createRes.AppId)
		assert.NotEmpty(st, createRes.Name)
		assert.NotEmpty(st, createRes.Method)
		assert.NotEmpty(st, createRes.Uri)
		assert.Greater(st, createRes.CreatedAt, int64(0))
		assert.Greater(st, createRes.UpdatedAt, int64(0))

		updateCtx, cancel := Context()
		defer cancel()
		updateReq := &kanthorsdk.EndpointUpdateReq{
			Name:   f.App().Name(),
			Method: http.MethodPut,
		}
		updateRes, err := sdk.Endpoint.Update(updateCtx, createRes.Id, updateReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, updateRes.Id)
		assert.NotEmpty(st, updateRes.AppId)
		assert.NotEqual(st, updateRes.Name, createRes.Name)
		assert.NotEqual(st, updateRes.Method, createRes.Method)
		assert.Equal(st, updateRes.CreatedAt, createRes.CreatedAt)
		assert.Greater(st, updateRes.UpdatedAt, createRes.UpdatedAt)

		getCtx, cancel := Context()
		defer cancel()
		getRes, err := sdk.Endpoint.Get(getCtx, createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, getRes.Id)
		assert.NotEmpty(st, getRes.AppId)
		assert.Equal(st, getRes.Name, updateRes.Name)
		assert.Equal(st, getRes.Method, updateRes.Method)
		assert.Equal(st, getRes.Uri, createRes.Uri)
		assert.Equal(st, getRes.CreatedAt, createRes.CreatedAt)
		assert.Greater(st, getRes.UpdatedAt, createRes.UpdatedAt)

		ctx, cancel := Context()
		defer cancel()
		listRes, err := sdk.Endpoint.List(ctx, kanthorsdk.WithIds([]string{getRes.Id}))
		assert.Nil(st, err)
		assert.NotNil(st, listRes.Data)
		assert.GreaterOrEqual(st, listRes.Count, int64(1))
		assert.Equal(st, listRes.Data[0].Id, createRes.Id)

		deleteCtx, cancel := Context()
		defer cancel()
		deleteRes, err := sdk.Endpoint.Delete(deleteCtx, createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, deleteRes.Id)
	})
}
