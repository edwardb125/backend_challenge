package Struct

// Output this struct is for output data as we expected
type Output struct{
	Id          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}

// Input this struct come from the input
type Input struct{
	ID string `json:"id"`
}
