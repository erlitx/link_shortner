package v1

import "github.com/erlitx/link_shortner/internal/usecase"


type Handlers struct {
	usecase *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}
