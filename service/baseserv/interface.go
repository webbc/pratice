package baseserv

type IService interface {
	Init() bool
	Update(int64)
	Start() bool
	Stop()
	Process(interface{}) bool
	Panic(interface{}) bool
}
