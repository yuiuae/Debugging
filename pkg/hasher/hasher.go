// Copyright 2023 Serhii Khrystenko. All rights reserved.

/*
Package hasher implements add new user and password verification.

This package uses package bcrypt, witch implements Provos
and Mazières's bcrypt adaptive hashing algorithm
*/

package hasher

import (
	password "github.com/vzglad-smerti/password_hash"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword_passhash(pass string) (string, error) {
	hash, err := password.Hash(pass)
	return string(hash), err
}

func CheckPasswordHash_passhash(pass, hash string) bool {
	hash_veriry, err := password.Verify(hash, pass)
	_ = hash_veriry
	// err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(hash)) //move err
	return err == nil
}

// HashPassword_bcrypt generates a hash for the password...
func HashPassword_bcrypt(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	return string(hash), err
}

// CheckPasswordHash_bcrypt checks password by hash...
func CheckPasswordHash_bcrypt(pass, hash string) bool {
	// fmt.Println(password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(hash)) //move err
	return err == nil
}
