package port

import "cln-arch/usecase/input/data"

type UserInputPort interface {
	Create(*data.UserInputData)
}
