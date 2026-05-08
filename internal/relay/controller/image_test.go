package controller

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	relaymodel "github.com/yeying-community/router/internal/relay/model"
)

func TestGetImageRequestAppliesDefaults(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest("POST", "/v1/images/generations", strings.NewReader(`{"prompt":"draw a city skyline"}`))
	ctx.Request.Header.Set("Content-Type", "application/json")

	req, err := getImageRequest(ctx, 0)
	if err != nil {
		t.Fatalf("getImageRequest() error = %v", err)
	}
	if req.Prompt != "draw a city skyline" {
		t.Fatalf("Prompt = %q, want %q", req.Prompt, "draw a city skyline")
	}
	if req.N != 1 {
		t.Fatalf("N = %d, want %d", req.N, 1)
	}
	if req.Size != "1024x1024" {
		t.Fatalf("Size = %q, want %q", req.Size, "1024x1024")
	}
	if req.Model != "dall-e-2" {
		t.Fatalf("Model = %q, want %q", req.Model, "dall-e-2")
	}
}

func TestValidateImageRequest(t *testing.T) {
	tests := []struct {
		name       string
		request    *relaymodel.ImageRequest
		wantOK     bool
		wantErrMsg string
	}{
		{
			name: "valid request",
			request: &relaymodel.ImageRequest{
				Model:  "dall-e-3",
				Prompt: "draw a city skyline",
				Size:   "1024x1024",
				N:      1,
			},
			wantOK: true,
		},
		{
			name: "missing prompt",
			request: &relaymodel.ImageRequest{
				Model: "dall-e-3",
				Size:  "1024x1024",
				N:     1,
			},
			wantErrMsg: "prompt is required",
		},
		{
			name: "unsupported size",
			request: &relaymodel.ImageRequest{
				Model:  "dall-e-3",
				Prompt: "draw a city skyline",
				Size:   "512x512",
				N:      1,
			},
			wantErrMsg: "size not supported for this image model",
		},
		{
			name: "prompt too long",
			request: &relaymodel.ImageRequest{
				Model:  "dall-e-2",
				Prompt: strings.Repeat("a", 1001),
				Size:   "1024x1024",
				N:      1,
			},
			wantErrMsg: "prompt is too long",
		},
		{
			name: "invalid n",
			request: &relaymodel.ImageRequest{
				Model:  "dall-e-3",
				Prompt: "draw a city skyline",
				Size:   "1024x1024",
				N:      2,
			},
			wantErrMsg: "invalid value of n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateImageRequest(tt.request, nil)
			if tt.wantOK {
				if err != nil {
					t.Fatalf("validateImageRequest() error = %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("validateImageRequest() error = nil, want %q", tt.wantErrMsg)
			}
			if err.Error.Message != tt.wantErrMsg {
				t.Fatalf("validateImageRequest() message = %q, want %q", err.Error.Message, tt.wantErrMsg)
			}
		})
	}
}

func TestGetImageCostRatio(t *testing.T) {
	tests := []struct {
		name      string
		request   *relaymodel.ImageRequest
		wantRatio float64
		wantErr   bool
	}{
		{
			name: "dall-e-3 standard",
			request: &relaymodel.ImageRequest{
				Model: "dall-e-3",
				Size:  "1024x1792",
			},
			wantRatio: 2,
		},
		{
			name: "dall-e-3 hd square",
			request: &relaymodel.ImageRequest{
				Model:   "dall-e-3",
				Size:    "1024x1024",
				Quality: "hd",
			},
			wantRatio: 2,
		},
		{
			name: "dall-e-3 hd portrait",
			request: &relaymodel.ImageRequest{
				Model:   "dall-e-3",
				Size:    "1024x1792",
				Quality: "hd",
			},
			wantRatio: 3,
		},
		{
			name:    "nil request",
			request: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getImageCostRatio(tt.request)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("getImageCostRatio() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("getImageCostRatio() error = %v", err)
			}
			if got != tt.wantRatio {
				t.Fatalf("getImageCostRatio() = %v, want %v", got, tt.wantRatio)
			}
		})
	}
}
