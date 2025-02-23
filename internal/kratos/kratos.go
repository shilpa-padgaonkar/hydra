// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package kratos

import (
	"context"
	"fmt"
	"net/url"

	"go.opentelemetry.io/otel/attribute"

	"github.com/ory/hydra/v2/driver/config"
	"github.com/ory/hydra/v2/x"
	client "github.com/ory/kratos-client-go"
	"github.com/ory/x/httpx"
	"github.com/ory/x/otelx"
)

type (
	dependencies interface {
		config.Provider
		x.HTTPClientProvider
		x.TracingProvider
		x.RegistryLogger
	}
	Provider interface {
		Kratos() Client
	}
	Client interface {
		DisableSession(ctx context.Context, identityProviderSessionID string) error
	}
	Default struct {
		dependencies
	}
)

func New(d dependencies) Client {
	return &Default{dependencies: d}
}

func (k *Default) DisableSession(ctx context.Context, identityProviderSessionID string) (err error) {
	ctx, span := k.Tracer(ctx).Tracer().Start(ctx, "kratos.DisableSession")
	otelx.End(span, &err)

	adminURL, ok := k.Config().KratosAdminURL(ctx)
	span.SetAttributes(attribute.String("admin_url", fmt.Sprintf("%+v", adminURL)))
	if !ok {
		span.SetAttributes(attribute.Bool("skipped", true))
		span.SetAttributes(attribute.String("reason", "kratos admin url not set"))

		return nil
	}

	if identityProviderSessionID == "" {
		span.SetAttributes(attribute.Bool("skipped", true))
		span.SetAttributes(attribute.String("reason", "kratos session ID is empty"))

		return nil
	}

	configuration := k.clientConfiguration(ctx, adminURL)
	if header := k.Config().KratosRequestHeader(ctx); header != nil {
		configuration.HTTPClient.Transport = httpx.WrapTransportWithHeader(configuration.HTTPClient.Transport, header)
	}
	kratos := client.NewAPIClient(configuration)
	_, err = kratos.IdentityApi.DisableSession(ctx, identityProviderSessionID).Execute()

	return err

}

func (k *Default) clientConfiguration(ctx context.Context, adminURL *url.URL) *client.Configuration {
	configuration := client.NewConfiguration()
	configuration.Servers = client.ServerConfigurations{{URL: adminURL.String()}}
	configuration.HTTPClient = k.HTTPClient(ctx).StandardClient()

	return configuration
}
