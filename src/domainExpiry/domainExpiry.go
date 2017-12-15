package main

import (
	// "bytes"

	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// path, err := exec.LookPath("whois")
	cmd, err := exec.Command("whois", "plivo.com").Output()
	// fmt.Println(cmd)
	if err != nil {
		fmt.Println(err)
	}
	str := fmt.Sprintf("%s", cmd)
	strArray := strings.Split(str, "\n")
	// fmt.Println(strArray[6])
	// fmt.Printf("%q\n", strings.SplitN(str, "Registry Expiry Date:", 2))
	b := contains(strArray, "Registry")
	fmt.Println(b)
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
