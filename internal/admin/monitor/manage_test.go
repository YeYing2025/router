package monitor

import (
	"net/http"
	"testing"

	"github.com/yeying-community/router/common/config"
	relaymodel "github.com/yeying-community/router/internal/relay/model"
)

func TestShouldDisableChannelForZhipuInsufficientBalanceCode(t *testing.T) {
	previous := config.AutomaticDisableChannelEnabled
	config.AutomaticDisableChannelEnabled = true
	defer func() {
		config.AutomaticDisableChannelEnabled = previous
	}()

	err := &relaymodel.Error{
		Message: "余额不足或无可用资源包,请充值。",
		Code:    "1113",
	}

	if !ShouldDisableChannel(err, http.StatusTooManyRequests) {
		t.Fatalf("ShouldDisableChannel = false, want true")
	}
}

func TestIsInsufficientBalanceError(t *testing.T) {
	tests := []struct {
		name       string
		err        *relaymodel.Error
		statusCode int
		want       bool
	}{
		{
			name:       "payment required",
			err:        &relaymodel.Error{Message: "billing required"},
			statusCode: http.StatusPaymentRequired,
			want:       true,
		},
		{
			name:       "insufficient quota type",
			err:        &relaymodel.Error{Type: "insufficient_quota", Message: "quota exceeded"},
			statusCode: http.StatusTooManyRequests,
			want:       true,
		},
		{
			name:       "zhipu balance code",
			err:        &relaymodel.Error{Code: "1113", Message: "余额不足或无可用资源包,请充值。"},
			statusCode: http.StatusTooManyRequests,
			want:       true,
		},
		{
			name:       "permission error",
			err:        &relaymodel.Error{Type: "permission_error", Message: "permission denied"},
			statusCode: http.StatusForbidden,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInsufficientBalanceError(tt.err, tt.statusCode); got != tt.want {
				t.Fatalf("IsInsufficientBalanceError = %v, want %v", got, tt.want)
			}
		})
	}
}
