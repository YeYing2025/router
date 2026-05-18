package model

import "testing"

func TestBuildGroupModelBindingItemsPreservesDisabledGroupModelState(t *testing.T) {
	priority := int64(3)
	rows := []GroupModelChannel{
		{
			Group:         "group-a",
			Model:         "gpt-5.1-codex-mini",
			ChannelId:     "channel-1",
			UpstreamModel: "gpt-5.1-codex-mini",
			Priority:      &priority,
		},
	}
	channelByID := map[string]Channel{
		"channel-1": {
			Id:       "channel-1",
			Name:     "channel-1",
			Protocol: "openai",
			Status:   ChannelStatusEnabled,
		},
	}
	enabledByModel := map[string]bool{
		"gpt-5.1-codex-mini": false,
	}

	items := buildGroupModelBindingItems(rows, channelByID, enabledByModel)
	if len(items) != 1 {
		t.Fatalf("len(items) = %d, want 1", len(items))
	}
	if items[0].Enabled == nil {
		t.Fatalf("items[0].Enabled is nil, want false pointer")
	}
	if *items[0].Enabled {
		t.Fatalf("items[0].Enabled = true, want false")
	}
}
