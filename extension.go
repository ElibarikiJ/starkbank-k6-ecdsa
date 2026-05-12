package starkbankecdsa

import (
	"strconv"
	"time"

	"github.com/k6io/k6/js/modules"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/ecdsa"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
)

func init() {
	modules.Register("k6/x/starkbank", New())
}

type Module struct{}

func New() *Module {
	return &Module{}
}

func (*Module) Sign(accessID, privateKeyPEM, body string) map[string]interface{} {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	message := accessID + ":" + timestamp + ":" + body

	privKey := privatekey.FromPem(privateKeyPEM)
	signature := ecdsa.Sign(message, &privKey)

	return map[string]interface{}{
		"timestamp": timestamp,
		"signature": signature.ToBase64(),
		"message":   message,
	}
}

// Helper to load private key from PEM once
func (*Module) GetPrivateKeyFromPEM(pem string) string {
	return pem // just pass through
}
