package domain

type IHLSValidator interface {
	Validate(uri string) (string, bool, error)
}
