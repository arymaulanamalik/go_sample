package blowfish

import (
	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"

	//nolint:all
	bf "golang.org/x/crypto/blowfish"
)

func EncryptBlowfish(pt, key []byte) ([]byte, error) {
	block, err := bf.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Pad(pt)
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return ct, nil
}

func DecryptBlowfish(ct, key []byte) ([]byte, error) {
	block, err := bf.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Unpad(pt)
	if err != nil {
		panic(err.Error())
	}
	return pt, nil
}
