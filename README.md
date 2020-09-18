# macaddress-go

[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)

Golang client for [macaddress.io](https://macaddress.io/) API, get JSON data only.

## Usage

```go
package main

import (
  "fmt"
  "github.com/viiftw/macaddress-go"
)

func main() {
  apiKey := <YOUR-API-KEY>
  client := macaddress.NewClient(apiKey)

  // get full result data
  result, _ := client.Search("44:38:39:ff:ef:57")
  fmt.Println(result.VendorDetails.Oui)
  fmt.Println(result.VendorDetails.CompanyName)

  // get vendor company name only
  fmt.Println(client.GetVendor("44:38:39:ff:ef:57"))
}

// => Result:
// 443839
// Cumulus Networks, Inc
// Cumulus Networks, Inc
```
