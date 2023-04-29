package puver

import (
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
)

type GPGKey struct {
	Private string
	Public  string
}

func GenerateGPGKey(name string, email string, passphrase string) (GPGKey, error) {
	var k GPGKey

	key, err := helper.GenerateKey(name, email, []byte(passphrase), "rsa", 4096)
	if err != nil {
		return k, err
	}

	privateKeyRing, err := crypto.NewKeyFromArmored(key)
	if err != nil {
		return k, err
	}

	pub, err := privateKeyRing.GetArmoredPublicKey()
	if err != nil {
		return k, err
	}

	k.Private = key
	k.Public = pub

	return k, nil
}
