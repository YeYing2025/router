package channel

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yeying-community/router/internal/admin/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetChannelProtocolsHidesVolcengineRealtime(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:"+t.Name()+"?mode=memory&cache=private"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&model.ChannelProtocolCatalog{}); err != nil {
		t.Fatalf("AutoMigrate: %v", err)
	}
	if err := db.Create(&[]model.ChannelProtocolCatalog{
		{Name: "volcengine", ProtocolID: 40, Label: "VolcEngine", Enabled: true, SortOrder: 1},
		{Name: "volcengine-realtime", ProtocolID: 49, Label: "VolcEngine Realtime", Enabled: true, SortOrder: 2},
	}).Error; err != nil {
		t.Fatalf("seed channel protocols: %v", err)
	}

	originalDB := model.DB
	model.DB = db
	t.Cleanup(func() {
		model.DB = originalDB
	})

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/admin/channel/protocols", nil)

	GetChannelProtocols(c)

	var resp struct {
		Success bool `json:"success"`
		Data    []struct {
			Value string `json:"value"`
		} `json:"data"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal response %q: %v", recorder.Body.String(), err)
	}
	if !resp.Success {
		t.Fatalf("expected success response, got body=%s", recorder.Body.String())
	}
	for _, item := range resp.Data {
		if item.Value == "volcengine-realtime" {
			t.Fatalf("volcengine-realtime should be hidden from admin protocol selector")
		}
	}
	if len(resp.Data) != 1 || resp.Data[0].Value != "volcengine" {
		t.Fatalf("unexpected protocol options: %#v", resp.Data)
	}
}
