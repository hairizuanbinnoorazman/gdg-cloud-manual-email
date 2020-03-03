package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hairizuanbinnoorazman/gdg-cloud-manual-email/emailSender"
	"github.com/spf13/cobra"
)

var dryRun bool
var qwikLabCode string
var sendGridKey string
var emailListFile string
var subject string

var rootCmd = &cobra.Command{
	Use:   "email",
	Short: "Email is a bulk email tool to quickly fire off to email to users",
	Run: func(cmd *cobra.Command, args []string) {
		if dryRun == true {
			output, err := parseTemplate(qwikLabCode)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(output)
			return
		}
		output, err := parseTemplate(qwikLabCode)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		key, available := os.LookupEnv("SENDGRID_KEY")
		if available == false {
			panic("SENDGRID KEY IS NOT AVAILABLE")
		}
		sendGridKey = key

		fmt.Println(output)
		emailList, err := getEmails(emailListFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, item := range emailList {
			fmt.Println(item)
			emailer := emailSender.SendGrid{Key: sendGridKey}
			err = emailer.Send(item, "support@gdg-cloud-singapore.com", subject, output)
			if err != nil {
				fmt.Printf("%+v", err)
				return
			}
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&qwikLabCode, "qwiklab", "sample-do-not-use", "To pass qwiklab code to email template")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dryRun", true, "Used to run the tool in dry run mode. Default is true")
	rootCmd.PersistentFlags().StringVar(&emailListFile, "emailListFile", "email_list_file", "Provide the email list file that is to be used for this")
	rootCmd.PersistentFlags().StringVar(&subject, "subject", "default subject", "Provide a subject to be used for email")
}

func parseTemplate(code string) (string, error) {
	t := template.New("email_template.html")

	var err error
	t, err = t.ParseFiles("email_template.html")
	if err != nil {
		return "", err
	}

	data := struct {
		Code string
	}{
		Code: code,
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	result := tpl.String()
	return result, nil
}

func getEmails(fileName string) ([]string, error) {
	rawFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []string{}, err
	}
	data := string(rawFile)
	emailList := strings.Split(data, "\n")
	return emailList, nil
}
