package model

type OpenAPIResponse struct {
	DomainRecords DomainRecords `json:"DomainRecords"`
	PageNumber    int           `json:"PageNumber"`
	PageSize      int           `json:"PageSize"`
	RequestID     string        `json:"RequestId"`
	TotalCount    int           `json:"TotalCount"`
}
type DomainRecords struct {
	Record []Record `json:"Record"`
}
type Record struct {
	DomainName string `json:"DomainName"`
	Line       string `json:"Line"`
	Locked     bool   `json:"Locked"`
	RR         string `json:"RR"`
	RecordID   string `json:"RecordId"`
	Status     string `json:"Status"`
	TTL        int    `json:"TTL"`
	Type       string `json:"Type"`
	Value      string `json:"Value"`
	Weight     int    `json:"Weight"`
}

type RecordResponse struct {
	RequestID string `json:"RequestId"`
	RecordID  string `json:"RecordId"`
}
