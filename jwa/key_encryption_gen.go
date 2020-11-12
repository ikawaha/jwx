// this file was auto-generated by internal/cmd/gentypes/main.go: DO NOT EDIT

package jwa

import (
	"fmt"

	"github.com/pkg/errors"
)

// KeyEncryptionAlgorithm represents the various encryption algorithms as described in https://tools.ietf.org/html/rfc7518#section-4.1
type KeyEncryptionAlgorithm string

// Supported values for KeyEncryptionAlgorithm
const (
	A128GCMKW          KeyEncryptionAlgorithm = "A128GCMKW"          // AES-GCM key wrap (128)
	A128KW             KeyEncryptionAlgorithm = "A128KW"             // AES key wrap (128)
	A192GCMKW          KeyEncryptionAlgorithm = "A192GCMKW"          // AES-GCM key wrap (192)
	A192KW             KeyEncryptionAlgorithm = "A192KW"             // AES key wrap (192)
	A256GCMKW          KeyEncryptionAlgorithm = "A256GCMKW"          // AES-GCM key wrap (256)
	A256KW             KeyEncryptionAlgorithm = "A256KW"             // AES key wrap (256)
	DIRECT             KeyEncryptionAlgorithm = "dir"                // Direct encryption
	ECDH_ES            KeyEncryptionAlgorithm = "ECDH-ES"            // ECDH-ES
	ECDH_ES_A128KW     KeyEncryptionAlgorithm = "ECDH-ES+A128KW"     // ECDH-ES + AES key wrap (128)
	ECDH_ES_A192KW     KeyEncryptionAlgorithm = "ECDH-ES+A192KW"     // ECDH-ES + AES key wrap (192)
	ECDH_ES_A256KW     KeyEncryptionAlgorithm = "ECDH-ES+A256KW"     // ECDH-ES + AES key wrap (256)
	PBES2_HS256_A128KW KeyEncryptionAlgorithm = "PBES2-HS256+A128KW" // PBES2 + HMAC-SHA256 + AES key wrap (128)
	PBES2_HS384_A192KW KeyEncryptionAlgorithm = "PBES2-HS384+A192KW" // PBES2 + HMAC-SHA384 + AES key wrap (192)
	PBES2_HS512_A256KW KeyEncryptionAlgorithm = "PBES2-HS512+A256KW" // PBES2 + HMAC-SHA512 + AES key wrap (256)
	RSA1_5             KeyEncryptionAlgorithm = "RSA1_5"             // RSA-PKCS1v1.5
	RSA_OAEP           KeyEncryptionAlgorithm = "RSA-OAEP"           // RSA-OAEP-SHA1
	RSA_OAEP_256       KeyEncryptionAlgorithm = "RSA-OAEP-256"       // RSA-OAEP-SHA256
)

// Accept is used when conversion from values given by
// outside sources (such as JSON payloads) is required
func (v *KeyEncryptionAlgorithm) Accept(value interface{}) error {
	var tmp KeyEncryptionAlgorithm
	if x, ok := value.(KeyEncryptionAlgorithm); ok {
		tmp = x
	} else {
		var s string
		switch x := value.(type) {
		case fmt.Stringer:
			s = x.String()
		case string:
			s = x
		default:
			return errors.Errorf(`invalid type for jwa.KeyEncryptionAlgorithm: %T`, value)
		}
		tmp = KeyEncryptionAlgorithm(s)
	}
	switch tmp {
	case A128GCMKW, A128KW, A192GCMKW, A192KW, A256GCMKW, A256KW, DIRECT, ECDH_ES, ECDH_ES_A128KW, ECDH_ES_A192KW, ECDH_ES_A256KW, PBES2_HS256_A128KW, PBES2_HS384_A192KW, PBES2_HS512_A256KW, RSA1_5, RSA_OAEP, RSA_OAEP_256:
	default:
		return errors.Errorf(`invalid jwa.KeyEncryptionAlgorithm value`)
	}

	*v = tmp
	return nil
}

// String returns the string representation of a KeyEncryptionAlgorithm
func (v KeyEncryptionAlgorithm) String() string {
	return string(v)
}

// IsSymmetric returns true if the algorithm is a symmetric type
func (v KeyEncryptionAlgorithm) IsSymmetric() bool {
	switch v {
	case A128GCMKW, A128KW, A192GCMKW, A192KW, A256GCMKW, A256KW, DIRECT:
		return true
	}
	return false
}
