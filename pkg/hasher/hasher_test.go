package hasher

import (
	"fmt"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkHashPassword_bcrypt(b)
		benchmarkCheckPasswordHash_bcrypt(b)
	}
}

// func BenchmarkMain2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		benchmarkHashPassword_passhash(b)
// 		benchmarkCheckPasswordHash_passhash(b)
// 	}
// }

func benchmarkHashPassword_passhash(b *testing.B) {
	pass := "password"
	hash, err := HashPassword_passhash(pass)
	if err != nil {
		b.Errorf("Error #{err}:")
	}
	_ = hash
	// fmt.Println(hash)
}

func benchmarkCheckPasswordHash_passhash(b *testing.B) {
	pass := "password"
	hash := "ptkTBNQ3rIAKTYJvmtRn2C2iO3A9dkzv6FReP12eyXRMTQgSwECZPeh9NopLF3geAdXfeMP9pXaEZZ3W8j8UmG095aFrMEsiS7Fp1Ksf$#$17$#$9ff19a5fefc4fb3701b4ca1344980f3b3b991e544922c73bb740af55$#$94dc589262bf6aea2bd5c9e75ce0a3757525e72aea211ffd0a60f15bcf498d2a"
	ok := CheckPasswordHash_passhash(pass, hash)
	_ = ok
	// fmt.Println(ok)
}

func benchmarkHashPassword_bcrypt(b *testing.B) {
	pass := "password"
	hash, err := HashPassword_bcrypt(pass)
	if err != nil {
		b.Errorf("Error #{err}:")
	}
	_ = hash
	// fmt.Println(hash)
}

func benchmarkCheckPasswordHash_bcrypt(b *testing.B) {
	pass := "password"
	hash := "$2a$08$cuzJYqNhKhGY2bxYCHExV.kUwUwFSQUrwHZGISR7TXveseNozjpru"
	ok := CheckPasswordHash_bcrypt(pass, hash)
	_ = ok
	// fmt.Println(ok)
	// if !b {
	// 	t.Errorf("Error in CheckPasswordHash")
	// }
}

func _TestHashPassword(t *testing.T) {
	pass := "password"
	hash, err := HashPassword_bcrypt(pass)
	if err != nil {
		t.Errorf("Error #{err}:")
	}
	fmt.Println(hash)
}

func _TestCheckPasswordHash(t *testing.T) {
	pass := "password"
	hash := "$2a$08$cuzJYqNhKhGY2bxYCHExV.kUwUwFSQUrwHZGISR7TXveseNozjpru"
	b := CheckPasswordHash_bcrypt(pass, hash)
	fmt.Println(b)
	// if !b {
	// 	t.Errorf("Error in CheckPasswordHash")
	// }
}
