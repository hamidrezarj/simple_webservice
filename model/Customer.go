package model

import (
	"fmt"
)

// type Stringer interface {
// 	string()
// }

type Customer struct {
	ID           uint64 `json:"cID"`
	Name         string `json:"cName"`
	Tel          uint64 `json:"cTel"`
	Address      string `json:"cAddress"`
	RegisterDate string `json:"cRegisterDate"`
}

func (c Customer) String() string {
	return fmt.Sprintf("name:%s tel:%d id: %d address:%s", c.Name, c.Tel, c.ID, c.Address)
}
