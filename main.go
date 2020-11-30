package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	usernameFile string
	passwordFile string
	outFile      string
)

func mustBeSpecified(v, msg string) {
	if len(v) == 0 {
		log.Fatal(msg)
	}
}

func mustRead(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(b), "\n")
}

func mustWrite(file string, s string) {
	if err := ioutil.WriteFile(file, []byte(s), 0666); err != nil {
		log.Fatal(err)
	}
}

func mustMakeMD5userlist(username, password string) string {
	hasher := md5.New()
	_, err := hasher.Write([]byte(password + username))
	if err != nil {
		log.Fatal(err)
	}
	md5paswd := "md5" + hex.EncodeToString(hasher.Sum(nil))
	return fmt.Sprintf("%q %q", username, md5paswd)
}

func main() {
	flag.StringVar(&usernameFile, "u", usernameFile, "Input file with username. Required")
	flag.StringVar(&passwordFile, "p", passwordFile, "Input file with password. Required")
	flag.StringVar(&outFile, "o", outFile, "Output auth file. Required")
	flag.Parse()

	mustBeSpecified(usernameFile, "Expected username file specified")
	mustBeSpecified(passwordFile, "Expected password file specified")
	mustBeSpecified(outFile, "Expected output file specified")

	mustWrite(
		outFile,
		mustMakeMD5userlist(
			mustRead(usernameFile),
			mustRead(passwordFile),
		),
	)
}
