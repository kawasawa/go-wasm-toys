package entity

type CaseRequest struct {
	PlainText string
}

type CaseResponse struct {
	Camel  string
	Pascal string
	Snake  string
	Kebab  string
	Upper  string
	Lower  string
}
