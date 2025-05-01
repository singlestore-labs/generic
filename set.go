package generic

func ToSet[T comparable](slice []T) map[T]struct{} {
	m := make(map[T]struct{}, len(slice))
	for _, item := range slice {
		m[item] = struct{}{}
	}
	return m
}
