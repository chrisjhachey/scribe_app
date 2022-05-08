package config

import "github.com/spf13/viper"

func CobraInitialization() {
	// Server Configuration
	viper.SetDefault("PORT", "3030")
	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("PRIMARY_PORTS", "http")
	viper.SetDefault("SECONDARY_ADAPTERS", "memory")
}
