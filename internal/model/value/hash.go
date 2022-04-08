package value

import (
	"crypto/sha512"
	"hash"
)

type Hash struct {
	hash hash.Hash `json:"-"`
}

func (h *Hash) Equals(hash Hash) bool {
	return h.hash == hash.hash
}

func HashFromString(str string) (*Hash, error) {
	hash := sha512.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return nil, err
	}

	return &Hash{hash}, nil
}
