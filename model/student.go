package model

type Student struct {
	FirstName string `json:"first_name,omitempty"`
	ID        uint64 `json:"id,omitempty"`
}

