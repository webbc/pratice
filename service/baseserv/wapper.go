package baseserv

var servManager *ServiceManager

func init()  {
	servManager = NewServiceManager()
}

func AddService(cfg *Config) bool {
	return servManager.AddService(cfg)
}
