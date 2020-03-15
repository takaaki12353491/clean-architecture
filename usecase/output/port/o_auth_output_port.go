package outputport

import outputdata "cln-arch/usecase/output/data"

type OAuthOutputPort interface {
	Login(state string, url string) *outputdata.Login
}
