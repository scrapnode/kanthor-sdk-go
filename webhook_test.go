package kanthorsdk_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	kanthorsdk "github.com/scrapnode/kanthor-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	f := faker.New().Json()

	t.Run("internal", func(st *testing.T) {
		webhook, err := kanthorsdk.NewWebhook(strings.ReplaceAll(uuid.NewString(), "-", ""))
		assert.Nil(t, err)

		id := uuid.NewString()
		ts := time.Now().UnixMilli()
		payload := []byte(f.String())

		signature, err := webhook.Sign(id, ts, payload)
		assert.Nil(st, err)

		assert.NotEmpty(st, signature)
		assert.Equal(st, 64, len(strings.Split(signature, ",")[1]))

		headers := http.Header{}
		headers.Set(kanthorsdk.HeaderWebhookId, id)
		headers.Set(kanthorsdk.HeaderWebhookTs, fmt.Sprintf("%d", ts))
		headers.Set(kanthorsdk.HeaderWebhookSign, signature)
		err = webhook.Verify(payload, headers)
		assert.Nil(st, err)
	})
}
