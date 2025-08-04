package config

import (
	"os"
	"strings"
)

type Config struct {
	Server   ServerConfig
	Supabase SupabaseConfig
}

type ServerConfig struct {
	Port           string
	TrustedProxies []string
}

type SupabaseConfig struct {
	URL    string
	APIKey string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port:           getEnv("PORT", "8080"),
			TrustedProxies: getTrustedProxies(),
		},
		Supabase: SupabaseConfig{
			URL:    os.Getenv("SUPABASE_URL"),
			APIKey: os.Getenv("SUPABASE_API_KEY"),
		},
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getTrustedProxies() []string {
	proxies := os.Getenv("TRUSTED_PROXIES")
	if proxies == "" {
		return []string{}
	}

	proxyList := strings.Split(proxies, ",")
	for i, proxy := range proxyList {
		proxyList[i] = strings.TrimSpace(proxy)
	}

	return proxyList
}
