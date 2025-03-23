package memberships

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockServiceInterface := NewMockserviceInterface(ctrlMock)

	tests := []struct {
		name string
		mockFn func()
		wantErr bool
		expectedStatusCode int
		expectedBody memberships.LoginResponse
	}{
		// TODO: Add test cases.
		{
			name: "success",
			mockFn: func() {
				mockServiceInterface.EXPECT().Login(&memberships.LoginRequest{
					Email: "testaa@gmail.com",
					Password: "password",
				}).Return("accessToken", nil)
			},
			wantErr: false,
			expectedStatusCode: 202,
			expectedBody: memberships.LoginResponse{
				AccessToken: "accessToken",
			},
		},

		{
			name: "failed",
			mockFn: func() {
				mockServiceInterface.EXPECT().Login(&memberships.LoginRequest{
					Email: "testaa@gmail.com",
					Password: "password",
				}).Return("", assert.AnError)
			},
			wantErr: true,
			expectedStatusCode: 202,
			expectedBody: memberships.LoginResponse{
				AccessToken: "accessToken",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api:= gin.New()
			h:= Handler{
				Engine: api,
				serviceInterface: mockServiceInterface,
			}
			h.RegisterRoute()
			responseWriter := httptest.NewRecorder()
			var endpoint string = `/api/v1/memberships/login`
			model:= &memberships.LoginRequest{
				Email: "testaa@gmail.com",
				Password: "password",
			}
			val, err := json.Marshal(model)
			assert.NoError(t, err)
			body:= bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(responseWriter, req)
			assert.Equal(t, tt.expectedStatusCode, responseWriter.Code)
			if !tt.wantErr{
				res:= responseWriter.Result()
				defer res.Body.Close()

				responseModel := memberships.LoginResponse{}
				err = json.Unmarshal(responseWriter.Body.Bytes(), &responseModel)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, responseModel)
			}
		})
	}
}
