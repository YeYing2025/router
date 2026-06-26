package monitor

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yeying-community/router/common/config"
	"github.com/yeying-community/router/internal/relay/model"
)

func ShouldDisableChannel(err *model.Error, statusCode int) bool {
	if !config.AutomaticDisableChannelEnabled {
		return false
	}
	if err == nil {
		return false
	}
	if statusCode == http.StatusUnauthorized {
		return true
	}
	switch err.Type {
	case "insufficient_quota", "authentication_error", "permission_error", "forbidden":
		return true
	}
	code := strings.ToLower(strings.TrimSpace(fmt.Sprint(err.Code)))
	if code == "invalid_api_key" || code == "account_deactivated" || code == "1113" {
		return true
	}

	lowerMessage := strings.ToLower(err.Message)
	if strings.Contains(lowerMessage, "your access was terminated") ||
		strings.Contains(lowerMessage, "violation of our policies") ||
		strings.Contains(lowerMessage, "your credit balance is too low") ||
		strings.Contains(lowerMessage, "organization has been disabled") ||
		strings.Contains(lowerMessage, "credit") ||
		strings.Contains(lowerMessage, "balance") ||
		strings.Contains(lowerMessage, "permission denied") ||
		strings.Contains(lowerMessage, "organization has been restricted") || // groq
		strings.Contains(lowerMessage, "api key not valid") || // gemini
		strings.Contains(lowerMessage, "api key expired") || // gemini
		strings.Contains(lowerMessage, "已欠费") ||
		strings.Contains(lowerMessage, "余额不足") ||
		strings.Contains(lowerMessage, "无可用资源包") {
		return true
	}
	return false
}

func IsInsufficientBalanceError(err *model.Error, statusCode int) bool {
	if err == nil {
		return false
	}
	if statusCode == http.StatusPaymentRequired {
		return true
	}
	lowerType := strings.ToLower(strings.TrimSpace(err.Type))
	if lowerType == "insufficient_quota" || lowerType == "billing_error" {
		return true
	}
	code := strings.ToLower(strings.TrimSpace(fmt.Sprint(err.Code)))
	if code == "insufficient_quota" || code == "billing_hard_limit_reached" || code == "1113" {
		return true
	}
	lowerMessage := strings.ToLower(strings.TrimSpace(err.Message))
	return strings.Contains(lowerMessage, "your credit balance is too low") ||
		strings.Contains(lowerMessage, "credit") ||
		strings.Contains(lowerMessage, "balance") ||
		strings.Contains(lowerMessage, "已欠费") ||
		strings.Contains(lowerMessage, "余额不足") ||
		strings.Contains(lowerMessage, "无可用资源包")
}

func ShouldEnableChannel(err error, openAIErr *model.Error) bool {
	if !config.AutomaticEnableChannelEnabled {
		return false
	}
	if err != nil {
		return false
	}
	if openAIErr != nil {
		return false
	}
	return true
}
