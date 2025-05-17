package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, Base64 Encoding")

	// Encode base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode form base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println("decoding error :", err)
		return
	}
	fmt.Println(decoded)

	// URL safe, avoid '/' and '+'
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe encoded: ", urlSafeEncoded)

}
