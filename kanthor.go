package kanthorsdk

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/scrapnode/kanthor-sdk-go/internal/openapi"
)

type Kanthor struct {
	Account      *Account
	Application  *Application
	Endpoint     *Endpoint
	EndpointRule *EndpointRule
	Message      *Message
}

func New(credentials string, options ...Option) (*Kanthor, error) {
	var proj Project
	if err := json.Unmarshal(project, &proj); err != nil {
		return nil, err
	}

	conf := openapi.NewConfiguration()
	conf.Scheme = "https"
	conf.Middleware = func(r *http.Request) {
		r.Header.Set("Idempotency-Key", uuid.NewString())
	}

	conf.AddDefaultHeader("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(credentials))))

	h, err := host(credentials, &proj)
	if err != nil {
		return nil, err
	}
	conf.Host = h

	// forllowing https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent#syntax
	conf.UserAgent = fmt.Sprintf("kanthor/%s OpenAPI/go", proj.Version)

	// override configuration with custom options'
	opts := &Options{}
	for _, option := range options {
		option(opts)
	}
	conf.Debug = opts.Debug
	if opts.Host != "" {
		conf.Host = opts.Host
		// if the host is localhost, should change the scheme back to HTTP
		conf.Scheme = scheme(conf.Host)
	}
	if opts.Scheme != "" {
		conf.Scheme = opts.Scheme
	}

	api := openapi.NewAPIClient(conf)
	sdk := &Kanthor{
		Account:     &Account{api: api},
		Application: &Application{api: api},
		Endpoint:    &Endpoint{api: api},
		Message:     &Message{api: api},
	}
	return sdk, nil
}
