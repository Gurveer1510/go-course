package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func hashing() {
	password := "password123"
	// hashPassword := sha256.Sum256([]byte(password))

	// fmt.Println(password)
	// fmt.Println(hashPassword)
	// fmt.Printf("SHA-256 hash hex value: %x\n", hashPassword)
	// resultantHash := fmt.Sprintf("%x", hashPassword)
	// fmt.Println(resultantHash)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	signupHash := hashPassword(password, salt)

	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)
	fmt.Println("Hash:", signupHash)

	// Verify the password
	// retrieve the saltStr and decode it
	decodeStr, _ := base64.StdEncoding.DecodeString(saltStr)
	loginHash := hashPassword(password, decodeStr)

	// compare the stored signup hash with loginhash
	if signupHash == loginHash {
		fmt.Println("ACCESS GRANTED")
		return
	} 

	fmt.Println("ACCESS DENIED")


}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return string(base64.StdEncoding.EncodeToString(hash[:]))
}
