package test

import (
	"errors"
	"fmt"
	"os"
	"simpleBE/internal"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// run with: go test ./test -v

var testDB *gorm.DB

func TestMain(m *testing.M) {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		if os.Getenv("CI") != "" {
			fmt.Fprintln(os.Stderr, "TEST_DATABASE_URL not set (CI)")
			os.Exit(1)
		}
		// local: skip the whole package if not configured
		fmt.Println("Skipping DB tests: TEST_DATABASE_URL not set")
		os.Exit(0)
	}

	var err error
	testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if err := testDB.AutoMigrate(&internal.Message{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	code := m.Run()

	// optional: close the underlying connection
	if sqlDB, err := testDB.DB(); err == nil {
		_ = sqlDB.Close()
	}
	os.Exit(code)
}

func withTx(t *testing.T, fn func(tx *gorm.DB)) {
	t.Helper()
	tx := testDB.Begin()
	if tx.Error != nil {
		t.Fatalf("begin tx: %v", tx.Error)
	}
	defer func() {
		_ = tx.Rollback().Error
	}()
	fn(tx)
}

func TestMessageCRUD(t *testing.T) {
	withTx(t, func(tx *gorm.DB) {
		// Create
		m := internal.Message{Text: "hello"}
		if err := tx.Create(&m).Error; err != nil {
			t.Fatalf("create: %v", err)
		}
		if m.ID == 0 {
			t.Fatalf("expected auto-increment ID")
		}

		// Read
		var got internal.Message
		if err := tx.First(&got, m.ID).Error; err != nil {
			t.Fatalf("read: %v", err)
		}
		if got.Text != "hello" {
			t.Fatalf("expected 'hello', got %q", got.Text)
		}

		// Update
		if err := tx.Model(&got).Update("Text", "updated").Error; err != nil {
			t.Fatalf("update: %v", err)
		}
		var after internal.Message
		if err := tx.First(&after, m.ID).Error; err != nil {
			t.Fatalf("reselect: %v", err)
		}
		if after.Text != "updated" {
			t.Fatalf("expected 'updated', got %q", after.Text)
		}

		// Delete
		res := tx.Delete(&internal.Message{}, m.ID)
		if res.Error != nil {
			t.Fatalf("delete: %v", res.Error)
		}
		if res.RowsAffected != 1 {
			t.Fatalf("expected RowsAffected=1, got %d", res.RowsAffected)
		}

		err := tx.First(&internal.Message{}, m.ID).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("expected ErrRecordNotFound after delete, got %v", err)
		}
	})
}

func TestMessageNotNull(t *testing.T) {
	withTx(t, func(tx *gorm.DB) {
		err := tx.Table("messages").Create(map[string]any{
			"text": nil,
		}).Error
		if err == nil {
			t.Fatalf("expected NOT NULL")
		}
	})
}
