// Copyright (c) Matthew Mellor
// SPDX-License-Identifier: MPL-2.0

package haproxy

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestBackendModelTimeoutsToAPI(t *testing.T) {
	model := BackendResourceModel{
		TimeoutConnect: types.StringValue("5s"),
		TimeoutCheck:   types.StringValue("5s"),
		TimeoutServer:  types.StringValue("10m"),
		CustomOptions:  types.StringValue("timeout tunnel 1h"),
	}

	request := model.toAPI(context.Background())
	if request.TimeoutConnect != "5s" || request.TimeoutCheck != "5s" || request.TimeoutServer != "10m" {
		t.Fatalf("unexpected timeout request values: %#v", request)
	}
	if request.CustomOptions != "timeout tunnel 1h" {
		t.Fatalf("unexpected custom options: %q", request.CustomOptions)
	}
}

func TestBackendModelTimeoutsFromAPI(t *testing.T) {
	response := backendAPIResponse{
		TimeoutConnect: "5s",
		TimeoutCheck:   "5s",
		TimeoutServer:  "10m",
		CustomOptions:  "timeout tunnel 1h",
	}
	var model BackendResourceModel
	model.fromAPI(context.Background(), &response, "backend-uuid")

	if model.TimeoutConnect.ValueString() != "5s" || model.TimeoutCheck.ValueString() != "5s" || model.TimeoutServer.ValueString() != "10m" {
		t.Fatalf("unexpected timeout model values: %#v", model)
	}
	if model.CustomOptions.ValueString() != "timeout tunnel 1h" {
		t.Fatalf("unexpected custom options: %q", model.CustomOptions.ValueString())
	}
}
