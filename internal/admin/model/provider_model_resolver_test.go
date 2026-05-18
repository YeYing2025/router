package model

import "testing"

func TestNormalizeProviderLookupCandidates(t *testing.T) {
	values := NormalizeProviderLookupCandidates(
		"",
		"gpt-5.4",
		"openai/gpt-5.4",
		"openai/gpt-5.4",
		"claude-opus-4-6",
	)
	want := []string{"gpt-5.4", "openai/gpt-5.4", "claude-opus-4-6"}
	if len(values) != len(want) {
		t.Fatalf("len(candidates)=%d, want=%d candidates=%v", len(values), len(want), values)
	}
	for i := range want {
		if values[i] != want[i] {
			t.Fatalf("candidates[%d]=%q, want=%q all=%v", i, values[i], want[i], values)
		}
	}
}

func TestResolveProviderFromModelMap(t *testing.T) {
	providerByModel := map[string]string{
		"gpt-5.4":      "openai",
		"claude-4.1":   "anthropic",
		"legacy-model": "custom",
	}
	if got := ResolveProviderFromModelMap(providerByModel, "openai/gpt-5.4"); got != "openai" {
		t.Fatalf("ResolveProviderFromModelMap(openai/gpt-5.4)=%q, want openai", got)
	}
	if got := ResolveProviderFromModelMap(providerByModel, "claude-4.1"); got != "anthropic" {
		t.Fatalf("ResolveProviderFromModelMap(claude-4.1)=%q, want anthropic", got)
	}
	if got := ResolveProviderFromModelMap(providerByModel, "unknown-model"); got != "" {
		t.Fatalf("ResolveProviderFromModelMap(unknown-model)=%q, want empty", got)
	}
	if got := ResolveProviderFromModelMap(providerByModel, "legacy-model"); got != "" {
		t.Fatalf("ResolveProviderFromModelMap(legacy-model)=%q, want empty", got)
	}
}
