package wager

type Config struct {
	ProviderURI string `mapstructure:"providerURI"`
	Address     string `mapstructure:"wagerAddress"`
}
