package myHash

import "log"

func Password_hash(password string) uint64 {
	if len(password) < 8 {
		log.Fatal("Password can't be shorter than 8 symbols")
	}
	var hash_pas uint64
	for _, i := range password {
		hash_pas ^= uint64(i)
		hash_pas += uint64(i) * 256
		hash_pas ^= uint64(1009)
		hash_pas %= 18446744073709551609
	}
	return hash_pas
}
