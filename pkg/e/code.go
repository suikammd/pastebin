package e

const (
	SUCCESS = iota

	NOTFOUND
	EXPIRED

	ReadRequestError
	ParseRequestError
	CreateError
	NoShortLinkError
	ERROR
)
