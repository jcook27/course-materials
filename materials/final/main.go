package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"flag"
	"strconv"
	"time"
	"os"
	"io"
	"net/smtp"
)

func Log(l *log.Logger, msg string){
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05"))
	l.Print(msg)
}
func main() {
	var ipAddress string
	var portNumber string
	var fuzzStart int
	var fuzzEnd int
	var emailAddress string
	var emailPassword string
	var sendTo string
	var smtpPort string
	var smtpHost string
	logFile, err := os.OpenFile("fuzzer.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	flag.StringVar(&ipAddress, "ipAddress", "127.0.0.1", "The ip address of the server to fuzz. This defaults to localhost so make sure you set this.")
	flag.StringVar(&portNumber, "portNumber", "21", "The port of the server to fuzz.")
	flag.IntVar(&fuzzStart, "fuzzStart", 30, "The number of characters to start testing with.")
	flag.IntVar(&fuzzEnd, "fuzzEnd", 2500, "The number of characters to end testing with.")
	flag.StringVar(&emailAddress, "emailAddress", "", "The email address to send an email from.")
	flag.StringVar(&emailPassword, "emailPassword", "", "The password to the email address to send an email from.")
	flag.StringVar(&sendTo, "sendTo", "", "The email address to send an email to.")
	flag.StringVar(&smtpPort, "smtpPort", "587", "The smtp port to send the email to. Gmail's is 587.")
	flag.StringVar(&smtpHost, "smtpHost", "smtp.gmail.com", "The smtp host to send email.")
	flag.Parse()
	if (fuzzStart > fuzzEnd){
		log.Fatalf("ERROR: fuzzStart must be less than or equal to fuzzEnd.")
	}
	fmt.Println("\u1f31a")
	fullAddress := ipAddress + ":" + portNumber
	log.Println("INFO: Fuzzing: " + fullAddress + " starting at: " + strconv.Itoa(fuzzStart) + " characters and ending with: " + strconv.Itoa(fuzzEnd))
	for i := fuzzStart; i < fuzzEnd; i++ {
		conn, err := net.Dial("tcp", fullAddress)
		if err != nil {
			if (emailAddress != "" && emailPassword != "" && sendTo != "" ){
				messageContents := "The service on " + fullAddress + " is no longer running after " + strconv.Itoa(i) + " attempts."
				message := []byte(messageContents)
				auth := smtp.PlainAuth("", emailAddress, emailPassword, smtpHost)
				to := []string{sendTo}
				err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailAddress, to, message)
				if err != nil{
					log.Println("ERROR: Error sending email.")
				}
				log.Println("INFO: Email sent.")
			} 
			log.Fatalf("ERROR: Error at offset %d: %s\n", i, err)
		}
		bufio.NewReader(conn).ReadString('\n')
		user := ""
		for n := 0; n < i; n++ {
			user += "A"
		}
		raw := "USER %s\n"
		fmt.Fprintf(conn, raw, user)
		bufio.NewReader(conn).ReadString('\n')

		raw = "PASS password\n"
		fmt.Fprint(conn, raw)
		bufio.NewReader(conn).ReadString('\n')
		log.Println("ATTEMPT: Begining attempt: " + strconv.Itoa(i))
		if err := conn.Close(); err != nil {
			log.Println("INFO: Unable to close connection on attempt: " + strconv.Itoa(i))
		}
	}
	log.Println("INFO: Fuzzing of: " + fullAddress + " complete.")
	if (emailAddress != "" && emailPassword != "" && sendTo != "" ){
		messageContents := "The program has finished running."
		message := []byte(messageContents)
		auth := smtp.PlainAuth("", emailAddress, emailPassword, smtpHost)
		to := []string{sendTo}
		err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, emailAddress, to, message)
		if err != nil{
			log.Println("ERROR: Error sending email.")
		}
		log.Println("INFO: Email sent.")
	} 
	logFile.Close()
}
