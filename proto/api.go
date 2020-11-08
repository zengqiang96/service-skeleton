package proto

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type GetUserReq struct {
	Id string
}

type GetUserRsp struct {
	Id   string
	Name string
}
