package model

import "gorm.io/gorm"

func migrateGroupModelChannelsTableWithDB(tx *gorm.DB) error {
	if tx == nil {
		return nil
	}
	if tx.Migrator().HasTable("group_model_routes") && !tx.Migrator().HasTable(GroupModelChannelsTableName) {
		if err := tx.Migrator().RenameTable("group_model_routes", GroupModelChannelsTableName); err != nil {
			return err
		}
	}
	return tx.AutoMigrate(&GroupModelChannel{})
}

func migrateGroupModelRoutesTableWithDB(tx *gorm.DB) error {
	return migrateGroupModelChannelsTableWithDB(tx)
}
