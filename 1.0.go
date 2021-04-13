package main

import (
	"sync"
	"testing"

	"github.com/cgrates/birpc"
	"github.com/cgrates/birpc/context"
	"github.com/cgrates/birpc/jsonrpc"
	"github.com/cgrates/cgrates/engine"
	"github.com/cgrates/cgrates/utils"
)

var (
	newRPC     *birpc.Client
	newCfgPath = "./1.0"
	addressNew = "127.0.0.1:4012"
	onceNew    sync.Once
)

func BenchmarkNew(b *testing.B) {
	onceNew.Do(func() {
		err := StartEngine("cgr-engine2", newCfgPath, addressNew, 500)
		if err != nil {
			b.Fatal(err)
		}
	})
	var err error
	if newRPC, err = jsonrpc.Dial(utils.TCP, addressNew); err != nil {
		b.Fatal(err)
	}
	var rpl string
	var a engine.ExternalAttributeProfile
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := newRPC.Call(ctx, "AdminSv1.SetAttributeProfile", attr, &rpl); err != nil {
			b.Fatal(err)
		}
		if err := newRPC.Call(ctx, "AdminSv1.GetAttributeProfile", tntID, &a); err != nil {
			b.Fatal(err)
		}
		if err := newRPC.Call(ctx, "AdminSv1.RemoveAttributeProfile", tntID, &rpl); err != nil {
			b.Fatal(err)
		}
	}
}
