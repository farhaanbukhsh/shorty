package storage

// Service helps to define the fuctions needed to communicate with database
type Service interface {
	Save(url string, slug string) (string, error)
	Load(code string) (string, error)
	Close() error
}
