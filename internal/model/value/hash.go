package value

import (
	"crypto/sha512"
	"encoding/hex"
)

type Hash NullableString

func HashFromString(str string) Hash {
	return Hash{str}
}

func HashString(str string) (Hash, error) {
	hash := sha512.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return Hash{}, err
	}

	return Hash{hex.EncodeToString(hash.Sum(nil))}, nil
}
