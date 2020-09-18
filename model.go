package macaddress

// Response struct of result
type Response struct {
	BlockDetails struct {
		AssignmentBlockSize string `json:"assignmentBlockSize"`
		BlockFound          bool   `json:"blockFound"`
		BlockSize           int64  `json:"blockSize"`
		BorderLeft          string `json:"borderLeft"`
		BorderRight         string `json:"borderRight"`
		DateCreated         string `json:"dateCreated"`
		DateUpdated         string `json:"dateUpdated"`
	} `json:"blockDetails"`
	MacAddressDetails struct {
		AdministrationType string   `json:"administrationType"`
		Applications       []string `json:"applications"`
		Comment            string   `json:"comment"`
		IsValid            bool     `json:"isValid"`
		SearchTerm         string   `json:"searchTerm"`
		TransmissionType   string   `json:"transmissionType"`
		VirtualMachine     string   `json:"virtualMachine"`
		WiresharkNotes     string   `json:"wiresharkNotes"`
	} `json:"macAddressDetails"`
	VendorDetails struct {
		CompanyAddress string `json:"companyAddress"`
		CompanyName    string `json:"companyName"`
		CountryCode    string `json:"countryCode"`
		IsPrivate      bool   `json:"isPrivate"`
		Oui            string `json:"oui"`
	} `json:"vendorDetails"`
}
