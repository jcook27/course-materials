package main

import "bhg-scanner/scanner"

func main(){
	scanner.PortScanner("scanme.nmap.org", 1, 100)
}