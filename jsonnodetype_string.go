// Code generated by "stringer -type=JSONNodeType"; DO NOT EDIT.

package jsongo

import "strconv"

const _JSONNodeType_name = "TypeUndefinedTypeMapTypeArrayTypeValuetypeError"

var _JSONNodeType_index = [...]uint8{0, 13, 20, 29, 38, 47}

func (i JSONNodeType) String() string {
	if i >= JSONNodeType(len(_JSONNodeType_index)-1) {
		return "JSONNodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _JSONNodeType_name[_JSONNodeType_index[i]:_JSONNodeType_index[i+1]]
}
