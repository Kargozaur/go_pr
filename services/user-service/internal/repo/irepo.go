package repo

type IRepo[T RepoType] interface {
	Create(T) (*any, error)
	Read(T) (*any, error)
	Update(T) (*any, error)
	Delete(T) (bool, error)
}
