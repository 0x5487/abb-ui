package routers

import "github.com/jasonsoft/abb-ui/config"

var (
	_config *config.Configuration
)

func init() {
	_config = config.Config()
}
