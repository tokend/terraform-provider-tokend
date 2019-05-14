package regources

import "time"

// Operation - represent operation
type Operation struct {
	Key
	Relationships OperationRelation `json:"relationships"`
	Attributes    OperationAttr     `json:"attributes"`
}

//OperationRelation - represents operation relationships
type OperationRelation struct {
	Tx      *Relation `json:"tx"`
	Source  *Relation `json:"source"`
	Details *Relation `json:"details"`
}

//OperationAttr - represents attributes of operation
type OperationAttr struct {
	AppliedAt time.Time `json:"applied_at"`
}
