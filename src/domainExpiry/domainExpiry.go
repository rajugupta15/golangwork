package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func handler(b error) error {
	if b != nil {
		fmt.Println(b)
	}
	return b
}
func domainExpiry(domain string) string {
	cmd, err := exec.Command("whois", domain).Output()
	handler(err)
	str := fmt.Sprintf("%s", cmd)
	strArray := strings.Split(str, "\n")
	s := strings.Split(strArray[6], ":")
	expiry := s[1:]
	format := strings.Join(expiry, ":")
	trim := strings.TrimSpace(format)
	return (trim)
}
func timeFormat(t string) time.Time {
	t1, e := time.Parse(time.RFC3339, t)
	handler(e)
	s := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.UTC)
	return s
}
func main() {
	domain := "thegeeklinux.com"
	a := domainExpiry(domain)
	b := timeFormat(a)
	c := time.Now()
	duration := b.Sub(c)
	fmt.Println(domain, "will expire in", int(duration.Hours()/24), "days")
}
