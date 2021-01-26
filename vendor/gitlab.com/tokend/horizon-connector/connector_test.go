package horizon_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/horizon-connector"
	"gitlab.com/tokend/keypair"
)

func TestConnector_WithSigner(t *testing.T) {
	t.Run("new instance", func(t *testing.T) {
		base := &url.URL{}
		kp, _ := keypair.Random()
		were := horizon.NewConnector(base)
		got := were.WithSigner(kp)
		assert.NotEqual(t, got, were)
	})
}
