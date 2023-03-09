package entity

type EncodeUrlRequest struct {
	UrlStr string
}

type EncodeUrlResponse struct {
	EncodedUrl string
}

type DecodeUrlRequest struct {
	UrlStr string
}

type DecodeUrlResponse struct {
	DecodedUrl string
}
