package service

type Service struct {
	repo Repository
}

type Repository interface {
	GetNewWordFromStorage(length int) (string, error)
	GetDailyWordFromStorage() (string, error)
}

type Dependencies struct {
	Repo Repository
}

func NewService(deps Dependencies) *Service {
	return &Service{
		repo: deps.Repo,
	}
}
