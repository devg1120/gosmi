// Code generated by "enumer -type=Decl -autotrimprefix -json"; DO NOT EDIT

package types

import (
	"encoding/json"
	"fmt"
)

const (
	_Decl_name_0 = "UnknownImplicitTypeTypeAssignment"
	_Decl_name_1 = "ImplSequenceOfValueAssignmentObjectTypeObjectIdentityModuleIdentityNotificationTypeTrapTypeObjectGroupNotificationGroupModuleComplianceAgentCapabilitiesTextualConventionMacroComplGroupComplObjectImplObject"
	_Decl_name_2 = "ModuleExtensionTypedefNodeScalarTableRowColumnNotificationGroupComplianceIdentityClassAttributeEvent"
)

var (
	_Decl_index_0 = [...]uint8{0, 7, 19, 33}
	_Decl_index_1 = [...]uint8{0, 14, 29, 39, 53, 67, 83, 91, 102, 119, 135, 152, 169, 174, 184, 195, 205}
	_Decl_index_2 = [...]uint8{0, 6, 15, 22, 26, 32, 37, 40, 46, 58, 63, 73, 81, 86, 95, 100}
)

func (i Decl) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _Decl_name_0[_Decl_index_0[i]:_Decl_index_0[i+1]]
	case 4 <= i && i <= 19:
		i -= 4
		return _Decl_name_1[_Decl_index_1[i]:_Decl_index_1[i+1]]
	case 33 <= i && i <= 47:
		i -= 33
		return _Decl_name_2[_Decl_index_2[i]:_Decl_index_2[i+1]]
	default:
		return fmt.Sprintf("Decl(%d)", i)
	}
}

var _DeclNameToValue_map = map[string]Decl{
	_Decl_name_0[0:7]:     0,
	_Decl_name_0[7:19]:    1,
	_Decl_name_0[19:33]:   2,
	_Decl_name_1[0:14]:    4,
	_Decl_name_1[14:29]:   5,
	_Decl_name_1[29:39]:   6,
	_Decl_name_1[39:53]:   7,
	_Decl_name_1[53:67]:   8,
	_Decl_name_1[67:83]:   9,
	_Decl_name_1[83:91]:   10,
	_Decl_name_1[91:102]:  11,
	_Decl_name_1[102:119]: 12,
	_Decl_name_1[119:135]: 13,
	_Decl_name_1[135:152]: 14,
	_Decl_name_1[152:169]: 15,
	_Decl_name_1[169:174]: 16,
	_Decl_name_1[174:184]: 17,
	_Decl_name_1[184:195]: 18,
	_Decl_name_1[195:205]: 19,
	_Decl_name_2[0:6]:     33,
	_Decl_name_2[6:15]:    34,
	_Decl_name_2[15:22]:   35,
	_Decl_name_2[22:26]:   36,
	_Decl_name_2[26:32]:   37,
	_Decl_name_2[32:37]:   38,
	_Decl_name_2[37:40]:   39,
	_Decl_name_2[40:46]:   40,
	_Decl_name_2[46:58]:   41,
	_Decl_name_2[58:63]:   42,
	_Decl_name_2[63:73]:   43,
	_Decl_name_2[73:81]:   44,
	_Decl_name_2[81:86]:   45,
	_Decl_name_2[86:95]:   46,
	_Decl_name_2[95:100]:  47,
}

func DeclFromString(s string) (Decl, error) {
	if val, ok := _DeclNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Decl values", s)
}

func DeclAsList() []Decl {
	list := make([]Decl, len(_DeclNameToValue_map))
	idx := 0
	for _, v := range _DeclNameToValue_map {
		list[idx] = v
		idx++
	}
	return list
}

func DeclAsListString() []string {
	list := make([]string, len(_DeclNameToValue_map))
	idx := 0
	for k := range _DeclNameToValue_map {
		list[idx] = k
		idx++
	}
	return list
}

func DeclIsValid(t Decl) bool {
	for _, v := range DeclAsList() {
		if t == v {
			return true
		}
	}
	return false
}

func (i Decl) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *Decl) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Decl should be a string, got %s", data)
	}

	var err error
	*i, err = DeclFromString(s)
	return err
}