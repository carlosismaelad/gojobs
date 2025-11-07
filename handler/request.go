package handler

import "fmt"

// Create Opening
type CreateOpeningRequest struct {
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote *bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}

func errorParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("request body empty or malformed")
	}
	if r.Role == "" {
		return errorParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errorParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errorParamIsRequired("location", "string")
	}
	if r.Remote == nil {
		return errorParamIsRequired("remote", "bool")
	}
	if r.Link == "" {
		return errorParamIsRequired("link", "string")
	}
	if r.Salary <= 0  {
		return errorParamIsRequired("salary", "integer")
	}
	return nil
}
