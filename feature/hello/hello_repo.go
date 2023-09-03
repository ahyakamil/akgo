package hello

func helloRepo(helloReq HelloReq) Hello {
	return Hello{
		Name: helloReq.Name,
		Greeting: "Bonjour!",
		Secret: "secret",
	}
}
