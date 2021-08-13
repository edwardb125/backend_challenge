package Struct

// this struct is for input data as we expected
type Input struct{
	Id          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}

// this struct return the final message after post method
type Output struct{
	Message string `json:"message"`
}
