package model

type OpenAPIResponse struct {
	TotalCount    int            `json:"TotalCount"`
	RequestID     string         `json:"RequestId"`
	PageSize      int            `json:"PageSize"`
	DomainRecords []DomainRecord `json:"DomainRecords"`
	PageNumber    int            `json:"PageNumber"`
}

type DomainRecord struct {
	Rr         string `json:"RR"`
	Line       string `json:"Line"`
	Status     string `json:"Status"`
	Locked     bool   `json:"Locked"`
	Type       string `json:"Type"`
	DomainName string `json:"DomainName"`
	Value      string `json:"Value"`
	RecordID   string `json:"RecordId"`
	TTL        int    `json:"TTL"`
	Weight     int    `json:"Weight"`
}

type RecordResponse struct {
	RequestID string `json:"RequestId"`
	RecordID  string `json:"RecordId"`
}
