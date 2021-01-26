package regources

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"testing"
)

type testEntry struct {
	Key
	Info int
}

func TestIncludeCollection(t *testing.T) {
	var collection Included
	entries := make([]testEntry, 10)
	for i := range entries {
		entries[i] = testEntry{
			Key: Key{
				ID:   strconv.FormatInt(int64(i), 10),
				Type: ResourceType("test_entry"),
			},
			Info: rand.Int(),
		}

		collection.Add(&entries[i])
	}

	rawCollection, err := json.Marshal(collection)
	if err != nil {
		t.Fatal("failed to marshal collection", err)
	}

	var jsonCollection Included
	err = json.Unmarshal(rawCollection, &jsonCollection)
	if err != nil {
		t.Fatal("failed to unmarshal collection", err)
	}

	for _, entry := range entries {
		var actualEntry testEntry
		exists := jsonCollection.tryFindEntry(entry.GetKey(), &actualEntry)
		if !exists {
			t.Fatal("failed to find entry in collection")
		}

		if actualEntry != entry {
			t.Fatal(actualEntry, " actual entry does not match expected entry", entry)
		}
	}
}
