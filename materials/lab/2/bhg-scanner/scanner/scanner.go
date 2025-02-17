// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage: To run this code either:  
//		1. From: bhg-scanner/main -> `go build` then `./main`
 //		2. From: bhg-scanner/scanner -> `go test`
 //		The code has been modified from the origianl to: use DialTimeout, keep track of the closed ports, print values as csv, return the number of open ports and closed ports, allow selecting a diffrent target to scan, and select a port to start scaning from and where to stop scaning.


package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)


func worker(target string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", target, p)    
		conn, err := net.DialTimeout("tcp", address, 60 * time.Second) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!) // Dial changed to dial timeout
		if err != nil { 
			results <- -1 * p
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object 
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner(target string, start int, end int) (int, int){  

	//TODO 3 : ADD closed ports; currently code only tracks open ports // closedports added
	var closedports []int
	var openports []int  // notice the capitalization here. access limited!
	ports := make(chan int, 150)   // 75 = 1m10.055s, 100 = 0m45.117s, 150 = 0m35.057s
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(target, ports, results)
	}

	go func() {
		for i := start; i <= end; i++ {
			ports <- i
		}
	}()

	for i := 0; i <= end - start; i++ {
		port := <-results
		if port >= 0 {
			openports = append(openports, port)
		}else{
			closedports = append(closedports, -1 * port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)
	//TODO 5 : Enhance the output for easier consumption, include closed ports

	for _, port := range openports {
		fmt.Printf("%d, open\n", port)
	}
	for _, port := range closedports {
		fmt.Printf("%d, closed\n", port)
	}
	return len(openports), len(closedports) // TODO 6 : Return total number of ports scanned (number open, number closed); 
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
