// Code generated by "enumer -type=RegistryVisibility -json -trimprefix RegistryVisibility -transform lower"; DO NOT EDIT.

package puver

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _RegistryVisibilityName = "publicprivate"

var _RegistryVisibilityIndex = [...]uint8{0, 6, 13}

const _RegistryVisibilityLowerName = "publicprivate"

func (i RegistryVisibility) String() string {
	if i < 0 || i >= RegistryVisibility(len(_RegistryVisibilityIndex)-1) {
		return fmt.Sprintf("RegistryVisibility(%d)", i)
	}
	return _RegistryVisibilityName[_RegistryVisibilityIndex[i]:_RegistryVisibilityIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _RegistryVisibilityNoOp() {
	var x [1]struct{}
	_ = x[RegistryVisibilityPublic-(0)]
	_ = x[RegistryVisibilityPrivate-(1)]
}

var _RegistryVisibilityValues = []RegistryVisibility{RegistryVisibilityPublic, RegistryVisibilityPrivate}

var _RegistryVisibilityNameToValueMap = map[string]RegistryVisibility{
	_RegistryVisibilityName[0:6]:       RegistryVisibilityPublic,
	_RegistryVisibilityLowerName[0:6]:  RegistryVisibilityPublic,
	_RegistryVisibilityName[6:13]:      RegistryVisibilityPrivate,
	_RegistryVisibilityLowerName[6:13]: RegistryVisibilityPrivate,
}

var _RegistryVisibilityNames = []string{
	_RegistryVisibilityName[0:6],
	_RegistryVisibilityName[6:13],
}

// RegistryVisibilityString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RegistryVisibilityString(s string) (RegistryVisibility, error) {
	if val, ok := _RegistryVisibilityNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _RegistryVisibilityNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to RegistryVisibility values", s)
}

// RegistryVisibilityValues returns all values of the enum
func RegistryVisibilityValues() []RegistryVisibility {
	return _RegistryVisibilityValues
}

// RegistryVisibilityStrings returns a slice of all String values of the enum
func RegistryVisibilityStrings() []string {
	strs := make([]string, len(_RegistryVisibilityNames))
	copy(strs, _RegistryVisibilityNames)
	return strs
}

// IsARegistryVisibility returns "true" if the value is listed in the enum definition. "false" otherwise
func (i RegistryVisibility) IsARegistryVisibility() bool {
	for _, v := range _RegistryVisibilityValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for RegistryVisibility
func (i RegistryVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for RegistryVisibility
func (i *RegistryVisibility) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RegistryVisibility should be a string, got %s", data)
	}

	var err error
	*i, err = RegistryVisibilityString(s)
	return err
}
