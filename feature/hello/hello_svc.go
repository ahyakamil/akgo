package hello

import "akgo/aklog"

func DoHello(helloReq HelloReq) HelloResp {
	aklog.Info("log to test mdc")
	dataFromRepo := helloRepo(helloReq)
	helloResp := HelloResp{
		Name:     dataFromRepo.Name,
		Greeting: dataFromRepo.Greeting,
	}
	return helloResp
}
