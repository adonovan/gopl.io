/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package hasher

type HashAlgorithm int

const (
	Sha256 HashAlgorithm = iota
	Sha384
	Sha512
)
