package inputport

import inputdata "cln-arch/usecase/input/data"

type UserInputPort interface {
	Create(*inputdata.UserInputData)
}
