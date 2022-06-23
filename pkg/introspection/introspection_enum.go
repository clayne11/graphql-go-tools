// Code generated by go-enum
// DO NOT EDIT!

package introspection

import (
	"fmt"
)

const (
	// SCALAR is a __TypeKind of type SCALAR
	SCALAR __TypeKind = iota
	// LIST is a __TypeKind of type LIST
	LIST
	// NONNULL is a __TypeKind of type NON_NULL
	NONNULL
	// OBJECT is a __TypeKind of type OBJECT
	OBJECT
	// ENUM is a __TypeKind of type ENUM
	ENUM
	// INTERFACE is a __TypeKind of type INTERFACE
	INTERFACE
	// UNION is a __TypeKind of type UNION
	UNION
	// INPUTOBJECT is a __TypeKind of type INPUT_OBJECT
	INPUTOBJECT
)

const ___TypeKindName = "SCALARLISTNON_NULLOBJECTENUMINTERFACEUNIONINPUT_OBJECT"

var ___TypeKindMap = map[__TypeKind]string{
	0: ___TypeKindName[0:6],
	1: ___TypeKindName[6:10],
	2: ___TypeKindName[10:18],
	3: ___TypeKindName[18:24],
	4: ___TypeKindName[24:28],
	5: ___TypeKindName[28:37],
	6: ___TypeKindName[37:42],
	7: ___TypeKindName[42:54],
}

// String implements the Stringer interface.
func (x __TypeKind) String() string {
	if str, ok := ___TypeKindMap[x]; ok {
		return str
	}
	return fmt.Sprintf("__TypeKind(%d)", x)
}

var ___TypeKindValue = map[string]__TypeKind{
	___TypeKindName[0:6]:   0,
	___TypeKindName[6:10]:  1,
	___TypeKindName[10:18]: 2,
	___TypeKindName[18:24]: 3,
	___TypeKindName[24:28]: 4,
	___TypeKindName[28:37]: 5,
	___TypeKindName[37:42]: 6,
	___TypeKindName[42:54]: 7,
}

// Parse__TypeKind attempts to convert a string to a __TypeKind
func Parse__TypeKind(name string) (__TypeKind, error) {
	if x, ok := ___TypeKindValue[name]; ok {
		return x, nil
	}
	return __TypeKind(0), fmt.Errorf("%s is not a valid __TypeKind", name)
}

// MarshalText implements the text marshaller method
func (x *__TypeKind) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *__TypeKind) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := Parse__TypeKind(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
