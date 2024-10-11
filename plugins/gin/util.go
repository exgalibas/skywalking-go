package gin

func limit(in []byte, limit uint) []byte {
	l := uint(len(in))
	if l <= limit {
		return in
	}
	return in[:limit]
}
