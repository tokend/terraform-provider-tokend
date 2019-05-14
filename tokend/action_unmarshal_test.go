package tokend

import (
	"encoding/json"
	"fmt"
	"testing"

	"gitlab.com/tokend/go/xdr"
)

func TestDecode(t *testing.T) {
	i := `AAAAAJmKrDlLkvckmYUVW/adHTiWug6ci+mHe1BNTF0xdX9ZAAAAAAAAAAAAAAAAAAAAAAAAAABcepENAAAAAAAAAAEAAAAAAAAACwAAAAAAAAAAAAAAAAAAAANBTEkAAAAAAJmKrDlLkvckmYUVW/adHTiWug6ci+mHe1BNTF0xdX9ZAAAAF0h26AAAAAAU9GsEAAAAAAAAAAB7eyJuYW1lIjoiQWxpY28iLCJsb2dvIjp7Im1pbWVfdHlwZSI6IiIsIm5hbWUiOiIiLCJrZXkiOiIiLCJ0eXBlIjoiIn0sInRlcm1zIjp7Im1pbWVfdHlwZSI6IiIsIm5hbWUiOiIiLCJrZXkiOiIiLCJ0eXBlIjoiIn19AAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAB3G0klgAAAED1BeM0QbD/tgQeRnG2aMiO1wrYTx4Sx0p53KsrqxMbdmRKsC5Z2+lal1eMsEGRw5HwkdmuewX/bYkt4iqTVhoI`
	var d xdr.TransactionEnvelope
	if err := xdr.SafeUnmarshalBase64(i, &d); err != nil {
		t.Fatal(err)
	}
	b, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}

func TestSOOOOOOOOOOOOOOOQA(t *testing.T) {
	r := xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeTransaction,
		Ext:  &xdr.EmptyExt{},
	}
	_, err := xdr.MarshalBase64(&r)
	if err != nil {
		t.Fatal(err)
	}
}
