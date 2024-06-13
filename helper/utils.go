package helper

func Ternary(condition bool, tru, fals any) any {
	if condition {
		return tru
	}
	return fals
}
