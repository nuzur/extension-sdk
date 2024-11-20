package domainhelpers

import (
	nemgen "github.com/nuzur/extension-sdk/v1/proto_deps/nem/idl/gen"
)

func EntityPrimaryKeys(entity *nemgen.Entity) []*nemgen.Field {
	res := []*nemgen.Field{}
	for _, f := range entity.Fields {
		if f.Key {
			res = append(res, f)
		}
	}
	return res
}