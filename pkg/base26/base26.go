package base26

// base charset [a-z]
const base int64 = 26

// Encode encodes m to a string
func Encode(m int64, length int) string {
	bits := make([]int64, length)
	strs := make([]byte, length)
	
	for i := 0; m != 0 && i < length; i++ {
		res := m % base
		bits[i] = res
		m /= base
	}
	
	for i, v := range bits {
		strs[i] = byte(v) + 'a'
	}

	return string(strs)
}

// Decode decodes a string to a int64
func Decode(str string) int64 {
	var magnitude int64 = 1
	var origin int64

	for _, v := range str {
		origin += int64(v - 'a')*magnitude
		magnitude = magnitude*base
	}

	return origin
}
