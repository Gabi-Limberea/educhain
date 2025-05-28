package models

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Provider struct {
	Email            string           `json:"email"`
	Password         string           `json:"password,omitempty"`
	ContractAddress  string           `json:"contractAddress"`
	OrganizationInfo OrganizationInfo `json:"organizationInfo"`
}

type OrganizationInfo struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Website     string `json:"website,omitempty"`
}

func (p *Provider) ToSQL() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf(`('%s'`, p.Email))
	strBuilder.WriteString(
		fmt.Sprintf(
			`, '%s'`, base64.StdEncoding.EncodeToString([]byte(p.Password)),
		),
	)
	strBuilder.WriteString(fmt.Sprintf(`, '%s'`, p.ContractAddress))
	strBuilder.WriteString(fmt.Sprintf(`, '%s'`, p.OrganizationInfo.Name))
	strBuilder.WriteString(fmt.Sprintf(`, '%s'`, p.OrganizationInfo.Address))
	strBuilder.WriteString(fmt.Sprintf(`, '%s'`, p.OrganizationInfo.PhoneNumber))
	strBuilder.WriteString(fmt.Sprintf(`, '%s'`, p.OrganizationInfo.Website))
	strBuilder.WriteString(")")
	return strBuilder.String()
}

func (p *Provider) IsValid() bool {
	return p.Email != "" && p.Password != "" && p.OrganizationInfo.Name != "" && p.OrganizationInfo.Address != ""
}
