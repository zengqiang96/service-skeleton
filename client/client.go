package client

import "github.com/smallnest/rpcx/client"

func NewClient(addr *string) client.XClient {
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("User", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	return xclient
}
