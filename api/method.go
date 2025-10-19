package api

import (
	"fmt"
	"strings"
)

type ParsedName struct {
	FullName     string `json:"fullName"` // FullName = OrgName + ApiName
	OrgName      string `json:"orgName"`
	ApiName      string `json:"apiName"` // ApiName = EntityName + EntityMethod
	EntityName   string `json:"entityName"`
	EntityMethod string `json:"entityMethod"`
}

func ParseMethodName(method string) (*ParsedName, error) {
	parts := strings.Split(method, ".")
	if len(parts) < 4 {
		return nil, fmt.Errorf("invalid method format '%s': expected at least 4 parts", method)
	}

	name := &ParsedName{
		FullName: method,
	}
	if len(parts) > 4 {
		name.OrgName = strings.Join(parts[:len(parts)-4], ".")
	}
	name.ApiName = strings.Join(parts[len(parts)-4:], ".")
	name.EntityName = strings.Join(parts[len(parts)-4:len(parts)-1], ".")
	name.EntityMethod = parts[len(parts)-1]

	return name, nil
}
