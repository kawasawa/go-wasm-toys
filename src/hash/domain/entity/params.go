package entity

type HashRequest struct {
	PlainText string
}

type HashResponse struct {
	MD5    string
	SHA1   string
	SHA256 string
	SHA384 string
	SHA512 string
}
