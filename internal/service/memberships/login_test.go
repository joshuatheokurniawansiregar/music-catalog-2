package memberships

import (
	"testing"

	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/configs"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestService_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	mockRepositoryInterface := NewMockrepositoryInterface(ctrlMock)

	type args struct {
		request *memberships.LoginRequest
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
				request: &memberships.LoginRequest{
					Email: "test@gmail.com",
					Password: "testpassword",
				},
			},
			wantErr: false,
			mockFn: func(args args){
				mockRepositoryInterface.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: &gorm.Model{
						ID: 1,
					},
					Email: "test@gmail.com",
					Username: "testusername",
					Password: "$2a$10$Z6zv7jEa8.ieiRGX1K//oeqITSIBeyl.MwQl66eT5wRyfKhd41NcO",
				}, nil)
			},
		},

		{
			name: "failed when get user",
			args: args{
				request: &memberships.LoginRequest{
					Email: "test@gmail.com",
					Password: "testpassword",
				},
			},
			wantErr: true,
			mockFn: func(args args){
				mockRepositoryInterface.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when get user with wrong password",
			args: args{
				request: &memberships.LoginRequest{
					Email: "test@gmail.com",
					Password: "testpassword",
				},
			},
			wantErr: true,
			mockFn: func(args args){
				mockRepositoryInterface.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: &gorm.Model{
						ID: 1,
					},
					Email: "test@gmail.com",
					Username: "username",
					Password: "testpassword",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			var service *Service = &Service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJWT: "abc",
					},
				},
				repositoryInterface: mockRepositoryInterface,
			}
			
			got, err := service.Login(tt.args.request)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if tt.wantErr == false{
				assert.NotEmpty(t, got)
			}else{
				assert.Empty(t, got)
			}
		})
	}
}
