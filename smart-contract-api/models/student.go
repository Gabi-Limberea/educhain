package models

import (
	"fmt"
)

type Student struct {
	ExternalID    string `json:"id"`
	Provider      int64  `json:"-"`
	WalletAddress string `json:"walletAddress"`
}

func (s *Student) ToSQL() string {
	return fmt.Sprintf("('%s', '%s', %d)", s.ExternalID, s.WalletAddress, s.Provider)
}
