package sgw_test

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddytest"
)

func TestParams(t *testing.T) {
	tester := caddytest.NewTester(t)
	tester.InitServer(`
{
	order waf first
}

:9080 {
	waf {
		output stdout
		ipvendor "ipv4_en.ipdb"
		orders wafModule1 wafModule2 wafModule3
		strategyOrders strategy1 strategy2 strategy3
	}
	respond "Yahaha! You found me!"
}
  `, "caddyfile")

	tester.AssertGetResponse("http://localhost:9080", 200, "Yahaha! You found me!")
}
