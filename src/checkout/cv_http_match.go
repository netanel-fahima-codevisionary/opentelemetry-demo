// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package main

import "net/http"

// CodeVisionary discovers net/http HandleFunc, not gRPC; names match Coralogix spans.
type CheckoutService struct{}

func (s *CheckoutService) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func registerCVHTTPMatchRoutes() {
	svc := &CheckoutService{}
	http.HandleFunc("POST /oteldemo.CheckoutService/PlaceOrder", svc.PlaceOrder)
}
