package xor

func Bytes(a, b []byte) []byte {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if n == 0 {
		return nil
	}
	res := make([]byte, n)
	xorBytes(res, a, b)
	return res
}

func DstBytes(dst, a, b []byte) int {
	return xorBytes(dst, a, b)
}
