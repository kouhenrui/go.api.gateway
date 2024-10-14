package config

import (
	"strings"
	"testing"
)

func TestNewCasbinEnforcer(t *testing.T) {

	lp := []string{
		"p, admin, /api/v1/resource, GET",
		"p, admin, /api/v1/resource, POST",
		"p, user, /api/v1/resource, GET",
	}

	for _, s := range lp {
		ts := strings.Split(s, ", ")
		t.Log(ts[1])
	}

}
