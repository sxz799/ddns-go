package tencentModel

type RecordList struct {
	DefaultNS     interface{} `json:"DefaultNS"`
	Line          string      `json:"Line"`
	LineID        string      `json:"LineId"`
	MX            int         `json:"MX"`
	MonitorStatus string      `json:"MonitorStatus"`
	Name          string      `json:"Name"`
	RecordID      int         `json:"RecordId"`
	Remark        string      `json:"Remark"`
	Status        string      `json:"Status"`
	TTL           int         `json:"TTL"`
	Type          string      `json:"Type"`
	UpdatedOn     string      `json:"UpdatedOn"`
	Value         string      `json:"Value"`
	Weight        interface{} `json:"Weight"`
}
