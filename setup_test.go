package caddywaf_test

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddytest"
)

func TestParams(t *testing.T) {
	tester := caddytest.NewTester(t)
	tester.InitServer(`
	:9080
	waf stdout {
		ipvendor  "xxx.ipdb"
		orders  wafModule1  wafModule2 wafModule3
		strategyOrders strategy1 strategy2 strategy3
	}
	respond "Yahaha! You found me!"
  `, "caddyfile")
}
