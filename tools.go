package diwef

func nvl(a, b any) any {
	if a == nil || a == "" || a == 0 {
		return b
	} else {
		return a
	}
}
