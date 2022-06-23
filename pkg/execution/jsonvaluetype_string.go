// Code generated by "stringer -type=JSONValueType"; DO NOT EDIT.

package execution

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownValueType-0]
	_ = x[StringValueType-1]
	_ = x[IntegerValueType-2]
	_ = x[FloatValueType-3]
	_ = x[BooleanValueType-4]
}

const _JSONValueType_name = "UnknownValueTypeStringValueTypeIntegerValueTypeFloatValueTypeBooleanValueType"

var _JSONValueType_index = [...]uint8{0, 16, 31, 47, 61, 77}

func (i JSONValueType) String() string {
	if i < 0 || i >= JSONValueType(len(_JSONValueType_index)-1) {
		return "JSONValueType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _JSONValueType_name[_JSONValueType_index[i]:_JSONValueType_index[i+1]]
}
