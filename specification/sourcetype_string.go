// Code generated by "stringer -type=SourceType"; DO NOT EDIT.

package specification

import "strconv"

const _SourceType_name = "SourceTypeFileSourceTypeText"

var _SourceType_index = [...]uint8{0, 14, 28}

func (i SourceType) String() string {
	if i < 0 || i >= SourceType(len(_SourceType_index)-1) {
		return "SourceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SourceType_name[_SourceType_index[i]:_SourceType_index[i+1]]
}