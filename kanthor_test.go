package kanthorsdk_test

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/jaswdr/faker"
	kanthorsdk "github.com/scrapnode/kanthor-sdk-go"
)

func New() (*kanthorsdk.Kanthor, error) {
	file := os.Getenv("TEST_SDK_CREDENTIALS_FILE")
	if file == "" {
		return nil, errors.New("TEST_SDK_CREDENTIALS_FILE must be set")
	}
	credentials, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if string(credentials) == "" {
		return nil, errors.New("TEST_SDK_CREDENTIALS_FILE is empty")
	}

	return kanthorsdk.New(
		string(credentials),
		kanthorsdk.WithHost(os.Getenv("KANTHOR_SDK_HOST")),
		kanthorsdk.WithDebug(os.Getenv("TEST_DEBUG") == "TRUE"),
	)
}

func Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*5)
}

func AppId(sdk *kanthorsdk.Kanthor) (string, error) {
	f := faker.New()

	ctx, cancel := Context()
	defer cancel()

	req := &kanthorsdk.ApplicationCreateReq{}
	req.SetName(f.App().Name())
	res, err := sdk.Application.Create(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Id, nil
}

func EpId(sdk *kanthorsdk.Kanthor) (string, error) {
	f := faker.New()

	appId, err := AppId(sdk)
	if err != nil {
		return "", err
	}

	ctx, cancel := Context()
	defer cancel()

	req := &kanthorsdk.EndpointCreateReq{}
	req.SetAppId(appId)
	req.SetName(f.App().Name())
	req.SetMethod(http.MethodPost)
	req.SetUri(f.Internet().URL())
	res, err := sdk.Endpoint.Create(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Id, nil
}
