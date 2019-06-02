package service

import (
	"testing"

	"github.com/tangzixiang/goprojectinit/pkg/utils"
)

func init() {
	utils.Verbose = true
}

func Test_checkGOModules(t *testing.T) {
	t.Logf("checkGOModules(): %v\n", checkGOModules())
}
