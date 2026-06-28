package channel

import (
	"strings"

	"github.com/yeying-community/router/internal/relay/apitype"
)

func ToAPIType(channelProtocol int) int {
	apiType := apitype.OpenAI
	switch channelProtocol {
	case Anthropic:
		apiType = apitype.Anthropic
	case Baidu:
		apiType = apitype.Baidu
	case PaLM:
		apiType = apitype.PaLM
	case Zhipu:
		apiType = apitype.Zhipu
	case Ali:
		apiType = apitype.Ali
	case Xunfei:
		apiType = apitype.Xunfei
	case AIProxyLibrary:
		apiType = apitype.AIProxyLibrary
	case Tencent:
		apiType = apitype.Tencent
	case Gemini:
		apiType = apitype.Gemini
	case Ollama:
		apiType = apitype.Ollama
	case AwsClaude:
		apiType = apitype.AwsClaude
	case Coze:
		apiType = apitype.Coze
	case Cohere:
		apiType = apitype.Cohere
	case Cloudflare:
		apiType = apitype.Cloudflare
	case DeepL:
		apiType = apitype.DeepL
	case VertextAI:
		apiType = apitype.VertexAI
	case Replicate:
		apiType = apitype.Replicate
	case Proxy:
		apiType = apitype.Proxy
	}

	return apiType
}

func IsVolcengineRealtimeRequest(channelProtocol int, requestPaths ...string) bool {
	switch channelProtocol {
	case VolcEngine:
		for _, path := range requestPaths {
			normalized := strings.TrimSpace(strings.ToLower(path))
			if strings.HasPrefix(normalized, "/v1/realtime") {
				return true
			}
		}
	}
	return false
}

func ToAPITypeForRequest(channelProtocol int, requestPaths ...string) int {
	if IsVolcengineRealtimeRequest(channelProtocol, requestPaths...) {
		return apitype.VolcengineRealtime
	}
	return ToAPIType(channelProtocol)
}

func ToAPITypeByProtocol(protocol string) int {
	return ToAPIType(TypeByProtocol(protocol))
}
