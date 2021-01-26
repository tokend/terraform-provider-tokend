/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type SubmitTransactionBody struct {
	// base-64 encoded XDR representation of transaction itself
	Tx string `json:"tx"`
	// defines whether to wait for ingest of transaction or not
	WaitForIngest *bool `json:"wait_for_ingest,omitempty"`
}
