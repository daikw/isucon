// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println(digest("aaa"))
	fmt.Println(digest2("aaa"))
	fmt.Println(digest3("aaa"))

	fmt.Println(escapeshellarg("aaa"))
}

func escapeshellarg(arg string) string {
	return "'" + strings.Replace(arg, "'", "'\\''", -1) + "'"
}

func escapeshellarg2(arg string) string {
	return strings.Replace(arg, "'", "'\\''", -1)
}

func digest(src string) string {
	// opensslのバージョンによっては (stdin)= というのがつくので取る
	out, err := exec.Command("/bin/bash", "-c", `printf "%s" `+escapeshellarg(src)+` | openssl dgst -sha512 | sed 's/^.*= //'`).Output()
	if err != nil {
		log.Print(err)
		return ""
	}

	return strings.TrimSuffix(string(out), "\n")
}

func digest2(src string) string {
	checksum := sha512.Sum512([]byte(escapeshellarg2(src)))
	return hex.EncodeToString(checksum[:])
}

func digest3(src string) string {
	s := sha512.New()
	s.Write([]byte(escapeshellarg2(src)))
	return hex.EncodeToString(s.Sum(nil))
}

