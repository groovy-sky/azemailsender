package main

import (
	"fmt"
	"log"
	"os"

	"github.com/groovy-sky/azemailsender"
)

func main() {
	// Simple usage example with attachment and without debug logging
	client := azemailsender.NewClient(
		"https://<RESOURCE-NAME>.communication.azure.com",
		"<YOUR_ACCESS_KEY>",
		nil, // Use default options
	)

	filePath := "/tmp/test.txt"
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	if _, err = file.WriteString("Hello World"); err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	defer func() {
		_ = file.Close()
	}()

	message, err := client.NewMessage().
		From("sender@yourdomain.com").
		To("recipient@example.com").
		Subject("Simple Test Email").
		PlainText("Hello, this is a simple test email!").
		AddAttachment(filePath).
		Build()

	if err != nil {
		log.Fatalf("Failed to build message: %v", err)
	}

	resp, err := client.Send(message)
	if err != nil {
		log.Fatalf("Email failed: %v", err)
	}

	fmt.Printf("Email sent successfully! ID: %s\n", resp.ID)
}
