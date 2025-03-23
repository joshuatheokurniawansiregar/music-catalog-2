package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_Signup(t *testing.T) {
	
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	membershipsServiceInterface := NewMockserviceInterface(ctrlMock)

	tests := []struct {
		name string
		mockFn func()
		expectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			name: "success",
			mockFn: func(){
				membershipsServiceInterface.EXPECT().Signup(memberships.SignUpRequest{
					Email: "testaa@gmail.com",
					Username: "testusername",
					Password: "testpassword",
				}).Return(nil)
			},
			expectedStatusCode: 201,
		},

		{
			name: "failed",
			mockFn: func(){
				membershipsServiceInterface.EXPECT().Signup(memberships.SignUpRequest{
					Email: "testaa@gmail.com",
					Username: "testusername",
					Password: "testpassword",
				}).Return(errors.New("email or username exists"))
			},
			expectedStatusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api:= gin.New()
			h:= Handler{
				Engine: api,
				serviceInterface: membershipsServiceInterface,
			}
			h.RegisterRoute()

			responseWriter := httptest.NewRecorder()
			var endpoint string= `/api/v1/memberships/sign_up`
			model:= memberships.SignUpRequest{
				Email: "testaa@gmail.com",
				Username: "testusername",
				Password: "testpassword",
			}
			val,err:= json.Marshal(model)
			assert.NoError(t,err)

			body:= bytes.NewReader(val)

			req, err:= http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(responseWriter, req)
			assert.Equal(t, tt.expectedStatusCode, responseWriter.Code)
		})
	}
}
