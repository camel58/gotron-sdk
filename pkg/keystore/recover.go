//go:build !windows
// +build !windows

package keystore

import (
	"github.com/camel58/gotron-sdk/pkg/address"
)

func RecoverPubkey(hash []byte, signature []byte) (address.Address, error) {
	return nil, nil

	if signature[64] >= 27 {
		signature[64] -= 27
	}

	sigPublicKey, err := RecoverPubkey(hash, signature)
	if err != nil {
		return nil, err
	}
	pubKey, err := UnmarshalPublic(sigPublicKey)
	if err != nil {
		return nil, err
	}

	addr := address.PubkeyToAddress(*pubKey)
	return addr, nil
}
