package model

import "testing"

func TestValidateProtocolConfiguration_DeepSeekRejectsBaseURLWithV1(t *testing.T) {
	baseURL := "https://api.deepseek.com/v1"
	channel := &Channel{
		Protocol: "deepseek",
		BaseURL:  &baseURL,
	}

	err := channel.ValidateProtocolConfiguration()
	if err == nil {
		t.Fatalf("expected validation error")
	}
	if got := err.Error(); got == "" {
		t.Fatalf("expected non-empty validation error")
	}
}

func TestValidateProtocolConfiguration_DeepSeekRejectsConfigAPIBaseURLWithV1(t *testing.T) {
	channel := &Channel{
		Protocol: "deepseek",
		Config:   `{"api_base_url":"https://api.deepseek.com/beta/v1"}`,
	}

	err := channel.ValidateProtocolConfiguration()
	if err == nil {
		t.Fatalf("expected validation error")
	}
	if got := err.Error(); got == "" {
		t.Fatalf("expected non-empty validation error")
	}
}

func TestValidateProtocolConfiguration_DeepSeekAcceptsRootAndBetaBaseURL(t *testing.T) {
	baseURL := "https://api.deepseek.com/beta"
	channel := &Channel{
		Protocol: "deepseek",
		BaseURL:  &baseURL,
		Config:   `{"api_base_url":"https://api.deepseek.com"}`,
	}

	if err := channel.ValidateProtocolConfiguration(); err != nil {
		t.Fatalf("expected validation success, got %v", err)
	}
}

func TestValidateProtocolConfiguration_VolcengineRealtimeEndpointRequiresAppID(t *testing.T) {
	channel := &Channel{
		Protocol: "volcengine",
		ChannelModels: []ChannelModel{
			{
				Model:     "speech-realtime-1",
				Selected:  true,
				Type:      ProviderModelTypeAudio,
				Endpoint:  ChannelModelEndpointRealtime,
				Endpoints: []string{ChannelModelEndpointRealtime},
			},
		},
		Config: `{"resource_id":"volc.speech.dialog"}`,
	}

	err := channel.ValidateProtocolConfiguration()
	if err == nil {
		t.Fatalf("expected validation error")
	}
}

func TestValidateProtocolConfiguration_VolcengineNonRealtimeDoesNotRequireAppID(t *testing.T) {
	channel := &Channel{
		Protocol: "volcengine",
		ChannelModels: []ChannelModel{
			{
				Model:    "doubao-seed-2-0-pro-260215",
				Selected: true,
				Type:     ProviderModelTypeText,
				Endpoint: ChannelModelEndpointResponses,
			},
		},
	}

	if err := channel.ValidateProtocolConfiguration(); err != nil {
		t.Fatalf("expected validation success, got %v", err)
	}
}
