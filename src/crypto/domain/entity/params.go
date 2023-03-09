package entity

type EncryptRequest struct {
	PlainText string
	Key       string
}

type EncryptResponse struct {
	EncryptedText string
}

type DecryptRequest struct {
	EncryptedText string
	Key           string
}

type DecryptResponse struct {
	PlainText string
}
