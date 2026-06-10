package model

import (
	"strings"
	"sync"

	"gorm.io/gorm"
)

var (
	groupBillingRatioLock       sync.RWMutex
	groupBillingRatioMap        = map[string]float64{}
	groupChannelBillingRatioMap = map[string]map[string]float64{}
)

func normalizeGroupBillingRatio(value float64) float64 {
	if value < 0 {
		return 1
	}
	return value
}

func buildGroupBillingRatioMap(rows []GroupCatalog) map[string]float64 {
	ratios := make(map[string]float64, len(rows))
	for _, row := range rows {
		groupID := strings.TrimSpace(row.Id)
		if groupID == "" {
			continue
		}
		ratios[groupID] = normalizeGroupBillingRatio(row.BillingRatio)
	}
	return ratios
}

func setGroupBillingRatioRuntime(ratios map[string]float64) {
	groupBillingRatioLock.Lock()
	groupBillingRatioMap = ratios
	groupBillingRatioLock.Unlock()
}

func setGroupChannelBillingRatioRuntime(ratios map[string]map[string]float64) {
	groupBillingRatioLock.Lock()
	groupChannelBillingRatioMap = ratios
	groupBillingRatioLock.Unlock()
}

func GetGroupBillingRatio(id string) float64 {
	groupID := strings.TrimSpace(id)
	if groupID == "" {
		return 1
	}
	groupBillingRatioLock.RLock()
	ratio, ok := groupBillingRatioMap[groupID]
	groupBillingRatioLock.RUnlock()
	if !ok {
		return 1
	}
	return normalizeGroupBillingRatio(ratio)
}

func GetGroupChannelBillingRatio(group string, channelID string) float64 {
	groupID := strings.TrimSpace(group)
	normalizedChannelID := strings.TrimSpace(channelID)
	if groupID == "" || normalizedChannelID == "" {
		return GetGroupBillingRatio(groupID)
	}
	groupBillingRatioLock.RLock()
	ratio, ok := groupChannelBillingRatioMap[groupID][normalizedChannelID]
	groupBillingRatioLock.RUnlock()
	if !ok {
		return GetGroupBillingRatio(groupID)
	}
	return normalizeGroupBillingRatio(ratio)
}

func buildGroupChannelBillingRatioMap(rows []GroupChannel) map[string]map[string]float64 {
	ratios := make(map[string]map[string]float64)
	for _, row := range rows {
		groupID := strings.TrimSpace(row.Group)
		channelID := strings.TrimSpace(row.ChannelId)
		if groupID == "" || channelID == "" || !row.Enabled {
			continue
		}
		if _, ok := ratios[groupID]; !ok {
			ratios[groupID] = make(map[string]float64)
		}
		ratios[groupID][channelID] = normalizeGroupBillingRatio(row.BillingRatio)
	}
	return ratios
}

func syncGroupBillingRatiosRuntimeWithDB(db *gorm.DB) error {
	if db == nil {
		return nil
	}
	rows, err := listGroupCatalogWithDB(db)
	if err != nil {
		return err
	}
	setGroupBillingRatioRuntime(buildGroupBillingRatioMap(rows))
	channelRows := make([]GroupChannel, 0)
	if err := db.Find(&channelRows).Error; err != nil {
		return err
	}
	setGroupChannelBillingRatioRuntime(buildGroupChannelBillingRatioMap(channelRows))
	return nil
}
