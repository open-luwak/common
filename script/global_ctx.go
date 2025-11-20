package script

import (
	"github.com/spf13/cast"
)

type CtxData map[string]any

func (c CtxData) Roles() []string {
	return cast.ToStringSlice(c[DefaultRolesKey])
}

func (c CtxData) AddRoles(roles []string) {
	if v, ok := c[DefaultRolesKey]; ok {
		if vv, ok := v.([]string); ok {
			vv = append(vv, roles...)
		}
	}
	c[DefaultRolesKey] = roles
}
