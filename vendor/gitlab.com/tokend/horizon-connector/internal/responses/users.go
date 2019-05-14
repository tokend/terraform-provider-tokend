package responses

import "gitlab.com/tokend/horizon-connector/internal/resources"

type Users struct {
	Data   []resources.User `json:"data"`
	Links  Links            `json:"links"`
	Errors []Error          `json:"errors"`
}
