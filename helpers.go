package hasty

// String helper allows to quickly get a pointer to a string
func String(v string) *string { return &v }

// Bool helper allows to quickly get a pointer to a boolean
func Bool(v bool) *bool { return &v }
