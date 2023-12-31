package main

import (
  "bufio"
  "fmt"
  "log"
  "net"
  "os"
  "strings"
)

func main() {
  // fmt.Println("Hello, World!")
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Printf("domain, hasMK, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

  for scanner.Scan() {
    checkDomain(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal("Could not scan input: ", err)
  }
}

func checkDomain(domain string) {
  var hasMK, hasSPF, hasDMARC bool
  var sprRecord, dmarcRecord string

  // Connect to server
  mxRecords, err := net.LookupMX(domain)
  if err != nil {
    log.Printf("Error: %v\n", err)
  }
  if len(mxRecords) > 0 {
    hasMK = true
  }

  txtRecords, err := net.LookupTXT(domain)
  if err != nil {
    log.Printf("Error: %v\n", err)
  }

  for _, record := range txtRecords {
    if strings.HasPrefix(record, "v=spf1") {
      hasSPF = true
      sprRecord = record
      break
    }
  }

  dmarcRecords, err := net.LookupTXT("_dmarc" + domain)
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

  // Print results
  fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMK, hasSPF, sprRecord, hasDMARC, dmarcRecord)
}
