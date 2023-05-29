package cardreader

type Interface interface {
	ReadCardNumber() (string, error)
}
