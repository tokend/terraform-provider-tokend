package user

import (
	"encoding/json"

	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/horizon-connector/internal/resources"
	"gitlab.com/tokend/horizon-connector/internal/responses"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"context"
	"fmt"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

// User obtains a single User by AccountID.
// If User doesn't exist - nil,nil is returned.
func (q *Q) User(accountID string) (*resources.User, error) {
	respBB, err := q.client.Get(fmt.Sprintf("/users/%s", accountID))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to send GET request")
	}

	if respBB == nil {
		// No such User
		return nil, nil
	}

	var response struct {
		Data resources.User `json:"data"`
	}
	if err := json.Unmarshal(respBB, &response); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response bytes", logan.F{
			"raw_response": string(respBB),
		})
	}

	return &response.Data, nil
}

// GetAccountID obtains accountID of User by emailAddress.
// If emailAddress doesn't exist - empty string, nil is returned.
func (q *Q) GetAccountID(emailAddress string) (string, error){
	requestURL := fmt.Sprintf("/user_id?email=%s", emailAddress)
	fields := logan.F{
		"request_url": requestURL,
	}

	respBB, err := q.client.Get(requestURL)
	if err != nil {
		return "", errors.Wrap(err, "Failed to get map with accountID", fields)
	}

	if respBB == nil {
		// No such emailAddress
		return "", nil
	}
	fields["raw_response"] = string(respBB)

	var response struct {
		AccountID string `json:"account_id"`
	}
	err = json.Unmarshal(respBB, &response)
	if err != nil {
		return "", errors.Wrap(err, "Failed to unmarshal response bytes", fields)
	}

	return response.AccountID, nil
}

// Users requests Users from Horizon using pagination
// and streams each User into returned Users channel.
//
// Once found empty page (no more Users) - closes both returned channels.
func (q *Q) Users(ctx context.Context) (<- chan resources.User, <- chan error) {
	userStream := make(chan resources.User)
	errChan := make(chan error)

	go func() {
		defer func() {
			close(userStream)
			close(errChan)
		}()

		url := "/users"
		for {
			select {
			case <-ctx.Done():
				return
			default:
				break
			}

			respBB, err := q.client.Get(url)
			if err != nil {
				errChan <- errors.Wrap(err, "Failed to get Users page")
				continue
			}

			var response responses.Users
			if err := json.Unmarshal(respBB, &response); err != nil {
				errChan <- errors.Wrap(err, "Failed to unmarshal response bytes", logan.F{
					"raw_response": string(respBB),
				})
				continue
			}

			if len(response.Data) == 0 {
				// No more users. Channels will be closed in defer.
				return
			}

			for _, user := range response.Data {
				ohigo := user

				ok := q.streamUser(ctx, ohigo, userStream)
				if !ok {
					// Ctx was canceled
					return
				}
			}

			url = response.Links.Next
		}
	}()


	return userStream, nil
}

func (q *Q) streamUser(ctx context.Context, user resources.User, userStream chan<- resources.User) bool {
	select {
	case <- ctx.Done():
		return false
	case userStream <- user:
		return true
	}
}
