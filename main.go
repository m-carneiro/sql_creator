package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[index.Int64()]
	}
	return string(result), nil
}

func main() {
	var role string

	outputFile, err := os.Create("data.sql")
	if err != nil {
		fmt.Println("Error creating data.sql:", err)
		return
	}
	defer outputFile.Close()

	for i := 1; i <= 50; i++ {
		username := fmt.Sprintf("user%d", i)
		accessToken, err := randomString(10)
		if err != nil {
			fmt.Println("Error generating ", err)
			return
		}

		if i%2 == 0 {
			role = "ADMIN"
			sql := fmt.Sprintf("INSERT INTO matabase (id, username, access_token, role) VALUES (%d, '%s', '%s', '%s');\n", i, username, accessToken, role)

			if _, err := outputFile.WriteString(sql); err != nil {
				fmt.Println("Error writing to data.sql:", err)
				return
			}
		} else {
			role = "USER"
			sql := fmt.Sprintf("INSERT INTO matabase (id, username, access_token, role) VALUES (%d, '%s', '%s', '%s');\n", i, username, accessToken, role)

			if _, err := outputFile.WriteString(sql); err != nil {
				fmt.Println("Error writing to data.sql:", err)
				return
			}
		}
	}
}
