package api

import (
	"fmt"
	"strings"
)

type NameSourceType int

const (
	NameSourceUnknown NameSourceType = iota
	NameSourceHTTPJsonRPC
	NameSourceInternalDAO
	NameSourceCronJob
)

type ParsedName struct {
	NameSource   NameSourceType
	FullName     string `json:"fullName"` // FullName = OrgName + ApiName
	OrgName      string `json:"orgName"`
	ApiName      string `json:"apiName"` // ApiName = EntityName + EntityMethod
	EntityName   string `json:"entityName"`
	EntityMethod string `json:"entityMethod"`
}

func ParseMethodName(method string, ns NameSourceType) (*ParsedName, error) {
	parts := strings.Split(method, ".")
	if len(parts) < 4 {
		return nil, fmt.Errorf("invalid method format '%s': expected at least 4 parts", method)
	}

	name := &ParsedName{
		NameSource: ns,
		FullName:   method,
	}
	if len(parts) > 4 {
		name.OrgName = strings.Join(parts[:len(parts)-4], ".")
	}
	name.ApiName = strings.Join(parts[len(parts)-4:], ".")
	name.EntityName = strings.Join(parts[len(parts)-4:len(parts)-1], ".")
	name.EntityMethod = parts[len(parts)-1]

	return name, nil
}
