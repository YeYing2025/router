package channel

import (
	"fmt"
	"strings"

	"github.com/yeying-community/router/internal/admin/model"
	channelsvc "github.com/yeying-community/router/internal/admin/service/channel"
)

func shouldProbeInsufficientBalanceRecovery(channel *model.Channel, state model.ChannelCircuitBreakerState) bool {
	if channel == nil || channel.Status != model.ChannelStatusAutoDisabled {
		return false
	}
	return model.IsInsufficientBalanceCircuitBreakerState(state)
}

func EnqueueInsufficientBalanceChannelRecoveryTests(limit int) (int, error) {
	if limit <= 0 {
		limit = 100
	}
	channels, err := channelsvc.GetAllBasic(0, 0, "all", true)
	if err != nil {
		return 0, err
	}
	channelByID := make(map[string]*model.Channel)
	channelIDs := make([]string, 0)
	for _, channelRow := range channels {
		if channelRow == nil || channelRow.Status != model.ChannelStatusAutoDisabled {
			continue
		}
		channelID := strings.TrimSpace(channelRow.Id)
		if channelID == "" {
			continue
		}
		channelByID[channelID] = channelRow
		channelIDs = append(channelIDs, channelID)
	}
	if len(channelIDs) == 0 {
		return 0, nil
	}
	states, err := model.ListChannelCircuitBreakerStatesByChannelIDsWithDB(model.DB, channelIDs)
	if err != nil {
		return 0, err
	}
	createdCount := 0
	for _, state := range states {
		channelID := strings.TrimSpace(state.ChannelId)
		channelRow := channelByID[channelID]
		if !shouldProbeInsufficientBalanceRecovery(channelRow, state) {
			continue
		}
		created, err := enqueueInsufficientBalanceRecoveryTest(channelRow, "")
		if err != nil {
			return createdCount, fmt.Errorf("enqueue recovery test for channel %s: %w", channelID, err)
		}
		if created {
			createdCount++
		}
		if createdCount >= limit {
			break
		}
	}
	return createdCount, nil
}
