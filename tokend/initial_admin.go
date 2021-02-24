package tokend

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func resourceInitialAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceInitialAdminCreate,
		Update: resourceInitialAdminUpdate,
		Read:   resourceInitialAdminRead,
		Delete: resourceInitialAdminDelete,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"job_function": {
				Type:     schema.TypeString,
				Required: true,
			},
			"job_title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type IdentityAttributes struct {
	Email       string   `json:"email"`
	Title       string   `json:"title"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	JobFunction string   `json:"job_function"`
	JobTitle    string   `json:"job_title"`
	RoleID      string   `json:"role"`
	Projects    []string `json:"projects"`
}

type IdentityCreationData struct {
	Attributes IdentityAttributes `json:"attributes"`
}

type CreateIdentityRequest struct {
	Data IdentityCreationData `json:"data"`
}

func resourceInitialAdminCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	email := d.Get("email").(string)
	title := d.Get("title").(string)
	firstName := d.Get("first_name").(string)
	lastName := d.Get("last_name").(string)
	jobFunction := d.Get("job_function").(string)
	jobTitle := d.Get("job_title").(string)
	roleID := d.Get("role_id").(string)
	projects := make([]string, 0)

	var request = CreateIdentityRequest{
		Data: IdentityCreationData{Attributes: IdentityAttributes{
			Email:       email,
			Title:       title,
			FirstName:   firstName,
			LastName:    lastName,
			JobFunction: jobFunction,
			JobTitle:    jobTitle,
			RoleID:      roleID,
			Projects:    projects,
		}},
	}

	code, body, err := m.Api.Client().PostJSON("identities", request)
	if err != nil {
		return errors.Wrap(err, "failed to send request to api")
	}

	if code == http.StatusConflict {
		return nil
	}

	if code != http.StatusCreated {
		return fmt.Errorf("failed to create initial admin, response: %s", string(body))
	}

	return nil
}

func resourceInitialAdminUpdate(d *schema.ResourceData, _m interface{}) (err error) {
	panic("update not implemented")
}

func resourceInitialAdminRead(d *schema.ResourceData, _m interface{}) (err error) {
	return nil
}

func resourceInitialAdminDelete(d *schema.ResourceData, _m interface{}) (err error) {
	panic("delete not implemented")
}
