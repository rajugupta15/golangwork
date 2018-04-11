package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeTofile(stack string, serviceName []string) {
	for i := 0; i < len(serviceName); i++ {
		filename := serviceName[i] + ".conf"
		file, err := os.Create(filename)
		check(err)
		defer file.Close()
		fmt.Fprintf(file, "<source>\n@type tail\npath /audit-logs/%s/%s/*.log\npos_file /audit-logs/td-agent/s3/logs/%s/%s/%s-log.pos\nformat none\ntag s3.%s.%s\n</source>", stack, serviceName[i], stack, serviceName[i], serviceName[i], stack, serviceName[i])
	}
}

func main() {
	stack := []string{"stack1", "stack2"}
	for i := 0; i < len(stack); i++ {
		if stack[i] == "stack1" {
			serviceList := []string{"stack1-service1", "stack1-service2", "stack1-service3"}
			writeTofile(stack[i], serviceList)
		} else if stack[i] == "stack2" {
			serviceList := []string{"stack2-service1", "stack2-service2"}
			writeTofile(stack[i], serviceList)
		} else {
			fmt.Println("Not found")
		}
	}

}
