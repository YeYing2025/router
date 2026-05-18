package anthropic

import (
	"testing"

	relaymodel "github.com/yeying-community/router/internal/relay/model"
)

func TestValidateMessagesRequest(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		wantErr string
	}{
		{
			name: "valid",
			raw: `{
				"model":"claude-sonnet-4-6",
				"messages":[{"role":"user","content":"hello"}],
				"max_tokens":128
			}`,
		},
		{
			name: "missing model",
			raw: `{
				"messages":[{"role":"user","content":"hello"}]
			}`,
			wantErr: "field model is required",
		},
		{
			name: "missing messages",
			raw: `{
				"model":"claude-sonnet-4-6"
			}`,
			wantErr: "field messages is required",
		},
		{
			name: "invalid max tokens",
			raw: `{
				"model":"claude-sonnet-4-6",
				"messages":[{"role":"user","content":"hello"}],
				"max_tokens":2147483647
			}`,
			wantErr: "max_tokens is invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMessagesRequest([]byte(tt.raw))
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("ValidateMessagesRequest returned error: %v", err)
				}
				return
			}
			if err == nil || err.Error() != tt.wantErr {
				t.Fatalf("ValidateMessagesRequest error = %v, want %q", err, tt.wantErr)
			}
		})
	}
}

func TestParseMessagesRequestMeta(t *testing.T) {
	raw := []byte(`{
		"model":"claude-sonnet-4-6",
		"messages":[{"role":"user","content":"hello"}],
		"max_tokens":256,
		"stream":true,
		"temperature":0.2
	}`)

	requestMeta, err := ParseMessagesRequestMeta(raw)
	if err != nil {
		t.Fatalf("ParseMessagesRequestMeta returned error: %v", err)
	}
	if requestMeta.Model != "claude-sonnet-4-6" {
		t.Fatalf("requestMeta.Model = %q, want %q", requestMeta.Model, "claude-sonnet-4-6")
	}
	if requestMeta.MaxTokens != 256 {
		t.Fatalf("requestMeta.MaxTokens = %d, want %d", requestMeta.MaxTokens, 256)
	}
	if !requestMeta.Stream {
		t.Fatalf("requestMeta.Stream = false, want true")
	}
}

func TestParseMessagesRequestToRelayRequest(t *testing.T) {
	raw := []byte(`{
		"model":"claude-sonnet-4-6",
		"system":[{"type":"text","text":"system prompt"}],
		"messages":[
			{"role":"user","content":"hello"},
			{"role":"assistant","content":[{"type":"text","text":"world"}]}
		],
		"max_tokens":256,
		"stream":true,
		"temperature":0.2,
		"top_p":0.9,
		"top_k":5,
		"stop_sequences":["done"],
		"tools":[
			{
				"name":"search",
				"description":"search docs",
				"input_schema":{
					"type":"object",
					"properties":{"query":{"type":"string"}},
					"required":["query"]
				}
			}
		],
		"tool_choice":{"type":"auto"}
	}`)

	request, err := ParseMessagesRequestToRelayRequest(raw)
	if err != nil {
		t.Fatalf("ParseMessagesRequestToRelayRequest returned error: %v", err)
	}
	if request.Model != "claude-sonnet-4-6" {
		t.Fatalf("request.Model = %q, want %q", request.Model, "claude-sonnet-4-6")
	}
	if request.MaxTokens != 256 {
		t.Fatalf("request.MaxTokens = %d, want %d", request.MaxTokens, 256)
	}
	if !request.Stream {
		t.Fatalf("request.Stream = false, want true")
	}
	if len(request.Messages) != 3 {
		t.Fatalf("len(request.Messages) = %d, want 3", len(request.Messages))
	}
	if request.Messages[0].Role != "system" {
		t.Fatalf("request.Messages[0].Role = %q, want system", request.Messages[0].Role)
	}
	if request.Messages[1].Role != "user" || request.Messages[1].StringContent() != "hello" {
		t.Fatalf("unexpected user message: %#v", request.Messages[1])
	}
	if request.Messages[2].Role != "assistant" {
		t.Fatalf("request.Messages[2].Role = %q, want assistant", request.Messages[2].Role)
	}
	parsed := request.Messages[2].ParseContent()
	if len(parsed) != 1 || parsed[0].Type != relaymodel.ContentTypeText || parsed[0].Text != "world" {
		t.Fatalf("unexpected assistant content: %#v", parsed)
	}
	if len(request.Tools) != 1 {
		t.Fatalf("len(request.Tools) = %d, want 1", len(request.Tools))
	}
	if request.Tools[0].Type != "function" || request.Tools[0].Function.Name != "search" {
		t.Fatalf("unexpected tool: %#v", request.Tools[0])
	}
	stop, ok := request.Stop.([]string)
	if !ok || len(stop) != 1 || stop[0] != "done" {
		t.Fatalf("unexpected stop sequences: %#v", request.Stop)
	}
}
