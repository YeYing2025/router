package openai

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestImageHandlerWrapsNonJSONUpstreamError(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusBadGateway,
		Status:     "502 Bad Gateway",
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("<html>bad gateway</html>")),
	}

	err, usage := ImageHandler(nil, resp)
	if usage != nil {
		t.Fatalf("usage = %+v, want nil", usage)
	}
	if err == nil {
		t.Fatal("ImageHandler() error = nil, want upstream error")
	}
	if err.StatusCode != http.StatusBadGateway {
		t.Fatalf("StatusCode = %d, want %d", err.StatusCode, http.StatusBadGateway)
	}
	if err.Type != "upstream_error" {
		t.Fatalf("Type = %q, want upstream_error", err.Type)
	}
	if err.Code != "upstream_http_error" {
		t.Fatalf("Code = %v, want upstream_http_error", err.Code)
	}
	if err.Message != "upstream returned 502 Bad Gateway" {
		t.Fatalf("Message = %q", err.Message)
	}
}

func TestImageHandlerTreatsTwoHundredErrorEnvelopeAsUpstreamError(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Header:     make(http.Header),
		Body: io.NopCloser(strings.NewReader(`{
			"error": {
				"message": "upstream error: do request failed ( ) (cch_session_id: sess_mpvz9mcc_1096e6fee442)",
				"type": "service_unavailable_error",
				"code": "service_unavailable_error"
			}
		}`)),
	}

	err, usage := ImageHandler(nil, resp)
	if usage != nil {
		t.Fatalf("usage = %+v, want nil", usage)
	}
	if err == nil {
		t.Fatal("ImageHandler() error = nil, want upstream error")
	}
	if err.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("StatusCode = %d, want %d", err.StatusCode, http.StatusServiceUnavailable)
	}
	if err.Type != "service_unavailable_error" {
		t.Fatalf("Type = %q, want service_unavailable_error", err.Type)
	}
	if err.Code != "service_unavailable_error" {
		t.Fatalf("Code = %v, want service_unavailable_error", err.Code)
	}
	if !strings.Contains(err.Message, "sess_mpvz9mcc_1096e6fee442") {
		t.Fatalf("Message = %q, want session id", err.Message)
	}
}
