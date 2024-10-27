package settings

// Consider integration of Logging and Monitoring services such as Sentry or Datadog

type (
	ThirdParty struct {
		GoogleMaps GoogleMaps `yaml:"google_maps"`
	}
	GoogleMaps struct {
		APIKey string `yaml:"api_key"`
	}
)
