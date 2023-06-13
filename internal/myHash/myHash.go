package myHash

import "log"

func PasswordHash(password string) uint64 {
	if len(password) < 8 {
		log.Fatal("Password can't be shorter than 8 symbols")
	}
	var hashPas uint64
	for _, i := range password {
		hashPas ^= uint64(i)
		hashPas += uint64(i) * 256
		hashPas ^= uint64(1009)
		hashPas %= 18446744073709551609
	}
	return hashPas
}
