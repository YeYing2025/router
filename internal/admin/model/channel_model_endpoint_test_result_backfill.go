package model

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func BackfillChannelModelEndpointTestResultsFromChannelTestsWithDB(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("database handle is nil")
	}
	if err := db.AutoMigrate(&ChannelModelEndpointTestResult{}); err != nil {
		return err
	}

	channelIDs := make([]string, 0)
	if err := db.Model(&ChannelTest{}).
		Distinct("channel_id").
		Where("channel_id <> ''").
		Pluck("channel_id", &channelIDs).Error; err != nil {
		return err
	}
	channelIDs = normalizeTrimmedValuesPreserveOrder(channelIDs)
	for _, channelID := range channelIDs {
		normalizedChannelID := strings.TrimSpace(channelID)
		if normalizedChannelID == "" {
			continue
		}
		rows, err := ListLatestChannelTestsByChannelIDWithDB(db, normalizedChannelID)
		if err != nil {
			return err
		}
		if len(rows) == 0 {
			continue
		}
		if err := UpsertChannelModelEndpointTestResultsWithDB(db, normalizedChannelID, "", rows); err != nil {
			return err
		}
	}
	return nil
}
