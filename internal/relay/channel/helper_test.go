package channel

import (
	"testing"

	"github.com/yeying-community/router/internal/relay/apitype"
)

func TestIsVolcengineRealtimeRequestSupportsCanonicalVolcengine(t *testing.T) {
	if !IsVolcengineRealtimeRequest(VolcEngine, "/v1/realtime") {
		t.Fatalf("expected canonical volcengine realtime request to be detected")
	}
	if !IsVolcengineRealtimeRequest(VolcEngine, "/v1/realtime/sessions") {
		t.Fatalf("expected canonical volcengine realtime session request to be detected")
	}
	if IsVolcengineRealtimeRequest(VolcEngine, "/v1/responses") {
		t.Fatalf("did not expect non-realtime request to be detected")
	}
}

func TestToAPITypeForRequestUsesVolcengineRealtimeAdaptor(t *testing.T) {
	if got := ToAPITypeForRequest(VolcEngine, "/v1/realtime"); got != apitype.VolcengineRealtime {
		t.Fatalf("ToAPITypeForRequest(volcengine, /v1/realtime) = %d, want %d", got, apitype.VolcengineRealtime)
	}
	if got := ToAPITypeForRequest(VolcEngine, "/v1/responses"); got != apitype.OpenAI {
		t.Fatalf("ToAPITypeForRequest(volcengine, /v1/responses) = %d, want %d", got, apitype.OpenAI)
	}
}
