package value

import "fmt"

type Key []byte

func NewKey(key string) Key {
	if len(key) >= 32 {
		return Key(key[:32])
	}

	i := 0
	for len(key) < 32 {
		key = fmt.Sprintf("%s%s", key, string([]rune(key)[i]))
		i++
		if i > len(key) {
			i = 0
		}
	}

	return Key(key)
}
