package main

import (
	"encoding/hex"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"log"
	"os"
	"io"
)
func main(){
	key := []byte("passphrasewhichneedstobe32bytes!")
	encrypt(key)
	decrypt(key)
}

func encrypt(key []byte) {
	log.Println("Encryption Program")
	text := []byte("My Super Secret Code Stuff")
	
	cpher, err := aes.NewCipher(key)
	if err!= nil {
		log.Fatal(err)
	}
	gcm,err:=cipher.NewGCM(cpher)
	if err!=nil{
		log.Fatal(err)
	}
	nonce:=make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err!= nil {
		log.Fatal(err)
	}
	log.Println("cipherHEX : ",hex.EncodeToString(gcm.Seal(nonce,nonce,text,nil)))
	err = os.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		log.Println(err)	
	}
}
func decrypt(key []byte){
	log.Println("Decryption Program")
	ciphertext,err:=os.ReadFile("myfile.data")
	if err != nil {
		log.Fatal(err)
	}
	c,err := aes.NewCipher(key)
	if err!= nil {
		log.Fatal(err)
	}
	gcm,err:=cipher.NewGCM(c)
	if err!=nil{
        log.Fatal(err)
    }
	nonceSize:=gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Fatal("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err!= nil {
		log.Fatal(err)
	}
	log.Println("plain : ",string(plaintext))
}