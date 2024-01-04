package kanthorsdk_test

import (
	"testing"

	"github.com/jaswdr/faker"
	kanthorsdk "github.com/scrapnode/kanthor-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestEndpointRule(t *testing.T) {
	sdk, err := New()
	assert.Nil(t, err)

	epId, err := EpId(sdk)
	assert.Nil(t, err)

	f := faker.New()

	t.Run("success", func(st *testing.T) {
		createCtx, cancel := Context()
		defer cancel()
		createReq := &kanthorsdk.EndpointRuleCreateReq{}
		createReq.SetEpId(epId)
		createReq.SetName(f.App().Name())
		createReq.SetExclusionary(true)
		createReq.SetPriority(1)
		createReq.SetConditionExpression("::any")
		createReq.SetConditionSource("type")
		createRes, err := sdk.EndpointRule.Create(createCtx, createReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, createRes.Id)
		assert.NotEmpty(st, createRes.EpId)
		assert.NotEmpty(st, createRes.Name)
		assert.True(st, createRes.Exclusionary)
		assert.GreaterOrEqual(st, createRes.Priority, int64(1))
		assert.NotEmpty(st, createRes.ConditionExpression)
		assert.NotEmpty(st, createRes.ConditionSource)
		assert.Greater(st, createRes.CreatedAt, int64(0))
		assert.Greater(st, createRes.UpdatedAt, int64(0))

		getCtx, cancel := Context()
		defer cancel()
		getRes, err := sdk.EndpointRule.Get(getCtx, createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, getRes.Id)
		assert.NotEmpty(st, getRes.EpId)
		assert.NotEmpty(st, getRes.Name)
		assert.True(st, getRes.Exclusionary)
		assert.GreaterOrEqual(st, getRes.Priority, int64(1))
		assert.NotEmpty(st, getRes.ConditionExpression)
		assert.NotEmpty(st, getRes.ConditionSource)
		assert.Greater(st, getRes.CreatedAt, int64(0))
		assert.Greater(st, getRes.UpdatedAt, int64(0))

		updateCtx, cancel := Context()
		defer cancel()
		updateReq := &kanthorsdk.EndpointRuleUpdateReq{}
		updateReq.SetName(f.App().Name())
		updateReq.SetPriority(createRes.Priority + 1)
		updateReq.SetExclusionary(false)
		updateReq.SetConditionSource(getRes.ConditionSource)
		updateReq.SetConditionExpression(getRes.ConditionExpression)
		updateRes, err := sdk.EndpointRule.Update(updateCtx, createRes.Id, updateReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, updateRes.Id)
		assert.NotEmpty(st, updateRes.EpId)
		assert.NotEmpty(st, updateRes.Name)
		assert.False(st, updateRes.Exclusionary)
		assert.GreaterOrEqual(st, updateRes.Priority, int64(createRes.Priority+1))
		assert.NotEmpty(st, updateRes.ConditionExpression)
		assert.NotEmpty(st, updateRes.ConditionSource)
		assert.Greater(st, updateRes.CreatedAt, int64(0))
		assert.Greater(st, updateRes.UpdatedAt, int64(0))

		ctx, cancel := Context()
		defer cancel()
		listRes, err := sdk.EndpointRule.List(ctx)
		assert.Nil(st, err)
		assert.NotNil(st, listRes.Data)
		assert.GreaterOrEqual(st, listRes.Count, int64(1))
		assert.Equal(st, listRes.Data[0].Id, createRes.Id)

		deleteCtx, cancel := Context()
		defer cancel()
		deleteRes, err := sdk.EndpointRule.Delete(deleteCtx, createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, deleteRes.Id)
	})
}
