package golibri

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
)

func Cry(keyCr string, st string) string {
	return hex.EncodeToString(encrypt([]byte(st), keyCr))

}

func Ucry(keyCr string, st string) string {
	sol := ""
	if st != "" {
		b, err := hex.DecodeString(st)
		if err != nil {
			sol = st
		} else {
			dsq := decrypt(b, keyCr)
			if dsq != nil {
				sol = string(dsq)
			} else {
				sol = st
			}

		}
	}
	return sol

}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		//panic(err.Error())

		return data
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		//panic(err.Error())
		return data
	}
	//fmt.Println("gcm:", gcm)
	nonceSize := gcm.NonceSize()
	//fmt.Println("nonSize:", nonceSize)
	if nonceSize > len(data) {
		return nil
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func EncryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}
