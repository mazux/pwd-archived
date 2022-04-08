package value

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type Password []byte

func (p Password) Decrypt(k Key) (string, error) {
	c, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(p) < nonceSize {
		return "", err
	}

	nonce, ciphertext := p[:nonceSize], p[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func Encrypt(password string, k Key) (Password, error) {
	c, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return Password(gcm.Seal(nonce, nonce, []byte(password), nil)), nil
}
