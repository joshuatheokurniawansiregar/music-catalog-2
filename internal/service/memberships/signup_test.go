package memberships

import (
	"testing"

	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/configs"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestService_Signup(t *testing.T) {
	ctrlMock:= gomock.NewController(t)
	defer ctrlMock.Finish()
	
	mockRepositoryInterface := NewMockrepositoryInterface(ctrlMock)
	
	type args struct {
		request memberships.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn func(args args)
	}{
		// TODO: Add test cases.
		{
			name: "successs",
			args: args{
				request: memberships.SignUpRequest{
					Email: "test@gmail.com",
					Username: "testusername",
					Password: "testpassword",
				},
			},
			wantErr: false,
			mockFn: func(args args){
				mockRepositoryInterface.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepositoryInterface.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},

		{
			name: "failed when get user",
			args: args{
				request: memberships.SignUpRequest{
					Email: "test@gmail.com",
					Username: "testusername",
					Password: "testpassword",
				},
			},
			wantErr: true,
			mockFn: func(args args){
				mockRepositoryInterface.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, assert.AnError)
			},
		},

		{
			name: "failed when create user",
			args: args{
				request: memberships.SignUpRequest{
					Email: "test@gmail.com",
					Username: "testusername",
					Password: "testpassword",
				},
			},
			wantErr: true,
			mockFn: func(args args){
				mockRepositoryInterface.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepositoryInterface.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			var service *Service = &Service{
				cfg: &configs.Config{},
				repositoryInterface: mockRepositoryInterface,
			}
			if err := service.Signup(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Service.Signup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
