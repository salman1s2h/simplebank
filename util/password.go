package util

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the given plain password using bcrypt and returns the hashed string.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent.
// Returns nil if they match, or an error (bcrypt.ErrMismatchedHashAndPassword on mismatch).

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
} // VerifyPassword verifies if the given password matches the stored hash.
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
