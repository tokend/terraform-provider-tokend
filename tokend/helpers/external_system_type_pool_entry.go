package helpers

import (
	"crypto/rand"
	"encoding/json"
	"math"
	"math/big"

	"gitlab.com/tokend/keypair"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/derive"
	"gitlab.com/tokend/go/xdrbuild"
)

func EthereumExternalTypeEnvelopes(count int, externalSystemType uint32, d *schema.ResourceData,
	builder xdrbuild.Builder, signer keypair.Full) ([]string, error) {
	ethereumData, err := EthereumDataFromRaw(d.Get("data"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get raw data")
	}

	if ethereumData == nil {
		return nil, errors.New("no data")
	}

	deriver, err := derive.NewETHDeriver(ethereumData.XPub)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init deriver")
	}

	var envelopes []string
	for i := 1; i <= count; i++ {
		envelope, err := ethereumExternalTypeEnvelope(count, externalSystemType, deriver, builder, signer)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create transaction envelope")
		}
		envelopes = append(envelopes, envelope)
	}
	return envelopes, nil
}

func ethereumExternalTypeEnvelope(count int, externalSystemType uint32,
	deriver *derive.ETHDeriver, builder xdrbuild.Builder, signer keypair.Full) (string, error) {
	tx := builder.Transaction(signer)
	for i := 1; i <= count; i++ {
		address, err := deriver.ChildAddress(uint64(i))
		if err != nil {
			return "", errors.Wrap(err, "failed to derive child address")
		}

		data := EthereumAddress{
			Type: "address",
			Data: EthereumAddressData{
				Address: address,
			},
		}
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return "", errors.Wrap(err, "failed to marshal data")
		}
		tx = tx.Op(xdrbuild.CreateExternalPoolEntry(int32(externalSystemType), cast.ToString(dataBytes), 0))
	}
	envelope, err := tx.Sign(signer).Marshal()
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal envelope")
	}
	return envelope, nil
}

func StellarExternalTypeEnvelopes(count int, externalSystemType uint32,
	d *schema.ResourceData, builder xdrbuild.Builder, signer keypair.Full) ([]string, error) {
	stellarData, err := StellarDataFromRaw(d.Get("data"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get raw data")
	}

	if stellarData == nil {
		return nil, errors.New("no data")
	}

	var envelopes []string
	for i := 1; i <= count; i++ {
		envelope, err := stellarExternalTypeEnvelope(externalSystemType, stellarData, builder, signer)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create transaction envelope")
		}
		envelopes = append(envelopes, envelope)
	}
	return envelopes, nil
}

func stellarExternalTypeEnvelope(externalSystemType uint32, stellarData *StellarData,
	builder xdrbuild.Builder, signer keypair.Full) (string, error) {
	tx := builder.Transaction(signer)

	data := StellarAddressWithPayload{
		Type: "address_with_payload",
		Data: StellarAddressData{
			Address: stellarData.Address,
			Payload: cast.ToString(generatePayload()),
		},
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal data")
	}
	tx = tx.Op(xdrbuild.CreateExternalPoolEntry(int32(externalSystemType), cast.ToString(dataBytes), 0))

	envelope, err := tx.Sign(signer).Marshal()
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

type StellarAddressWithPayload struct {
	Type string             `json:"type"`
	Data StellarAddressData `json:"data"`
}
type StellarAddressData struct {
	Address string `json:"address"`
	Payload string `json:"payload"`
}

type EthereumAddress struct {
	Type string              `json:"type"`
	Data EthereumAddressData `json:"data"`
}

type EthereumAddressData struct {
	Address string `json:"address"`
}
