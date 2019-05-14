package regources

import (
	"time"
)

// Contract represents singe contract entry with attached invoices
type Contract struct {
	ID                string                 `json:"id"`
	PT                string                 `json:"paging_token"`
	Contractor        string                 `json:"contractor"`
	Customer          string                 `json:"customer"`
	Escrow            string                 `json:"escrow"`
	StartTime         time.Time              `json:"start_time"`
	EndTime           time.Time              `json:"end_time"`
	InitialDetails    map[string]interface{} `json:"initial_details"`
	CustomerDetails   map[string]interface{} `json:"customer_details,omitempty"`
	AdditionalDetails []DetailsWithPayload   `json:"notes,omitempty"`
	Invoices          []ReviewableRequest    `json:"invoices,omitempty"`
	DisputeReason     *DetailsWithPayload    `json:"dispute,omitempty"`
	State             []Flag                 `json:"state"`
}

func (c Contract) PagingToken() string {
	return c.PT
}

type DetailsWithPayload struct {
	Details   map[string]interface{} `json:"details,omitempty"`
	Author    string                 `json:"author,omitempty"`
	CreatedAt time.Time              `json:"created_at,omitempty"`
}
