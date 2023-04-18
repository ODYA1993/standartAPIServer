package api

type API struct {
	config *Config
}

func NewAPI(config *Config) *API {
	return &API{
		config: config,
	}
}

func (api *API) Start() error {
	return nil
}
