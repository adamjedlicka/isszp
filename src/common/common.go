// Package common is a collection of helper functions, data types & constatns
package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"math/rand"
	"os"
	"unicode"
)

// DateFormat is a format string used tim time.Time.Format()
const (
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	DateTimeFormat = "2006-01-02 15:04:05"

	RandomRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// CamelToSnake transforms input string in CamelCase to sring in snake_case
func CamelToSnake(s string) string {
	buf := bytes.Buffer{}

	isPrevUpper := false
	lastRune := ' '

	for _, v := range s {
		if unicode.IsUpper(v) {
			if !isPrevUpper && lastRune != ' ' {
				buf.WriteRune('_')
			}
			isPrevUpper = true
		} else {
			isPrevUpper = false
		}

		buf.WriteRune(unicode.ToLower(v))
		lastRune = v
	}

	return buf.String()
}

// Copy file copies src file to the dst file. If any error occurs returns an error
func CopyFile(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

// RandomString generates random string of length n from RandomRunes rune pool
func RandomString(n int) string {
	str := make([]byte, n)

	for i := 0; i < n; i++ {
		str[i] = RandomRunes[rand.Intn(len(RandomRunes))]
	}

	return string(str)
}

// Encrypt encrypts text using AES encryption. key is encryption key used to encrypt text.
// returns encryptet text and errror if any occured
func Encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(crand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return ciphertext, nil
}

// Decrypt decpryts text using AES ancryption algorithm and key.
// returns decrypted data an error if any occured
func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}

	return data, nil
}
