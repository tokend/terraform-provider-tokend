package responses

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersUnmarshal(t *testing.T) {
	response := []byte(`{
	  "data": [
		{
		  "type": "not_verified",
		  "id": "GCGVW6JDUFPULEAZV2RXCDTIWBWWXMW5Q6VKMQJFVLMSELZOFMC4EW7Y",
		  "attributes": {
			"email": "random.email@example.com",
			"state": "nil",
			"kyc_sequence": 1,
			"reject_reason": "some_random_reject_reason",
			"recovery_address": "GCROYVH3STUDKB6FWQIGU2NNX4QVL2YWTATAGEMCC4TXF5QBI5QTFIPQ",
			"created_at": "2018-02-22T16:35:14.863787Z",
			"airdrop_state": "eligible"
		  }
		},
		{
		  "type": "not_verified",
		  "id": "GBVDUSYXOLPKWMRWTMFJXANGEIVDV4GUSYWSJONXTG3C7S53SX2YQH3R",
		  "attributes": {
			"email": "random.email2@example.com",
			"state": "nil",
			"kyc_sequence": 0,
			"reject_reason": "",
			"recovery_address": "GCMLKWQEM7I66LNLODEE4UACPKIGZLM6PWSWYDLHBYF7GDRDNMGMMBKK",
			"created_at": "2018-02-23T13:47:34.133953Z",
			"airdrop_state": "eligible"
		  }
		}
	  ],
	  "links": {
		"self": "/users?page=1",
		"next": "/users?page=2",
		"prev": "/users?"
	  }
	}`)

	var got Users
	if err := json.Unmarshal(response, &got); err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, "/users?page=1", got.Links.Self)
	assert.EqualValues(t, "/users?page=2", got.Links.Next)
	assert.EqualValues(t, "/users?", got.Links.Prev)

	assert.Len(t, got.Data, 2)
	user := got.Data[0]

	assert.EqualValues(t, "not_verified", user.Type)
	assert.EqualValues(t, "GCGVW6JDUFPULEAZV2RXCDTIWBWWXMW5Q6VKMQJFVLMSELZOFMC4EW7Y", user.ID)

	assert.EqualValues(t, "random.email@example.com", user.Attributes.Email)
	assert.EqualValues(t, "nil", user.Attributes.State)
	assert.EqualValues(t, 1, user.Attributes.KYCSequence)
	assert.EqualValues(t, "some_random_reject_reason", user.Attributes.RejectReason)
	assert.EqualValues(t, "GCROYVH3STUDKB6FWQIGU2NNX4QVL2YWTATAGEMCC4TXF5QBI5QTFIPQ", user.Attributes.RecoveryAddress)
	assert.EqualValues(t, "2018-02-22T16:35:14.863787Z", user.Attributes.CreatedAt)
	assert.EqualValues(t, "eligible", user.Attributes.AirdropState)
}
