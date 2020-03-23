package interactor

import (
	"cln-arch/domain/model"
	"cln-arch/errs"
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
	existingUser := &model.User{Model: gorm.Model{ID: 1}, Name: "existing"}
	newUser := &model.User{Model: gorm.Model{ID: 2}, Name: "new"}
	mockOAuthOutputPort := mockoutputport.NewMockOAuthOutputPort(ctrl)
	mockOAuthOutputPort.EXPECT().Callback(gomock.AssignableToTypeOf(existingUser)).DoAndReturn(
		func(user *model.User) *outputdata.Callback {
			return &outputdata.Callback{ID: user.ID, Name: user.Name, AvatorURL: user.AvatorURL}
		},
	).AnyTimes()
	mockUserRepository := mockrepository.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindByID(gomock.AssignableToTypeOf(existingUser.ID)).DoAndReturn(
		func(id uint) (*model.User, error) {
			if id == existingUser.ID {
				return existingUser, nil
			}
			return nil, errs.NotFound.New("user not found")
		},
	).AnyTimes()
	mockUserRepository.EXPECT().Store(gomock.AssignableToTypeOf(existingUser)).Return(nil).AnyTimes()
	mockOAuthStateRepository := mockrepository.NewMockOAuthStateRepository(ctrl)
	mockOAuthStateRepository.EXPECT().FindByState(gomock.Any()).DoAndReturn(
		func(state string) (*model.OAuthState, error) {
			if state == existingState.State {
				return existingState, nil
			}
			return nil, errs.Forbidden.New("state not found")
		},
	).AnyTimes()
	mockOAuthStateRepository.EXPECT().Delete(existingState).Return(nil).AnyTimes()
	mockOAuthTokenRepository := mockrepository.NewMockOAuthTokenRepository(ctrl)
	mockOAuthTokenRepository.EXPECT().Store(gomock.AssignableToTypeOf(&model.OAuthToken{})).Return(nil).AnyTimes()

	it := NewOAuthInteractor(
		mockOAuthOutputPort,
		mockUserRepository,
		mockOAuthStateRepository,
		mockOAuthTokenRepository,
	)

	// state doesn't match
	iRequest := &inputdata.CallbackRequest{Code: "code", State: "invalid state"}
	iUser := &inputdata.GithubUser{ID: existingUser.ID, Name: existingUser.Name}
	token := &oauth2.Token{}
	iCallback := &inputdata.Callback{
		Request:    iRequest,
		User:       iUser,
		OAuthToken: token,
	}
	_, err := it.Callback(iCallback)
	if err == nil {
		t.Error("State doesn't match but processed successfully")
	}
	// state matches and user already exists
	iRequest = &inputdata.CallbackRequest{Code: "code", State: existingState.State}
	iUser = &inputdata.GithubUser{ID: existingUser.ID, Name: existingUser.Name}
	token = &oauth2.Token{}
	iCallback = &inputdata.Callback{
		Request:    iRequest,
		User:       iUser,
		OAuthToken: token,
	}
	oCallback, err := it.Callback(iCallback)
	if err != nil || oCallback.ID != existingUser.ID {
		t.Error("User isn't recognized correctly")
	}
	// state matches and user doesn't exist
	iRequest = &inputdata.CallbackRequest{Code: "code", State: existingState.State}
	iUser = &inputdata.GithubUser{ID: newUser.ID, Name: newUser.Name}
	token = &oauth2.Token{}
	iCallback = &inputdata.Callback{
		Request:    iRequest,
		User:       iUser,
		OAuthToken: token,
	}
	oCallback, err = it.Callback(iCallback)
	if err != nil || oCallback.ID != newUser.ID {
		t.Error("User isn't recognized correctly")
	}
}
