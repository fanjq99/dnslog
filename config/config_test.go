package config

import "testing"

func TestParse(t *testing.T) {
	t.Log(Parse("../cmd/fixture/dev.yml"))
}
