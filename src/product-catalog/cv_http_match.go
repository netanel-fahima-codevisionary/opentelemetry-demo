// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package main

import "net/http"

// CodeVisionary discovers net/http HandleFunc, not gRPC; names match Coralogix spans.
type ProductCatalogService struct{}

func (s *ProductCatalogService) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *ProductCatalogService) ListProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func registerCVHTTPMatchRoutes() {
	svc := &ProductCatalogService{}
	http.HandleFunc("POST /oteldemo.ProductCatalogService/GetProduct", svc.GetProduct)
	http.HandleFunc("POST /oteldemo.ProductCatalogService/ListProducts", svc.ListProducts)
}
