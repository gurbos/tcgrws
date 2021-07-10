package main

import (
	"testing"
)

func TestStartupConfig(t *testing.T) {
	c := new(appConfigData)
	c.loadConfigData()
}
