package model

import (
	"fmt"
	"sort"
	"strings"

	"gorm.io/gorm"
)

type ChannelDisableImpactGroup struct {
	Group  string   `json:"group"`
	Models []string `json:"models,omitempty"`
}

type ChannelDisableImpact struct {
	ChannelID string                      `json:"channel_id"`
	Groups    []ChannelDisableImpactGroup `json:"groups"`
}

func (impact ChannelDisableImpact) InUse() bool {
	return len(impact.Groups) > 0
}

type ChannelDisableBlockedError struct {
	Impact ChannelDisableImpact
}

func (err *ChannelDisableBlockedError) Error() string {
	if err == nil {
		return ""
	}
	if !err.Impact.InUse() {
		return "渠道可以禁用"
	}
	parts := make([]string, 0, len(err.Impact.Groups))
	for _, item := range err.Impact.Groups {
		if len(item.Models) > 0 {
			parts = append(parts, fmt.Sprintf("%s(%s)", item.Group, strings.Join(item.Models, ", ")))
			continue
		}
		parts = append(parts, item.Group)
	}
	return fmt.Sprintf(
		"渠道仍被以下分组使用，不能直接禁用：%s。请先处理对应分组的渠道绑定或模型映射。",
		strings.Join(parts, "; "),
	)
}

func CollectChannelDisableImpactWithDB(db *gorm.DB, channelID string) (ChannelDisableImpact, error) {
	if db == nil {
		return ChannelDisableImpact{}, fmt.Errorf("database handle is nil")
	}
	normalizedChannelID := strings.TrimSpace(channelID)
	if normalizedChannelID == "" {
		return ChannelDisableImpact{}, fmt.Errorf("渠道 ID 不能为空")
	}

	groupRows := make([]GroupChannel, 0)
	if err := db.
		Select(`"group"`, "channel_id").
		Where("channel_id = ?", normalizedChannelID).
		Order(`"group" asc`).
		Find(&groupRows).Error; err != nil {
		return ChannelDisableImpact{}, err
	}

	rows := make([]GroupModelChannel, 0)
	if err := db.
		Select(`"group"`, "model", "channel_id").
		Where("channel_id = ?", normalizedChannelID).
		Order(`"group" asc, model asc`).
		Find(&rows).Error; err != nil {
		return ChannelDisableImpact{}, err
	}

	groupModelSet := make(map[string]map[string]struct{}, len(groupRows)+len(rows))
	for _, row := range groupRows {
		groupID := strings.TrimSpace(row.Group)
		if groupID == "" {
			continue
		}
		if _, ok := groupModelSet[groupID]; !ok {
			groupModelSet[groupID] = make(map[string]struct{})
		}
	}
	for _, row := range rows {
		groupID := strings.TrimSpace(row.Group)
		modelName := strings.TrimSpace(row.Model)
		if groupID == "" {
			continue
		}
		if _, ok := groupModelSet[groupID]; !ok {
			groupModelSet[groupID] = make(map[string]struct{})
		}
		if modelName != "" {
			groupModelSet[groupID][modelName] = struct{}{}
		}
	}

	groupIDs := make([]string, 0, len(groupModelSet))
	for groupID := range groupModelSet {
		groupIDs = append(groupIDs, groupID)
	}
	sort.Strings(groupIDs)

	groups := make([]ChannelDisableImpactGroup, 0, len(groupIDs))
	for _, groupID := range groupIDs {
		models := make([]string, 0, len(groupModelSet[groupID]))
		for modelName := range groupModelSet[groupID] {
			models = append(models, modelName)
		}
		sort.Strings(models)
		groups = append(groups, ChannelDisableImpactGroup{
			Group:  groupID,
			Models: models,
		})
	}

	return ChannelDisableImpact{
		ChannelID: normalizedChannelID,
		Groups:    groups,
	}, nil
}

func EnsureChannelCanBeManuallyDisabledWithDB(db *gorm.DB, channelID string) error {
	impact, err := CollectChannelDisableImpactWithDB(db, channelID)
	if err != nil {
		return err
	}
	if !impact.InUse() {
		return nil
	}
	return &ChannelDisableBlockedError{Impact: impact}
}
