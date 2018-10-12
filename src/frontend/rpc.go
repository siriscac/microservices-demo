// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"time"

	"github.com/siriscac/microservices-demo/fontend/apiclient/client"
	models "github.com/siriscac/microservices-demo/fontend/apiclient/models"

	"github.com/pkg/errors"
)

const (
	avoidNoopCurrencyConversionRPC = false
)

func (fe *frontendServer) getCurrencies(ctx context.Context) ([]string, error) {
	currs, err := fe.client.CurrencyService.GetSupportedCurrencies(ctx, &currency_service.GetSupportedCurrenciesParams{})
	if err != nil {
		return nil, err
	}
	var out []string
	for _, c := range currs.CurrencyCodes {
		if _, ok := whitelistedCurrencies[c]; ok {
			out = append(out, c)
		}
	}
	return out, nil
}

func (fe *frontendServer) getProducts(ctx context.Context) ([]*models.HipstershopProduct, error) {
	resp, err := fe.client.ProductCatalogService.ListProducts(ctx, &product_catalog_service.ListProductsParams{})
	return resp.Products, err
}

func (fe *frontendServer) getProduct(ctx context.Context, id string) (*models.HipstershopProduct, error) {
	resp, err := fe.client.ProductCatalogService.GetProduct(ctx, &product_catalog_service.GetProductParams{Id: id})
	return resp, err
}

func (fe *frontendServer) getCart(ctx context.Context, userID string) ([]*models.HipstershopCartItem, error) {
	resp, err := fe.client.CartService.GetCart(ctx, &cart_service.GetCartParams{UserId: userID})
	return resp.Items, err
}

func (fe *frontendServer) emptyCart(ctx context.Context, userID string) error {
	_, err := fe.client.CartService.EmptyCart(ctx, &cart_service.EmptyCartParams{UserId: userID})
	return err
}

func (fe *frontendServer) insertCart(ctx context.Context, userID, productID string, quantity int32) error {
	_, err := fe.client.CartService.AddItem(ctx, &cart_service.AddItemParams{
		Body: models.HipstershopAddItemRequest{
			UserId: userID,
			Item: models.HipstershopCartItem{
				ProductId: productID,
				Quantity:  quantity},
			}
		})
	return err
}

func (fe *frontendServer) convertCurrency(ctx context.Context, money *models.HipstershopMoney, currency string) (*models.HipstershopMoney, error) {
	if avoidNoopCurrencyConversionRPC && money.CurrencyCode == currency {
		return money, nil
	}
	return fe.client.CurrencyService.Convert(ctx,
		&currency_service.ConvertParams{
			FromCurrencyCode:   money,
			ToCode: currency
		})
}

func (fe *frontendServer) getShippingQuote(ctx context.Context, items []*models.HipstershopCartItem, currency string) (*models.HipstershopMoney, error) {
	quote, err := fe.client.ShippingService.GetQuote(ctx,
		&shipping_service.GetQuoteParams{
			Body: models.HipstershopGetQuoteRequest{
				Address: nil,
				Items:   items
			}
		})
	if err != nil {
		return nil, err
	}
	localized, err := fe.convertCurrency(ctx, quote.CostUsd, currency)
	return localized, errors.Wrap(err, "failed to convert currency for shipping cost")
}

func (fe *frontendServer) getRecommendations(ctx context.Context, userID string, productIDs []string) ([]*models.HipstershopProduct, error) {
	resp, err := fe.client.RecommendationService.ListRecommendations(ctx, &recommendation_service.ListRecommendationsParams{UserId: userID, ProductIds: productIDs})
	if err != nil {
		return nil, err
	}
	out := make([]*models.HipstershopProduct, len(resp.ProductIds))
	for i, v := range resp.ProductIds {
		p, err := fe.getProduct(ctx, v)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get recommended product info (#%s)", v)
		}
		out[i] = p
	}
	if len(out) > 4 {
		out = out[:4] // take only first four to fit the UI
	}
	return out, err
}

func (fe *frontendServer) getAd(ctx context.Context, ctxKeys []string) ([]*models.HipstershopAd, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	resp, err := fe.client.AdService.GetAds(ctx, &ad_service.GetAdsParams{
		ContextKeys: ctxKeys,
	})
	return resp.Ads, errors.Wrap(err, "failed to get ads")
}
