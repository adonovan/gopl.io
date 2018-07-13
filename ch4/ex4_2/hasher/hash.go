/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package hasher

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

const fmtVerb = "%x" // base 16, with lower-case letters for a-f

// Hash returns the sha hash as a hex string, for the given algorithm and byte slice.
func Hash(algorithm HashAlgorithm, data []byte) string {
	switch algorithm {
	case Sha256:
		return fmt.Sprintf(fmtVerb, sha256.Sum256(data))
	case Sha384:
		return fmt.Sprintf(fmtVerb, sha512.Sum384(data))
	case Sha512:
		return fmt.Sprintf(fmtVerb, sha512.Sum512(data))
	default:
		return fmt.Sprintf(fmtVerb, sha256.Sum256(data))
	}
}
