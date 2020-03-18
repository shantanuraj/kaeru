package hash

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type hashParams struct {
	saltLength  uint32
	iterations  uint32
	parallelism uint8
	memory      uint32
	keyLength   uint32
}

var params = hashParams{
	saltLength:  24,
	iterations:  3,
	parallelism: 2,
	memory:      64 * 1024,
	keyLength:   48,
}

func randomSalt(n uint32) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

// Password applies the argon2 hash to supplied password with random salt
func Password(password string) (passwordHash, encodedSalt string, err error) {
	p := params
	salt, err := randomSalt(p.saltLength)
	if err != nil {
		return "", "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Base64 encode the salt and hashed password.
	encodedSalt = base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	passwordHash = fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.memory,
		p.iterations,
		p.parallelism,
		encodedSalt,
		b64Hash,
	)

	return passwordHash, encodedSalt, nil
}
