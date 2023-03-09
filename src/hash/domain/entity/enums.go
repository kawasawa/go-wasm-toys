package entity

type HashMethod string

const (
	MD5    HashMethod = "md5"
	SHA1   HashMethod = "sha1"
	SHA256 HashMethod = "sha256"
	SHA384 HashMethod = "sha384"
	SHA512 HashMethod = "sha512"
)
