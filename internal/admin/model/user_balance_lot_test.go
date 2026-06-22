package model

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestResolveBalanceCreditExpiresAt(t *testing.T) {
	base := int64(1776073942) // 2026-04-13 15:12:22 UTC+8

	if got := ResolveBalanceCreditExpiresAt(base, 0); got != 0 {
		t.Fatalf("validity=0 should never expire, got %d", got)
	}

	wantThreeDays := base + 3*86400
	if got := ResolveBalanceCreditExpiresAt(base, 3); got != wantThreeDays {
		t.Fatalf("validity=3 expected %d, got %d", wantThreeDays, got)
	}
}

func TestEnsureUserBalanceLotAmountColumnsBackfillsLegacyYYCColumns(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:"+t.Name()+"?mode=memory&cache=private"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.Exec(`
		CREATE TABLE user_balance_lots (
			id char(36) PRIMARY KEY,
			user_id char(36) NOT NULL,
			source_type varchar(32) NOT NULL,
			source_id char(36) NOT NULL,
			total_yyc bigint NOT NULL DEFAULT 0,
			used_yyc bigint NOT NULL DEFAULT 0,
			remaining_yyc bigint NOT NULL DEFAULT 0,
			status varchar(16) NOT NULL DEFAULT 'active',
			granted_at bigint NOT NULL DEFAULT 0,
			expires_at bigint NOT NULL DEFAULT 0,
			expired_at bigint NOT NULL DEFAULT 0,
			created_at bigint,
			updated_at bigint
		)
	`).Error; err != nil {
		t.Fatalf("create legacy table: %v", err)
	}
	if err := db.Exec(`
		INSERT INTO user_balance_lots (
			id, user_id, source_type, source_id, total_yyc, used_yyc, remaining_yyc, status
		) VALUES (
			'lot-1', 'user-1', 'topup_order', 'source-1', 1000, 200, 800, 'active'
		)
	`).Error; err != nil {
		t.Fatalf("insert legacy row: %v", err)
	}

	if err := ensureUserBalanceLotAmountColumnsWithDB(db); err != nil {
		t.Fatalf("ensure columns: %v", err)
	}
	for _, column := range []string{"total_amount", "used_amount", "remaining_amount"} {
		if !db.Migrator().HasColumn(UserBalanceLotsTableName, column) {
			t.Fatalf("expected column %s to exist", column)
		}
	}

	row := UserBalanceLot{}
	if err := db.First(&row, "id = ?", "lot-1").Error; err != nil {
		t.Fatalf("load migrated row: %v", err)
	}
	if row.TotalAmount != 1000 {
		t.Fatalf("TotalAmount=%d, want 1000", row.TotalAmount)
	}
	if row.UsedAmount != 200 {
		t.Fatalf("UsedAmount=%d, want 200", row.UsedAmount)
	}
	if row.RemainingAmount != 800 {
		t.Fatalf("RemainingAmount=%d, want 800", row.RemainingAmount)
	}
}

func TestListUserBalanceLotsPageWithDBAllowsAllHistoricalStatuses(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:"+t.Name()+"?mode=memory&cache=private"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&UserBalanceLot{}); err != nil {
		t.Fatalf("AutoMigrate: %v", err)
	}
	rows := []UserBalanceLot{
		{
			Id:              "lot-active",
			UserID:          "user-1",
			SourceType:      UserBalanceLotSourceTopup,
			SourceID:        "source-active",
			TotalAmount:     100,
			RemainingAmount: 20,
			Status:          UserBalanceLotStatusActive,
			CreatedAt:       3,
		},
		{
			Id:          "lot-exhausted",
			UserID:      "user-1",
			SourceType:  UserBalanceLotSourceTopup,
			SourceID:    "source-exhausted",
			TotalAmount: 100,
			UsedAmount:  100,
			Status:      UserBalanceLotStatusExhaust,
			CreatedAt:   2,
		},
		{
			Id:          "lot-expired",
			UserID:      "user-1",
			SourceType:  UserBalanceLotSourceRedeem,
			SourceID:    "source-expired",
			TotalAmount: 50,
			UsedAmount:  50,
			Status:      UserBalanceLotStatusExpired,
			CreatedAt:   1,
		},
	}
	if err := db.Create(&rows).Error; err != nil {
		t.Fatalf("create rows: %v", err)
	}

	allRows, total, err := ListUserBalanceLotsPageWithDB(db, "user-1", "", "", 1, 20, false)
	if err != nil {
		t.Fatalf("list all rows: %v", err)
	}
	if total != 3 || len(allRows) != 3 {
		t.Fatalf("all status rows=%d total=%d, want 3/3", len(allRows), total)
	}

	expiredRows, expiredTotal, err := ListUserBalanceLotsPageWithDB(db, "user-1", "", UserBalanceLotStatusExpired, 1, 20, false)
	if err != nil {
		t.Fatalf("list expired rows: %v", err)
	}
	if expiredTotal != 1 || len(expiredRows) != 1 || expiredRows[0].Id != "lot-expired" {
		t.Fatalf("expired rows=%+v total=%d, want lot-expired/1", expiredRows, expiredTotal)
	}
}
