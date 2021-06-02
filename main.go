package main

import (
	"fmt"
	// "git.bluebird.id/bbone/drvmgmt/be-template/transport"
	// "git.bluebird.id/logistic/commons/config"
	// "git.bluebird.id/logistic/commons/constant"
	// "report/transport"

	"report/transport"
)

func main() {
	recSrv := transport.NewGRPCServer()
	transport.GRPCServerRun(
		fmt.Sprintf(":%s", config.Get(constant.ServicePortKey, "8089")),
		recSrv,
	)
}
