package baseserv

type Config struct {
	Id       uint32
	TimeTick uint32
	ChanSize uint32
	Serv     IService
}
