package source

import (

	"testing"
	"github.com/xiongshou/im/common/config"
)

func TestSource(t *testing.T) {
	config.Init("../../plato.yaml")
	Init()
}