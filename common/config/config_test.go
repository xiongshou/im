package config

import (
	"testing"
	"fmt"
)

func TestConfig(t *testing.T) {
	Init("../../plato.yaml")
	fmt.Println(GetEndpointsForDiscovery())
	fmt.Println(GetTimeoutForDiscovery())
}