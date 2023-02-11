package discovery

import (
	// "fmt"
	"context"
	"testing"
	"time"
	"github.com/xiongshou/im/common/config"
)

func TestServiceDiscovery(t *testing.T) {
	config.Init("../../plato.yaml")

	ctx := context.Background()
	ser := NewServiceDiscovery(&ctx)

	defer ser.Close()
	ser.WatchService("/web/", func(key, value string) {}, func(key, value string) {})
	for {
		select {
		case <-time.Tick(10 * time.Second):
		}
	}
}