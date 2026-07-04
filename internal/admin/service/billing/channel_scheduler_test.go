package billing

import (
	"testing"

	"github.com/yeying-community/router/internal/admin/model"
)

func TestShouldAutoRefreshChannelBillingSkipsInsufficientBalanceAutoDisabled(t *testing.T) {
	channel := &model.Channel{Id: "channel-1", Status: model.ChannelStatusAutoDisabled}

	if shouldAutoRefreshChannelBilling(channel) {
		t.Fatalf("insufficient-balance auto-disabled channel should not be auto-refreshed")
	}
}

func TestShouldAutoRefreshChannelBillingIncludesEnabled(t *testing.T) {
	channel := &model.Channel{Id: "channel-1", Status: model.ChannelStatusEnabled}

	if !shouldAutoRefreshChannelBilling(channel) {
		t.Fatalf("enabled channel should be auto-refreshed")
	}
}

func TestShouldAutoRefreshChannelBillingSkipsManualDisabled(t *testing.T) {
	channel := &model.Channel{Id: "channel-1", Status: model.ChannelStatusManuallyDisabled}

	if shouldAutoRefreshChannelBilling(channel) {
		t.Fatalf("manually disabled channel should not be auto-refreshed")
	}
}

func TestShouldAutoRefreshChannelBillingSkipsOtherAutoDisabled(t *testing.T) {
	channel := &model.Channel{Id: "channel-1", Status: model.ChannelStatusAutoDisabled}
	if shouldAutoRefreshChannelBilling(channel) {
		t.Fatalf("non-insufficient-balance auto-disabled channel should not be auto-refreshed")
	}
}
