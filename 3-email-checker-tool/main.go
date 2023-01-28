package main

import(
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Welcome!\n")
	fmt.Printf("Write the e-mail domain: ")

	for scanner.Scan(){
		text := scanner.Text()
		if len(text) > 0 {
			checkDomain(text)
		}
		fmt.Printf("\nWrite the e-mail domain: ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string){
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	
	fmt.Printf("\n")
	fmt.Printf("|____domain____|__hasMX__|__hasSPF__|____________sprRecord____________|__hasDMARC__|_______________________dmarcRecord_______________________|\n")
	fmt.Printf("| %12v | %7v | %8v | %31v | %10v | %55v |\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}