package signcontrol

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/keypair"
)

const (
	seed = "SCDMOOXVNMO6SA22AYUMZDIGLDJMBUTVEGB73FFNTLFJILBJWIU4NQ3D"
)

func TestSignRequest(t *testing.T) {
	t.Run("example", func(t *testing.T) {

		kp, _ := keypair.Parse(seed)
		request, _ := http.NewRequest("GET", "http://example.com/fo?ob=ar", nil)
		request.Header.Set("date", "Sun, 05 Jan 2014 21:31:40 GMT")
		err := SignRequest(request, kp)
		assert.NoError(t, err)
		assert.Equal(t,
			fmt.Sprintf(`keyId="%s",algorithm="%s",signature="zSkfg3R002fE1BKny3lAIqTibDzS+5u63vLjx75mc3B4ylO0Sxd4NlcLQpz4iXTRDw6oHuZZO69xrOoKIqZpAw==",headers="date real-request-target"`, kp.Address(), SignatureAlgorithm),
			request.Header.Get("signature"),
		)
	})

	t.Run("substitute", func(t *testing.T) {
		normalRequest, _ := http.NewRequest("GET", "https://foo.bar/baz", nil)

		sneakyRequest, _ := http.NewRequest("GET", "https://foo.bar/boogie_woogie", nil)
		sneakyRequest.Header.Add("Real-Request-Target", "get /baz")

		kp, _ := keypair.Random()

		err := SignRequest(normalRequest, kp)
		assert.NoError(t, err)

		err = SignRequest(sneakyRequest, kp)
		assert.NoError(t, err)

		assert.Equal(t, sneakyRequest.Header.Get("signature"), normalRequest.Header.Get("signature"))

		_, err = checkV2(normalRequest)
		assert.NoError(t, err)
		_, err = checkV2(sneakyRequest)
		assert.NoError(t, err)
	})
}

func TestCheckDate(t *testing.T) {
	t.Run("expired", func(t *testing.T) {
		assert.Equal(t, ErrExpired, checkDate("Tue, 15 Nov 1994 08:12:31 GMT"))
	})

	t.Run("fresh", func(t *testing.T) {
		assert.NoError(t, checkDate(time.Now().UTC().Format(http.TimeFormat)))
	})
}
