package lock

import (
	"fmt"
	"github.com/truechain/truechain-engineering-code/log"
)

const IfPrint = false

func PSLog(a ...interface{}) {
	if IfPrint {
		fmt.Println("[PSLog]", a)
		return
	}
	log.Debug("[PbftServer]", "[PSLog]", a)
}
