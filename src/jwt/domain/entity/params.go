package entity

type DecodeJwtRequest struct {
	TokenString string
	Key         string
}

type DecodeJwtResponse struct {
	Header   string
	Payload  string
	Verified bool
}
