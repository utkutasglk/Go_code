package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)


func main() {

	pass := "password123"

	// hash := sha256.Sum256([]byte(pass))
	// hash256 := sha512.Sum512([]byte(pass))


	// fmt.Println(pass)
	// fmt.Println(hash)
	// fmt.Println(hash256)

	// fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	// fmt.Printf("SHA-512 Hash hex val: %x\n", hash256)

	salt, err := generateSalt()
	fmt.Println("Original salt :", salt)
	fmt.Printf("Original salt : %x\n", salt)

	if err != nil {
		fmt.Println("error generating salt :",err)
		return
	}

	// hash the password with salt
	signUpHash := hashPassword(pass,salt)

	// store the salt and password in database, just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)  // Simulate as storing in database
	fmt.Println("Hash:", signUpHash)     // Simulate as storing in database
	hashOriginalPassword := sha256.Sum256([]byte(pass))
	fmt.Println("Hash of just the password string without salt :", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// verify
	// retrieve the saltStr and decode it 
	decodedSalt,err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt.", err)
		return
	}
	loginHash := hashPassword(pass,decodedSalt)

	// compare the stored signHash with loginhash
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credentials. ")
	}


}

func generateSalt()([]byte, error){

	salt := make([]byte,16)
	_,err := io.ReadFull(rand.Reader, salt)

	if err != nil {
		return nil,err
	}
	return salt,nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}