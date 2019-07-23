package tokend

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	"github.com/spf13/cast"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

	"gitlab.com/tokend/go/xdrbuild"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

const (
	externalSystemTypeStellarKey  = "external_system_type_stellar"
	externalSystemTypeEthereumKey = "external_system_type_ethereum"
)

func resourceExternalSystemPoolEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceExternalSystemPoolEntryCreate,
		Update: resourceExternalSystemPoolEntryUpdate,
		Read:   resourceExternalSystemPoolEntryRead,
		Delete: resourceExternalSystemPoolEntryDelete,
		Schema: map[string]*schema.Schema{
			"target_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"external_system_type": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data": {
				Type:      schema.TypeMap,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceExternalSystemPoolEntryCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	stellarExternalSystemType, err := m.Connector.KeyValues().Value(externalSystemTypeStellarKey)
	if err != nil {
		return errors.Wrap(err, "failed to get stellar external system type")
	}
	if stellarExternalSystemType == nil {
		return errors.New("stellar external system type key value not set")
	}

	ethereumExternalSystemType, err := m.Connector.KeyValues().Value(externalSystemTypeEthereumKey)
	if err != nil {
		return errors.Wrap(err, "failed to get ethereum external system type")
	}
	if ethereumExternalSystemType == nil {
		return errors.New("ethereum external system type key value not set")
	}

	externalType := d.Get("external_system_type").(int)

	var envelope string
	switch uint32(externalType) {
	case *stellarExternalSystemType.U32:
		envelope, err = stellarExternalTypeEnvelope(uint32(externalType), d, m)
	case *ethereumExternalSystemType.U32:
		//TODO implement
	default:
		panic(fmt.Sprintf("unknown external system type: %d", externalType))
	}

	if envelope == "" {
		return errors.New("empty transaction envelope")
	}

	result := m.Horizon.Submitter().Submit(context.TODO(), envelope)
	if result.Err != nil {
		return errors.Wrapf(result.Err, "failed to submit tx: %s %q", result.TXCode, result.OpCodes)
	}
	return nil
}

func resourceExternalSystemPoolEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_external_system_pool_entry update is not implemented")
}

func resourceExternalSystemPoolEntryRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceExternalSystemPoolEntryDelete(d *schema.ResourceData, _m interface{}) error {
	return errors.New("tokend_external_system_pool_entry delete is not implemented")
}

func stellarExternalTypeEnvelope(externalSystemType uint32, d *schema.ResourceData, m Meta) (string, error) {
	count := d.Get("target_count").(int)

	stellarData, err := helpers.StellarDataFromRaw(d.Get("data"))
	if err != nil {
		return "", errors.Wrap(err, "failed to get raw data")
	}

	if stellarData == nil {
		return "", errors.New("no data")
	}

	tx := m.Builder.Transaction(m.Signer)
	for i := 0; i < count; i++ {
		data := StellarAddressWithPayload{
			Type: "address_with_payload",
			Data: StellarData{
				Address: stellarData.Address,
				Payload: cast.ToString(generatePayload()),
			},
		}
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return "", errors.Wrap(err, "failed to marshal data")
		}
		tx = tx.Op(xdrbuild.CreateExternalPoolEntry(int32(externalSystemType), cast.ToString(dataBytes), 0))
	}
	envelope, err := tx.Sign(m.Signer).Marshal()
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal envelope")
	}
	return envelope, nil
}

func generatePayload() uint64 {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		panic("failed to generate uint64")
	}
	return val.Uint64()
}

//TODO move to regources
type StellarAddressWithPayload struct {
	Type string      `json:"type"`
	Data StellarData `json:"data"`
}
type StellarData struct {
	Address string `json:"address"`
	Payload string `json:"payload"`
}
