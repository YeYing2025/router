package model

import (
	"context"
	"strings"

	commonutils "github.com/yeying-community/router/common/utils"
)

type GroupModelChannel struct {
	Group         string `json:"group" gorm:"type:varchar(32);primaryKey;autoIncrement:false"`
	Model         string `json:"model" gorm:"primaryKey;autoIncrement:false"`
	ChannelId     string `json:"channel_id" gorm:"type:varchar(64);primaryKey;autoIncrement:false;index"`
	UpstreamModel string `json:"upstream_model" gorm:"type:varchar(255);default:'';index"`
	Provider      string `json:"provider" gorm:"type:varchar(128);default:'';index"`
	Priority      *int64 `json:"priority" gorm:"bigint;default:0;index"`
}

func (route GroupModelChannel) GetPriority() int64 {
	if route.Priority == nil {
		return 0
	}
	return *route.Priority
}

const (
	GroupModelChannelsTableName = "group_model_channels"
)

// GroupModelChannel is the static expansion table of group model to channel model
// mappings. It is derived from group_models, group_channels, and channel model
// configuration data. Admin flows must not treat this table as source-of-truth
// configuration, and temporary connectivity or test facts should not be written back into it.
func (GroupModelChannel) TableName() string {
	return GroupModelChannelsTableName
}

func GetRandomSatisfiedChannel(group string, model string, ignoreFirstPriority bool) (*Channel, error) {
	return mustGroupModelChannelRepo().GetRandomSatisfiedChannel(group, model, ignoreFirstPriority)
}

func ListSatisfiedChannels(group string, model string) ([]*Channel, error) {
	return mustGroupModelChannelRepo().ListSatisfiedChannels(group, model)
}

func (channel *Channel) AddGroupModelChannels() error {
	return mustGroupModelChannelRepo().AddGroupModelChannels(channel)
}

func (channel *Channel) DeleteGroupModelChannels() error {
	return mustGroupModelChannelRepo().DeleteGroupModelChannels(channel)
}

// UpdateGroupModelChannels updates static group-model channel rows of this channel.
// Make sure the channel is completed before calling this function.
func (channel *Channel) UpdateGroupModelChannels() error {
	return mustGroupModelChannelRepo().UpdateGroupModelChannels(channel)
}

func RefreshGroupModelChannelsByChannelStatus(channelId string, status bool) error {
	return mustGroupModelChannelRepo().RefreshGroupModelChannelsByChannelStatus(channelId, status)
}

// GetTopChannelByModel returns the highest-priority enabled channel for a given group+model.
// Order: priority desc, then channel_id asc (stable for UI usage).
func GetTopChannelByModel(group string, model string) (*Channel, error) {
	return mustGroupModelChannelRepo().GetTopChannelByModel(group, model)
}

func GetGroupModels(ctx context.Context, group string) ([]string, error) {
	return mustGroupModelChannelRepo().GetGroupModels(ctx, group)
}

func NormalizeGroupModelChannelUpstreamModel(modelName string, upstreamModel string) string {
	upstream := strings.TrimSpace(upstreamModel)
	if upstream != "" {
		return upstream
	}
	return strings.TrimSpace(modelName)
}

func NormalizeGroupModelChannelProvider(provider string) string {
	normalized := commonutils.NormalizeProvider(provider)
	if normalized == "custom" {
		return ""
	}
	return normalized
}

func normalizeGroupModelChannelRowsPreserveOrder(rows []GroupModelChannel) []GroupModelChannel {
	if len(rows) == 0 {
		return []GroupModelChannel{}
	}
	result := make([]GroupModelChannel, 0, len(rows))
	seen := make(map[string]struct{}, len(rows))
	for _, row := range rows {
		normalized := row
		normalized.Group = strings.TrimSpace(normalized.Group)
		normalized.Model = strings.TrimSpace(normalized.Model)
		normalized.ChannelId = strings.TrimSpace(normalized.ChannelId)
		normalized.UpstreamModel = NormalizeGroupModelChannelUpstreamModel(normalized.Model, normalized.UpstreamModel)
		normalized.Provider = NormalizeGroupModelChannelProvider(normalized.Provider)
		if normalized.Group == "" || normalized.Model == "" || normalized.ChannelId == "" {
			continue
		}
		key := normalized.Group + "::" + normalized.Model + "::" + normalized.ChannelId
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, normalized)
	}
	return result
}
