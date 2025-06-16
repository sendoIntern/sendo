package response

import "be/pkg/pagination"

type APIResponse struct {
	Status     string            `json:"status"`
	Message    string            `json:"message,omitempty"`
	Pagination pagination.Paging `json:"pagination,omitempty"`
	Data       interface{}       `json:"data,omitempty"`
	Error      string            `json:"error,omitempty"`
}
