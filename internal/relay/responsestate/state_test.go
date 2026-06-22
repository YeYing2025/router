package responsestate

import (
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestAnalyzeRequestBodyDetectsPreviousResponseAndToolOutput(t *testing.T) {
	raw := []byte(`{
		"model":"gpt-5.4",
		"previous_response_id":"resp_prev",
		"input":[
			{"type":"message","role":"user","content":"hello"},
			{"type":"function_call_output","call_id":"call_123","output":"ok"}
		]
	}`)

	state := AnalyzeRequestBody(raw)
	if state.PreviousResponseID != "resp_prev" {
		t.Fatalf("PreviousResponseID = %q, want resp_prev", state.PreviousResponseID)
	}
	if !state.HasToolOutput {
		t.Fatal("HasToolOutput = false, want true")
	}
	if !IsStatefulRequestBody(raw) {
		t.Fatal("IsStatefulRequestBody = false, want true")
	}
}

func TestStoreAndLookupRoute(t *testing.T) {
	ResetForTest()
	defer ResetForTest()

	StoreRoute(" resp_1 ", " channel-1 ")
	channelID, ok := LookupRoute("resp_1")
	if !ok {
		t.Fatal("LookupRoute ok = false, want true")
	}
	if channelID != "channel-1" {
		t.Fatalf("LookupRoute channel = %q, want channel-1", channelID)
	}
}

func TestLookupRouteExpires(t *testing.T) {
	ResetForTest()
	defer ResetForTest()

	now := time.Unix(100, 0)
	routeNow = func() time.Time { return now }
	routeTTL = time.Second

	StoreRoute("resp_1", "channel-1")
	now = now.Add(2 * time.Second)
	if channelID, ok := LookupRoute("resp_1"); ok {
		t.Fatalf("LookupRoute = (%q, true), want expired", channelID)
	}
}

func TestStoreAndLookupRouteUsesRedisWhenEnabled(t *testing.T) {
	ResetForTest()
	defer ResetForTest()

	redisStore := map[string]string{}
	redisRouteEnabledFunc = func() bool { return true }
	redisSetRouteFunc = func(key string, value string, expiration time.Duration) error {
		if key != "responses_route:resp_redis" {
			t.Fatalf("redis key = %q, want responses_route:resp_redis", key)
		}
		if expiration != defaultResponseRouteTTL {
			t.Fatalf("redis expiration = %v, want %v", expiration, defaultResponseRouteTTL)
		}
		redisStore[key] = value
		return nil
	}
	redisGetRouteFunc = func(key string) (string, error) {
		value, ok := redisStore[key]
		if !ok {
			return "", redis.Nil
		}
		return value, nil
	}

	StoreRoute(" resp_redis ", " channel-redis ")
	channelID, ok := LookupRoute("resp_redis")
	if !ok {
		t.Fatal("LookupRoute ok = false, want true")
	}
	if channelID != "channel-redis" {
		t.Fatalf("LookupRoute channel = %q, want channel-redis", channelID)
	}
	if _, ok := lookupMemoryRoute("resp_redis"); ok {
		t.Fatal("memory route should not be populated when Redis is enabled")
	}
}

func TestLookupRouteDoesNotFallbackToMemoryWhenRedisEnabled(t *testing.T) {
	ResetForTest()
	defer ResetForTest()

	redisRouteEnabledFunc = func() bool { return false }
	StoreRoute("resp_stateful", "channel-memory")
	redisRouteEnabledFunc = func() bool { return true }
	redisGetRouteFunc = func(key string) (string, error) {
		return "", redis.Nil
	}

	if channelID, ok := LookupRoute("resp_stateful"); ok {
		t.Fatalf("LookupRoute = (%q, true), want Redis miss without memory fallback", channelID)
	}
}

func TestLookupRouteDoesNotFallbackToMemoryOnRedisError(t *testing.T) {
	ResetForTest()
	defer ResetForTest()

	redisRouteEnabledFunc = func() bool { return false }
	StoreRoute("resp_stateful", "channel-memory")
	redisRouteEnabledFunc = func() bool { return true }
	redisGetRouteFunc = func(key string) (string, error) {
		return "", errors.New("redis unavailable")
	}

	if channelID, ok := LookupRoute("resp_stateful"); ok {
		t.Fatalf("LookupRoute = (%q, true), want Redis error without memory fallback", channelID)
	}
}

func TestExtractResponseID(t *testing.T) {
	if got := ExtractResponseID([]byte(`{"id":"resp_abc","object":"response"}`)); got != "resp_abc" {
		t.Fatalf("ExtractResponseID = %q, want resp_abc", got)
	}
}
