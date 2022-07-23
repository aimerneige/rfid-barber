package pkg

import "golang.org/x/crypto/bcrypt"

// HashPasswordWithBcrypt hash password with bcrypt
func HashPasswordWithBcrypt(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword := string(hashByte)

	return hashedPassword, err
}

// CheckPasswordWithBcrypt check password with bcrypt
func CheckPasswordWithBcrypt(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
