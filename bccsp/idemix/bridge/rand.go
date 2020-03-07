/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package bridge

import (
	cryptolib "TaiChainPKI/idemix"
	"github.com/hyperledger/fabric-amcl/amcl"
)

// NewRandOrPanic return a new amcl PRG or panic
func NewRandOrPanic() *amcl.RAND {
	rng, err := cryptolib.GetRand()
	if err != nil {
		panic(err)
	}
	return rng
}
