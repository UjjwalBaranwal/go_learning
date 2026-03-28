package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string
	UnreadCount   int
}

func getEmailTemplate() string {
	emailTemplate := `
Subject: {{.Subject}}
Hi {{.RecipientName}} {{.Body}}
{{if .Items}}
{{range .Items}}
- {{.}}
{{end}}
{{end}}

{{if gt .UnreadCount 0}}
You have {{.UnreadCount}} left
{{else}}
You have no message
{{end}}

- Thanks
{{.SenderName}}
	`

	return emailTemplate
}
func main() {
	fmt.Println("--- Text Template Example")
	emailTemplate := getEmailTemplate()
	temp1, err := template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	data := EmailData{
		SenderName:    "Ujjwal",
		RecipientName: "Baranwal",
		Subject:       "Regarding Learning Go",
		Body:          "You are doing your best , be consistent",
		Items:         []string{"Item1", "Item2", "Item3"},
		UnreadCount:   2,
	}
	// err = temp1.Execute(os.Stdout, data)// print on console output
	var output strings.Builder
	err = temp1.Execute(&output, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// fmt.Println(output.String())
	fmt.Println(strings.ToUpper(output.String()))
}
