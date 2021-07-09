package main

import (
	"fmt"
	"testing"
)

func TestAppConfigData(t *testing.T) {
	c := new(appConfigData)
	c.loadConfigData()
	switch {
	case c.dbHost == "":
		t.Error("Database host name not present")
	case c.dbPort == "":
		t.Error("Database port not present")
	case c.dbUser == "":
		t.Error("Database username not present")
	case c.dbPass == "":
		t.Error("Database user password not present")
	case c.dbName == "":
		t.Error("Database name not present")
	case c.host == "":
		t.Error("Public host name not present")
	case c.staticContent == "":
		t.Error("Static content path not present")
	}
	fmt.Println(c)
}
