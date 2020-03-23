package interactor

import (
	"cln-arch/domain/model"
	inputdata "cln-arch/usecase/input/data"
	mockrepository "cln-arch/usecase/mock_repository"
	outputdata "cln-arch/usecase/output/data"
	mockoutputport "cln-arch/usecase/output/mock_port"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"
)

func TestAuth(t *testing.T) {

}

func TestCallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// repository behavior
	existingState := &model.OAuthState{Model: gorm.Model{ID: 1}, State: "state"}
	existingUser := &model.User{Model: gorm.Model{ID: 1}, Name: "user"}
	mockOAuthOutputPort := mockoutputport.NewMockOAuthOutputPort(ctrl)
	mockOAuthOutputPort.EXPECT().Callback(existingUser).Return(
		&outputdata.Callback{ID: existingUser.ID, Name: existingUser.Name, AvatorURL: existingUser.AvatorURL},
	)
	mockUserRepository := mockrepository.NewMockUserRepository(ctrl)
	mockOAuthStateRepository := mockrepository.NewMockOAuthStateRepository(ctrl)
	mockOAuthStateRepository.EXPECT().FindByState(existingState.State).Return(existingState, nil)
	mockOAuthStateRepository.EXPECT().Delete(existingState).Return(nil)
	mockOAuthTokenRepository := mockrepository.NewMockOAuthTokenRepository(ctrl)
	mockUserRepository.EXPECT().FindByID(existingUser.ID).Return(existingUser, nil)

	it := NewOAuthInteractor(
		mockOAuthOutputPort,
		mockUserRepository,
		mockOAuthStateRepository,
		mockOAuthTokenRepository,
	)

	iRequest := &inputdata.CallbackRequest{Code: "code", State: "state"}
	iUser := &inputdata.GithubUser{ID: 1, Name: "test"}
	token := &oauth2.Token{}
	iCallback := &inputdata.Callback{
		Request:    iRequest,
		User:       iUser,
		OAuthToken: token,
	}

	_, err := it.Callback(iCallback)
	if err != nil {
		t.Error("Test failed")
	}
}
