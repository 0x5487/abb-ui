package config

type Configuration struct {
	Bind    string
	APIHost string      `yaml:"api_host"`
	Logs    []LogTarget `yaml:"logs"`
}

type LogTarget struct {
	Name             string
	Type             string
	Minlevel         string `yaml:"min_level"`
	ConnectionString string `yaml:"connection_string"`
}

func newConfiguration() *Configuration {
	return &Configuration{
		Bind: ":80",
	}
}

func (c *Configuration) isValid() error {
	return nil
}
