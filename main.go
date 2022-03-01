package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

//Fungsi untuk Hash Payload atau password menjadi 32byte/ Karakter
func HashPayload(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
	//res := hex.EncodeToString(hasher.Sum(nil))
	//return res -> untuk melihat hasil hash dengan println
}
func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(HashPayload(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	res := []byte(base64.StdEncoding.EncodeToString(append([]byte("LinkAja_"), ciphertext...)))
	return res
}
func Request(req string) string {
	res := "Request"
	return res
}
func main() {
	//fmt.Println(HashPayload("Bima"))
	ciphertext := Encrypt([]byte(Request("")), "Bima")
	fmt.Println(string(ciphertext))
	// var data = "bima afdani is win"
	// var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	// fmt.Println("encoded:", encodedString)

}
