package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
	"testing"

	"github.com/cgrates/cgrates/engine"
	"github.com/cgrates/cgrates/utils"
)

var (
	masterRPC     *rpc.Client
	masterCfgPath = "./master"
	addressMaster = "127.0.0.1:2012"
	onceMaster    sync.Once
)

func BenchmarkMaster(b *testing.B) {
	onceMaster.Do(func() {
		err := StartEngine("cgr-engine", masterCfgPath, addressMaster, 500)
		if err != nil {
			b.Fatal(err)
		}
	})
	var err error
	if masterRPC, err = jsonrpc.Dial(utils.TCP, addressMaster); err != nil {
		b.Fatal(err)
	}
	var rpl string
	var a engine.AttributeProfile

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := masterRPC.Call("APIerSv2.SetAttributeProfile", attr, &rpl); err != nil {
			b.Fatal(err)
		}
		if err := masterRPC.Call("APIerSv1.GetAttributeProfile", tntID, &a); err != nil {
			b.Fatal(err)
		}
		if err := masterRPC.Call("APIerSv1.RemoveAttributeProfile", tntID, &rpl); err != nil {
			b.Fatal(err)
		}
	}
}
