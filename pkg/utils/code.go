package utils

import "crypto/rand"

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomStringByPool(n int, pool string) (string, error) {
	l := byte(len(pool))

	b, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}

	for i := 0; i < n; i++ {
		b[i] = pool[(b[i])%l]
	}

	return string(b), nil
}
