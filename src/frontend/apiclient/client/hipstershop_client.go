// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/siriscac/microservices-demo/fontend/apiclient/client/ad_service"
	"github.com/siriscac/microservices-demo/fontend/apiclient/client/cart_service"
	"github.com/siriscac/microservices-demo/fontend/apiclient/client/checkout_service"
	"github.com/siriscac/microservices-demo/fontend/apiclient/client/currency_service"
	"github.com/siriscac/microservices-demo/fontend/apiclient/client/email_service"
	"github.com/siriscac/microservices-demo/fontend/apiclient/client/payment_service"
	"github.com/siriscac/microservices-demo/fontend/apiclient/client/product_catalog_service"
	"github.com/siriscac/microservices-demo/fontend/apiclient/client/recommendation_service"
	"github.com/siriscac/microservices-demo/fontend/apiclient/client/shipping_service"
)

// Default hipstershop HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new hipstershop HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Hipstershop {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new hipstershop HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Hipstershop {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new hipstershop client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Hipstershop {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Hipstershop)
	cli.Transport = transport

	cli.AdService = ad_service.New(transport, formats)

	cli.CartService = cart_service.New(transport, formats)

	cli.CheckoutService = checkout_service.New(transport, formats)

	cli.CurrencyService = currency_service.New(transport, formats)

	cli.EmailService = email_service.New(transport, formats)

	cli.PaymentService = payment_service.New(transport, formats)

	cli.ProductCatalogService = product_catalog_service.New(transport, formats)

	cli.RecommendationService = recommendation_service.New(transport, formats)

	cli.ShippingService = shipping_service.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Hipstershop is a client for hipstershop
type Hipstershop struct {
	AdService *ad_service.Client

	CartService *cart_service.Client

	CheckoutService *checkout_service.Client

	CurrencyService *currency_service.Client

	EmailService *email_service.Client

	PaymentService *payment_service.Client

	ProductCatalogService *product_catalog_service.Client

	RecommendationService *recommendation_service.Client

	ShippingService *shipping_service.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Hipstershop) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.AdService.SetTransport(transport)

	c.CartService.SetTransport(transport)

	c.CheckoutService.SetTransport(transport)

	c.CurrencyService.SetTransport(transport)

	c.EmailService.SetTransport(transport)

	c.PaymentService.SetTransport(transport)

	c.ProductCatalogService.SetTransport(transport)

	c.RecommendationService.SetTransport(transport)

	c.ShippingService.SetTransport(transport)

}
