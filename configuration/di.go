package configuration

type DI struct {
	config Config
}

func NewDI(cfg Config) *DI {
	di := &DI{config: cfg}
	return di
}

func (di *DI) GetConfig() Config {
	return di.config
}
