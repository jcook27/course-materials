To use this program:
go build main.go
SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term> <page number>

Example: SHODAN_API_KEY=YOURAPIKEYHERE ./main product:OpenSSH 5
This will return the 5th page of the results for OpenSSH on Shodan.io