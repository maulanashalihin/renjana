package services

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2Params holds the parameters for argon2id password hashing.
type Argon2Params struct {
	Memory      uint32 // KiB (e.g. 64*1024 = 64MB)
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// DefaultArgon2Params returns safe defaults for argon2id password hashing.
// Memory: 64MB, Time: 3, Threads: 4, Salt: 16 bytes, Key: 32 bytes.
var DefaultArgon2Params = Argon2Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 4,
	SaltLength:  16,
	KeyLength:   32,
}

// generateFromPassword hashes a password using argon2id and returns the
// encoded hash in the standard format:
//
//	$argon2id$v=19$m=<memory>,t=<iterations>,p=<parallelism>$<base64-salt>$<base64-hash>
func generateFromPassword(password string, params Argon2Params) (string, error) {
	salt := make([]byte, params.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, params.Memory, params.Iterations, params.Parallelism, b64Salt, b64Hash)

	return encoded, nil
}

// checkPassword verifies a password against an argon2id-encoded hash.
func checkPassword(encodedHash, password string) bool {
	params, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false
	}

	otherHash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	return subtle.ConstantTimeCompare(hash, otherHash) == 1
}

// hashPassword hashes with default argon2id params (package-level convenience).
func hashPassword(password string) (string, error) {
	return generateFromPassword(password, DefaultArgon2Params)
}

// HashPassword is an exported wrapper around the package-level hashPassword,
// usable from other packages (e.g. handlers).
func HashPassword(password string) (string, error) {
	return hashPassword(password)
}

func decodeHash(encodedHash string) (params Argon2Params, salt, hash []byte, err error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return params, nil, nil, fmt.Errorf("invalid argon2 hash format")
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return params, nil, nil, fmt.Errorf("invalid argon2 version: %w", err)
	}
	if version != argon2.Version {
		return params, nil, nil, fmt.Errorf("unexpected argon2 version: %d", version)
	}

	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism); err != nil {
		return params, nil, nil, fmt.Errorf("invalid argon2 params: %w", err)
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(parts[4])
	if err != nil {
		return params, nil, nil, fmt.Errorf("invalid argon2 salt encoding: %w", err)
	}

	hash, err = base64.RawStdEncoding.Strict().DecodeString(parts[5])
	if err != nil {
		return params, nil, nil, fmt.Errorf("invalid argon2 hash encoding: %w", err)
	}

	params.SaltLength = uint32(len(salt))
	params.KeyLength = uint32(len(hash))

	return params, salt, hash, nil
}
