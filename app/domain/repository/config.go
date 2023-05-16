package repository

import "github.com/americanas-go/config"

const (
	ProviderConfig = "app.domain.repository.attribute"
)

func init() {
	config.Add(ProviderConfig, "mock", "repository provider")
}

func ModuleProviderValue() string {
	return config.String((ProviderConfig))
}
