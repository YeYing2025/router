package token

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yeying-community/router/common/ctxkey"
	"github.com/yeying-community/router/internal/admin/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTokenControllerTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file:"+t.Name()+"?mode=memory&cache=private"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&model.Token{}); err != nil {
		t.Fatalf("AutoMigrate: %v", err)
	}
	originalDB := model.DB
	model.DB = db
	t.Cleanup(func() {
		model.DB = originalDB
	})
	return db
}

func seedUserTokenForTest(t *testing.T, db *gorm.DB, token model.Token) {
	t.Helper()
	if err := db.Create(&token).Error; err != nil {
		t.Fatalf("create token: %v", err)
	}
}

func decodeTokenResponseBody(t *testing.T, body []byte) map[string]any {
	t.Helper()
	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal response %q: %v", string(body), err)
	}
	return payload
}

func TestGetAllTokensReturnsMaskedKeys(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := newTokenControllerTestDB(t)
	seedUserTokenForTest(t, db, model.Token{
		Id:          "token-1",
		UserId:      "user-1",
		Key:         "sk-secretTokenValue1234",
		Status:      model.TokenStatusEnabled,
		Name:        "alpha",
		CreatedTime: 100,
		UpdatedTime: 100,
	})

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Set(ctxkey.Id, "user-1")
	c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/public/token/?page=1", nil)

	GetAllTokens(c)

	payload := decodeTokenResponseBody(t, recorder.Body.Bytes())
	if success, _ := payload["success"].(bool); !success {
		t.Fatalf("expected success response, got %v", payload)
	}
	data, ok := payload["data"].([]any)
	if !ok || len(data) != 1 {
		t.Fatalf("data=%T %#v, want one token row", payload["data"], payload["data"])
	}
	row, ok := data[0].(map[string]any)
	if !ok {
		t.Fatalf("row=%T %#v, want object", data[0], data[0])
	}
	key, _ := row["key"].(string)
	if key != "sk-secr****1234" {
		t.Fatalf("key=%q, want masked value", key)
	}
}

func TestGetTokenReturnsMaskedKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := newTokenControllerTestDB(t)
	seedUserTokenForTest(t, db, model.Token{
		Id:          "token-1",
		UserId:      "user-1",
		Key:         "secretTokenValue1234",
		Status:      model.TokenStatusEnabled,
		Name:        "alpha",
		CreatedTime: 100,
		UpdatedTime: 100,
	})

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Set(ctxkey.Id, "user-1")
	c.Params = gin.Params{{Key: "id", Value: "token-1"}}
	c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/public/token/token-1", nil)

	GetToken(c)

	payload := decodeTokenResponseBody(t, recorder.Body.Bytes())
	if success, _ := payload["success"].(bool); !success {
		t.Fatalf("expected success response, got %v", payload)
	}
	row, ok := payload["data"].(map[string]any)
	if !ok {
		t.Fatalf("data=%T %#v, want object", payload["data"], payload["data"])
	}
	key, _ := row["key"].(string)
	if key != "secr****1234" {
		t.Fatalf("key=%q, want masked value", key)
	}
}

func TestAddTokenReturnsRawKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_ = newTokenControllerTestDB(t)

	body := map[string]any{
		"name":            "created-token",
		"remain_quota":    1000,
		"unlimited_quota": false,
	}
	payloadBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Set(ctxkey.Id, "user-1")
	c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/public/token/", bytes.NewReader(payloadBytes))
	c.Request.Header.Set("Content-Type", "application/json")

	AddToken(c)

	payload := decodeTokenResponseBody(t, recorder.Body.Bytes())
	if success, _ := payload["success"].(bool); !success {
		t.Fatalf("expected success response, got %v", payload)
	}
	row, ok := payload["data"].(map[string]any)
	if !ok {
		t.Fatalf("data=%T %#v, want object", payload["data"], payload["data"])
	}
	key, _ := row["key"].(string)
	if key == "" {
		t.Fatal("expected raw key on create response")
	}
	if strings.Contains(key, "****") {
		t.Fatalf("create response key=%q, should remain raw", key)
	}
}
