package baseserv

type CommonService struct {
}

func (this *CommonService) Init() bool {
	return true
}

func (this *CommonService) Start() bool {
	return true
}

func (this *CommonService) Update(now int64) {
}

func (this *CommonService) Stop() {
}

func (this *CommonService) Process(msg interface{}) bool {
	return true
}

func (this *CommonService) Panic(msg interface{}) bool {
	return true
}
