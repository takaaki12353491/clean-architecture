package interactor

import "cln-arch/usecase/repository"

type UserInteractor struct {
	userRepository repository.UserRepository
}
