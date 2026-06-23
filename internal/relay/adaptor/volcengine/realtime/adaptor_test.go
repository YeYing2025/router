package realtime

import (
	"net/http"
	"testing"

	"github.com/yeying-community/router/internal/relay/meta"
	"github.com/yeying-community/router/internal/relay/relaymode"
)

func TestGetRequestURLUsesOfficialRealtimePath(t *testing.T) {
	adaptor := &Adaptor{}
	got, err := adaptor.GetRequestURL(&meta.Meta{
		Mode:    relaymode.Realtime,
		BaseURL: "https://openspeech.bytedance.com/",
	})
	if err != nil {
		t.Fatalf("GetRequestURL() error = %v", err)
	}
	want := "https://openspeech.bytedance.com/api/v3/realtime/dialogue"
	if got != want {
		t.Fatalf("GetRequestURL() = %q, want %q", got, want)
	}
}

func TestApplyRealtimeHeadersUsesOfficialHeaderNames(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"Bearer sk-test"},
		"OpenAI-Beta":   []string{"realtime=v1"},
	}
	ApplyRealtimeHeaders(header, "app-123", "access-456", "")
	if got := header.Get("Authorization"); got != "" {
		t.Fatalf("Authorization = %q, want empty", got)
	}
	if got := header.Get("OpenAI-Beta"); got != "" {
		t.Fatalf("OpenAI-Beta = %q, want empty", got)
	}
	if got := header.Get("X-Api-App-ID"); got != "app-123" {
		t.Fatalf("X-Api-App-ID = %q, want %q", got, "app-123")
	}
	if got := header.Get("X-Api-Access-Key"); got != "access-456" {
		t.Fatalf("X-Api-Access-Key = %q, want %q", got, "access-456")
	}
	if got := header.Get("X-Api-App-Key"); got != FixedAppKey {
		t.Fatalf("X-Api-App-Key = %q, want %q", got, FixedAppKey)
	}
	if got := header.Get("X-Api-Resource-Id"); got != DefaultResourceID {
		t.Fatalf("X-Api-Resource-Id = %q, want %q", got, DefaultResourceID)
	}
}
