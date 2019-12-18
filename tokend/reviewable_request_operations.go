package tokend

import (
	"context"
	"fmt"
	"gitlab.com/tokend/connectors/submit"
	"gitlab.com/tokend/go/xdr"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers"
	"gitlab.com/tokend/go/xdrbuild"
)

var reviewableRequestOperationsSchema = map[string]*schema.Schema{
	"security_type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"op_types": {
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}

func resourceReviewableRequestOperations() *schema.Resource {
	return &schema.Resource{
		Create: resourceReviewableRequestOperationsCreate,
		Update: resourceReviewableRequestOperationsUpdate,
		Read:   resourceReviewableRequestOperationsRead,
		Delete: resourceReviewableRequestOperationsDelete,
		Schema: reviewableRequestOperationsSchema,
	}
}

func resourceReviewableRequestOperationsCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)
	values := d.Get("op_types").([]interface{})

	opTypes := make([]xdr.Uint32, 0, len(values))
	for _, op := range values {
		opTypes = append(opTypes, xdr.Uint32(cast.ToUint32(op.(string))))
	}
	marshaledOpTypes := xdr.MustMarshalBase64(opTypes)

	key := fmt.Sprintf("reviewable_request_operations:%s", d.Get("security_type").(string))

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.PutKeyValue{
		Key:   key,
		Value: marshaledOpTypes,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	_, err = m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(submit.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
	}
	d.SetId(key)

	return nil
}

func resourceReviewableRequestOperationsUpdate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)
	values := d.Get("op_types").([]interface{})

	opTypes := make([]xdr.Uint32, 0, len(values))
	for _, op := range values {
		opTypes = append(opTypes, xdr.Uint32(cast.ToUint32(op.(string))))
	}
	marshaledOpTypes := xdr.MustMarshalBase64(opTypes)
	key := fmt.Sprintf("reviewable_request_operations:%s", d.Get("security_type").(string))
	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.PutKeyValue{
		Key:   key,
		Value: marshaledOpTypes,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	_, err = m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(submit.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
	}
	d.SetId(key)

	return nil
}

func resourceReviewableRequestOperationsRead(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)
	key := fmt.Sprintf("reviewable_request_operations:%s", d.Get("security_type").(string))

	rawValue, err := m.Connector.KeyValues().Value(key)
	if err != nil {
		return errors.Wrap(err, "failed to get value")
	}

	err = helpers.UpdateReviewableRequestOperations(d, rawValue)
	if err != nil {
		return errors.Wrap(err, "failed to update value")
	}

	return nil
}

func resourceReviewableRequestOperationsDelete(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	key := fmt.Sprintf("reviewable_request_operations:%s", d.Get("security_type").(string))

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.RemoveKeyValue{
		Key: key,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	_, err = m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(submit.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
	}
	return nil
}
