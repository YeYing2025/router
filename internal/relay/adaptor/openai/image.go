package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yeying-community/router/internal/relay/model"
)

func ImageHandler(c *gin.Context, resp *http.Response) (*model.ErrorWithStatusCode, *model.Usage) {
	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return ErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError), nil
	}
	err = resp.Body.Close()
	if err != nil {
		return ErrorWrapper(err, "close_response_body_failed", http.StatusInternalServerError), nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return wrapImageUpstreamError(resp, responseBody), nil
	}

	if errResponse := parseImageUpstreamError(responseBody); errResponse != nil {
		errResponse.StatusCode = statusCodeForImageUpstreamError(errResponse.Error)
		return errResponse, nil
	}

	var imageResponse ImageResponse
	err = json.Unmarshal(responseBody, &imageResponse)
	if err != nil {
		return ErrorWrapper(err, "unmarshal_response_body_failed", http.StatusInternalServerError), nil
	}

	resp.Body = io.NopCloser(bytes.NewBuffer(responseBody))

	for k, v := range resp.Header {
		c.Writer.Header().Set(k, v[0])
	}
	c.Writer.WriteHeader(resp.StatusCode)

	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		return ErrorWrapper(err, "copy_response_body_failed", http.StatusInternalServerError), nil
	}
	err = resp.Body.Close()
	if err != nil {
		return ErrorWrapper(err, "close_response_body_failed", http.StatusInternalServerError), nil
	}
	return nil, nil
}

func parseImageUpstreamError(responseBody []byte) *model.ErrorWithStatusCode {
	var errorResponse struct {
		Error model.Error `json:"error"`
	}
	if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
		return nil
	}
	if strings.TrimSpace(errorResponse.Error.Message) == "" {
		return nil
	}
	return &model.ErrorWithStatusCode{
		Error: errorResponse.Error,
	}
}

func statusCodeForImageUpstreamError(err model.Error) int {
	code := strings.ToLower(strings.TrimSpace(fmt.Sprint(err.Code)))
	errorType := strings.ToLower(strings.TrimSpace(err.Type))
	message := strings.ToLower(strings.TrimSpace(err.Message))
	combined := strings.Join([]string{code, errorType, message}, " ")

	switch {
	case strings.Contains(combined, "rate_limit"):
		return http.StatusTooManyRequests
	case strings.Contains(combined, "insufficient") ||
		strings.Contains(combined, "quota") ||
		strings.Contains(combined, "billing") ||
		strings.Contains(combined, "payment"):
		return http.StatusPaymentRequired
	case strings.Contains(combined, "unauthorized") ||
		strings.Contains(combined, "invalid_api_key"):
		return http.StatusUnauthorized
	case strings.Contains(combined, "forbidden"):
		return http.StatusForbidden
	case strings.Contains(combined, "invalid") ||
		strings.Contains(combined, "bad_request"):
		return http.StatusBadRequest
	case strings.Contains(combined, "service_unavailable") ||
		strings.Contains(combined, "no_available") ||
		strings.Contains(combined, "temporarily unavailable"):
		return http.StatusServiceUnavailable
	default:
		return http.StatusBadGateway
	}
}

func wrapImageUpstreamError(resp *http.Response, responseBody []byte) *model.ErrorWithStatusCode {
	if errResponse := parseImageUpstreamError(responseBody); errResponse != nil {
		errResponse.StatusCode = resp.StatusCode
		return errResponse
	}
	return &model.ErrorWithStatusCode{
		Error: model.Error{
			Message: fmt.Sprintf("upstream returned %s", resp.Status),
			Type:    "upstream_error",
			Code:    "upstream_http_error",
		},
		StatusCode: resp.StatusCode,
	}
}
