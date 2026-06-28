package channel

import "testing"

func TestNormalizeProtocolNameVolcengineRealtimeAlias(t *testing.T) {
	if got := NormalizeProtocolName("49"); got != "volcengine" {
		t.Fatalf("NormalizeProtocolName(49) = %q, want volcengine", got)
	}
}
