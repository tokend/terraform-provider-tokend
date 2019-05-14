package regources

// SystemStatistics holds system specific variables to verbose to be added to the root
// and not significant enough to have own resource
type SystemStatistics struct {
	// ExternalSystemPoolEntriesCount shows number of active entries per external system type
	ExternalSystemPoolEntriesCount map[string]uint64 `json:"external_system_pool_entries_count,omitempty"`
}
