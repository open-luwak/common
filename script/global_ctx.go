package script

import (
	"github.com/spf13/cast"
)

type Ctx map[string]any

func (c Ctx) Roles() []string {
	return cast.ToStringSlice(c[DefaultRolesKey])
}

func (c Ctx) AddRoles(roles []string) {
	if v, ok := c[DefaultRolesKey]; ok {
		if vv, ok := v.([]string); ok {
			vv = append(vv, roles...)
		}
	}
	c[DefaultRolesKey] = roles
}
