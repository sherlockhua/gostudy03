package main


import (
	"time"
	"flag"
	"fmt"
	"math/rand"
)

var (
	numCharset = "0123456789"
	strCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	mixCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	advanceCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()"
)

func main() {
	var length int
	var charset string
	flag.IntVar(&length, "l", 16, "-l the length of password")
	flag.StringVar(&charset, "t", "mix", "-t the charset of password")
	flag.Parse()
	
	rand.Seed(time.Now().UnixNano())

	var userCharset string
	switch charset {
	case "num":
		userCharset = numCharset
	case "char":
		userCharset = strCharset
	case "mix":
		userCharset = mixCharset
	case "advance":
		userCharset = advanceCharset
	default:
		userCharset = mixCharset
	}

	var password []byte
	for i := 0; i < length; i++{
		index := rand.Intn(len(userCharset))
		char := userCharset[index]
		password = append(password, char)
	}

	strPassword := string(password)
	fmt.Printf("%s\n", strPassword)

}