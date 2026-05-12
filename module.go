package starkbankecdsa

import "github.com/k6io/k6/js/modules"

func init() {
	modules.Register("k6/x/starkbank", New())
}
