package scanner

import (
	"testing"
)
var target = "scanme.nmap.org"
// Variables to set for changing the starting port and ending port for the scan. Do not set them as negative and do not set end to be less than start. Do not set either to being more than 65,535 because that is the maximum number of ports.
var start = 20
var end = 100
// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T){

    gotO, _ := PortScanner(target, 1, 1024) // Currently function returns only number of open ports
    want := 2 // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns 
	          //consider what would happen if you parameterize the portscanner address and ports to scan

    if gotO != want {
        t.Errorf("got %d, wanted %d", gotO, want)
    }
}

func TestTotalPortsScanned(t *testing.T){
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()

    gotO, gotC := PortScanner(target, 1, 1024) // Currently function returns only number of open ports
    want := 1024 // default value; consider what would happen if you parameterize the portscanner ports to scan

    if gotO + gotC != want {
        t.Errorf("got %d, wanted %d", gotO + gotC, want)
    }
}

func TestVariablePortsScanned(t *testing.T){
    gotO, gotC := PortScanner(target, start, end)
    want := end - start + 1

    if gotO + gotC != want {
        t.Errorf("got %d, wanted %d", gotO + gotC, want)
    }
}
