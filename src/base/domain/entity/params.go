package entity

type EncodeBaseRequest struct {
	PlainText string
}

type EncodeBaseResponse struct {
	Base32 string
	Base64 string
}

type DecodeBaseRequest struct {
	BaseText string
}

type DecodeBaseResponse struct {
	Plain32 string
	Plain64 string
}
