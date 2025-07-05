package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	conf := InitConfig()
	t.Log(conf)
}
