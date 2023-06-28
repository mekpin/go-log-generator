package main

import (
	"fmt"
	"go-log-generator/services/cron"
	"go-log-generator/services/generator"
)

func main() {
	// generate 1 time
	generator.GenerateLog(100)

	stop := make(chan bool)

	// Run your code or tasks here

	// Block the program from exiting
	<-stop

	// This line will never be reached
	fmt.Println("This line will never be printed")
}

func init() {
	cron.Routine()
}
