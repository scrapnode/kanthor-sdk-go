package kanthorsdk_test

import (
	"context"
	"errors"
	"os"
	"time"

	kanthorsdk "github.com/scrapnode/kanthor-sdk-go"
)

func New() (*kanthorsdk.Kanthor, error) {
	file := os.Getenv("TEST_KANTHOR_SDK_CREDENTIALS_FILE")
	if file == "" {
		return nil, errors.New("TEST_KANTHOR_SDK_CREDENTIALS_FILE must be set")
	}
	credentials, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if string(credentials) == "" {
		return nil, errors.New("TEST_KANTHOR_SDK_CREDENTIALS_FILE is empty")
	}

	return kanthorsdk.New(
		string(credentials),
		kanthorsdk.WithHost(os.Getenv("TEST_KANTHOR_SDK_HOST")),
	)
}

func Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*5)
}

func AppId(sdk *kanthorsdk.Kanthor) (string, error) {
	ctx, cancel := Context()
	defer cancel()

	req := &kanthorsdk.ApplicationCreateReq{}
	req.SetName("testing")
	res, err := sdk.Application.Create(ctx, req)
	if err != nil {
		return "", err
	}

	return *res.Id, nil
}
