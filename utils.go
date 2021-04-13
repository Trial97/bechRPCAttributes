package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	"os/exec"
	"time"

	"github.com/cgrates/cgrates/engine"
	"github.com/cgrates/cgrates/utils"
)

var (
	attr = &engine.ExternalAttributeProfile{
		Tenant:   "cgrates.org",
		ID:       "ATTR1",
		Contexts: []string{"*any"},
		Attributes: []*engine.ExternalAttribute{
			{
				Path:  utils.MetaReq + utils.NestingSep + "TestType",
				Value: "ConcurrentSessions",
			},
		},
		Weight: 20,
	}
	tntID = utils.TenantID{
		Tenant: "cgrates.org",
		ID:     "ATTR1",
	}
)

// Return reference towards the command started so we can stop it if necessary
func StartEngine(engineCMD, cfgPath, address string, waitEngine int) error {
	enginePath, err := exec.LookPath(engineCMD)
	if err != nil {
		return err
	}
	engine := exec.Command(enginePath, "-config_path", cfgPath)
	if err := engine.Start(); err != nil {
		return err
	}
	fib := utils.Fib()
	var connected bool
	for i := 0; i < 200; i++ {
		time.Sleep(time.Duration(fib()) * time.Millisecond)
		if _, err := jsonrpc.Dial("tcp", address); err == nil {
			connected = true
			break
		}
	}
	if !connected {
		return fmt.Errorf("engine did not open port <%s>", address)
	}
	time.Sleep(time.Duration(waitEngine) * time.Millisecond) // wait for rater to register all subsistems
	return nil
}
