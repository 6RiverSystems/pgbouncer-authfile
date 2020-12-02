package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/6RiverSystems/pgbouncer-authfile/pkg/flagtypes"
)

const (
	MD5 = "md5"
)

var (
	usernames flagtypes.InputData
	passwords flagtypes.InputData
	pwdType   string
	out       = flagtypes.OutFile{os.Stderr}
)

func md5password(username, password string) string {
	hasher := md5.New()
	_, err := hasher.Write([]byte(password + username))
	if err != nil {
		log.Fatalf("failed to write md5 hash: %v", err)
	}
	return "md5" + hex.EncodeToString(hasher.Sum(nil))
}

func maskPassword(pwdType, username, password string) string {
	switch pwdType {
	case MD5:
		return md5password(username, password)
	default:
		return password
	}
}

func main() {
	flag.Var(&usernames, "u", "Input file with username. Required")
	flag.Var(&passwords, "p", "Input file with password. Required")
	flag.StringVar(&pwdType, "t", pwdType, "Password type. Plain text if not specified. Possible values: md5. Optional")
	flag.Var(&out, "o", "Output file. Required")
	flag.Parse()

	if len(usernames) != len(passwords) {
		log.Fatalf("Number of usernames (%d) is not equal to number of passwors (%d)", len(usernames), len(passwords))
	}

	w := bufio.NewWriter(out.File)
	for i, username := range usernames {
		password := maskPassword(pwdType, username, passwords[i])
		_, err := fmt.Fprintf(w, "%q %q\n", username, password)
		if err != nil {
			log.Fatalf("failed to write record: %v", err)
		}
	}
	if err := w.Flush(); err != nil {
		log.Fatalf("failed to flush writer: %v", err)
	}
}
