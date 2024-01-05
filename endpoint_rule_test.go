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
		createReq := &kanthorsdk.EndpointRuleCreateReq{
			EpId:                epId,
			Name:                f.App().Name(),
			Exclusionary:        true,
			Priority:            1,
			ConditionSource:     "type",
			ConditionExpression: "any::",
		}
		createRes, err := sdk.EndpointRule.Create(createCtx, createReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, createRes.Id)
		assert.NotEmpty(st, createRes.EpId)
		assert.NotEmpty(st, createRes.Name)
		assert.True(st, createRes.Exclusionary)
		assert.NotEmpty(st, createRes.ConditionExpression)
		assert.NotEmpty(st, createRes.ConditionSource)
		assert.Greater(st, createRes.CreatedAt, int64(0))
		assert.Greater(st, createRes.UpdatedAt, int64(0))
		assert.Equal(st, int64(1), createRes.Priority)

		updateCtx, cancel := Context()
		defer cancel()
		updateReq := &kanthorsdk.EndpointRuleUpdateReq{
			Name:                f.App().Name(),
			Exclusionary:        false,
			Priority:            createRes.Priority + 1,
			ConditionSource:     createReq.ConditionSource,
			ConditionExpression: createReq.ConditionExpression,
		}
		updateRes, err := sdk.EndpointRule.Update(updateCtx, createRes.Id, updateReq)
		assert.Nil(st, err)
		assert.NotEmpty(st, updateRes.Id)
		assert.NotEmpty(st, updateRes.EpId)
		assert.NotEqual(st, updateRes.Name, createRes.Name)
		assert.False(st, updateRes.Exclusionary)
		assert.GreaterOrEqual(st, updateRes.Priority, int64(createRes.Priority+1))
		assert.NotEmpty(st, updateRes.ConditionExpression)
		assert.NotEmpty(st, updateRes.ConditionSource)
		assert.Greater(st, updateRes.UpdatedAt, createRes.UpdatedAt)
		assert.Equal(st, createRes.CreatedAt, updateRes.CreatedAt)

		getCtx, cancel := Context()
		defer cancel()
		getRes, err := sdk.EndpointRule.Get(getCtx, createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, getRes.Id)
		assert.NotEmpty(st, getRes.EpId)
		assert.False(st, getRes.Exclusionary)
		assert.Greater(st, getRes.UpdatedAt, createRes.UpdatedAt)
		assert.Equal(st, updateRes.Name, getRes.Name)
		assert.Equal(st, updateRes.Priority, getRes.Priority)
		assert.Equal(st, createRes.ConditionExpression, getRes.ConditionExpression)
		assert.Equal(st, createRes.ConditionSource, getRes.ConditionSource)
		assert.Equal(st, createRes.CreatedAt, getRes.CreatedAt)

		ctx, cancel := Context()
		defer cancel()
		listRes, err := sdk.EndpointRule.List(ctx, kanthorsdk.WithIds([]string{getRes.Id}))
		assert.Nil(st, err)
		assert.NotNil(st, listRes.Data)
		assert.GreaterOrEqual(st, listRes.Count, int64(1))
		assert.Equal(st, createRes.Id, listRes.Data[0].Id)

		deleteCtx, cancel := Context()
		defer cancel()
		deleteRes, err := sdk.EndpointRule.Delete(deleteCtx, createRes.Id)
		assert.Nil(st, err)
		assert.NotEmpty(st, deleteRes.Id)
	})
}
