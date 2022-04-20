package value

type NullableString struct {
	val string
}

func (v NullableString) IsNull() bool {
	return len(v.val) == 0
}

func (v NullableString) GetVal() string {
	return v.val
}
