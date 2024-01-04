package kanthorsdk

import (
	"context"

	"github.com/scrapnode/kanthor-sdk-go/internal/openapi"
)

type (
	MessageCreateReq = openapi.MessageCreateReq
	MessageCreateRes = openapi.MessageCreateRes
)

type Message struct {
	api *openapi.APIClient
}

func (instance *Message) Create(ctx context.Context, req *MessageCreateReq) (*MessageCreateRes, error) {
	request := instance.api.MessageAPI.MessagePost(ctx)
	request = request.Payload(*req)

	response, res, err := request.Execute()
	if err != nil {
		return nil, errorify(err, res)
	}

	return response, nil
}
