package script

import (
	"github.com/spf13/cast"
)

type InputWrapper struct {
	input any
}

func Input(v any) *InputWrapper {
	return &InputWrapper{input: v}
}

func (i *InputWrapper) LimitOffset() (int, int) {
	limit := i.PageSize()
	current := i.Current()

	if current < 1 {
		current = 0
	}

	if limit < 1 {
		limit = 100
	}

	offset := (current - 1) * limit
	return limit, offset
}

func (i *InputWrapper) PageSize() int {
	return i.getPaginationInt("pageSize", 0)
}

func (i *InputWrapper) Current() int {
	return i.getPaginationInt("current", 0)
}

func (i *InputWrapper) getPaginationInt(key string, defaultValue int) int {
	if pagination := i.getPaginationMap(); pagination != nil {
		if value, exists := pagination[key]; exists {
			return cast.ToInt(value)
		}
	}
	return defaultValue
}

func (i *InputWrapper) getPaginationMap() map[string]any {
	if v, ok := i.input.(map[string]any); ok {
		if pagination, ok := v["pagination"]; ok {
			if paginationMap, ok := pagination.(map[string]any); ok {
				return paginationMap
			}
		}
	}
	return nil
}
