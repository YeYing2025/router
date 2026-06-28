package realtime

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yeying-community/router/internal/relay/adaptor/openai"
	"github.com/yeying-community/router/internal/relay/meta"
	relaymodel "github.com/yeying-community/router/internal/relay/model"
	"github.com/yeying-community/router/internal/relay/relaymode"
)

const (
	DefaultResourceID = "volc.speech.dialog"
	FixedAppKey       = "PlgvMymc7f3tQnJ6"
)

type Adaptor struct{}

func NormalizeBaseURL(raw string) string {
	baseURL := strings.TrimRight(strings.TrimSpace(raw), "/")
	lower := strings.ToLower(baseURL)
	if strings.HasSuffix(lower, "/api/v3/realtime/dialogue") {
		return baseURL[:len(baseURL)-len("/api/v3/realtime/dialogue")]
	}
	return baseURL
}

func NormalizeResourceID(raw string) string {
	resourceID := strings.TrimSpace(raw)
	if resourceID == "" {
		return DefaultResourceID
	}
	return resourceID
}

func ApplyRealtimeHeaders(header http.Header, appID string, accessKey string, resourceID string) {
	if header == nil {
		return
	}
	header.Del("Authorization")
	header.Del("api-key")
	header.Del("OpenAI-Beta")
	header.Set("X-Api-App-Key", FixedAppKey)
	if trimmed := strings.TrimSpace(appID); trimmed != "" {
		header.Set("X-Api-App-ID", trimmed)
	}
	if trimmed := strings.TrimSpace(accessKey); trimmed != "" {
		header.Set("X-Api-Access-Key", trimmed)
	}
	header.Set("X-Api-Resource-Id", NormalizeResourceID(resourceID))
}

func (a *Adaptor) Init(meta *meta.Meta) {}

func (a *Adaptor) GetRequestURL(meta *meta.Meta) (string, error) {
	if meta == nil {
		return "", fmt.Errorf("meta is nil")
	}
	if meta.Mode != relaymode.Realtime {
		return "", fmt.Errorf("unsupported relay mode %d for volcengine realtime", meta.Mode)
	}
	baseURL := NormalizeBaseURL(meta.BaseURL)
	return baseURL + "/api/v3/realtime/dialogue", nil
}

func (a *Adaptor) SetupRequestHeader(c *gin.Context, req *http.Request, meta *meta.Meta) error {
	if req == nil || meta == nil {
		return fmt.Errorf("request or meta is nil")
	}
	ApplyRealtimeHeaders(req.Header, meta.Config.AppID, meta.APIKey, meta.Config.ResourceID)
	return nil
}

func (a *Adaptor) ConvertRequest(c *gin.Context, relayMode int, request *relaymodel.GeneralOpenAIRequest) (any, error) {
	return nil, errors.New("volcengine realtime only supports websocket /v1/realtime")
}

func (a *Adaptor) ConvertImageRequest(request *relaymodel.ImageRequest) (any, error) {
	return nil, errors.New("volcengine realtime only supports websocket /v1/realtime")
}

func (a *Adaptor) DoRequest(c *gin.Context, meta *meta.Meta, requestBody io.Reader) (*http.Response, error) {
	return nil, errors.New("volcengine realtime only supports websocket /v1/realtime")
}

func (a *Adaptor) DoResponse(c *gin.Context, resp *http.Response, meta *meta.Meta) (*relaymodel.Usage, *relaymodel.ErrorWithStatusCode) {
	return nil, openai.ErrorWrapper(errors.New("volcengine realtime only supports websocket /v1/realtime"), "unsupported_realtime_http_request", http.StatusBadRequest)
}

func (a *Adaptor) GetModelList() []string {
	return []string{}
}

func (a *Adaptor) GetChannelName() string {
	return "volcengine"
}
