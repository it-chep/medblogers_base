package salebot

import (
	"context"
	indto "medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/pkg/http/mocks"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

type fields struct {
	http *mocks.MockExecutor
}

func p(t *testing.T) fields {
	f := fields{
		http: mocks.NewMockExecutor(gomock.NewController(t)),
	}
	return f
}

func TestNotificatorCreateDoctor(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		deps := p(t)
		gateway := NewGateway("http://example.com", deps.http)

		createDTO := indto.CreateDoctorRequest{
			Slug:              "test-doctor",
			FullName:          "Test Doctor",
			InstagramUsername: "test_insta",
			TelegramUsername:  "test_tg",
		}

		deps.http.EXPECT().Do(gomock.Any()).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       http.NoBody,
			}, nil)

		gateway.NotificatorCreateDoctor(context.Background(), createDTO, 123)
	})

	t.Run("HttpError", func(t *testing.T) {
		t.Parallel()

		deps := p(t)
		gateway := NewGateway("http://example.com", deps.http)

		deps.http.EXPECT().Do(gomock.Any()).
			Return(nil, errors.New("http error"))

		gateway.NotificatorCreateDoctor(context.Background(), indto.CreateDoctorRequest{
			Slug:     "test-doctor",
			FullName: "Test Doctor",
		}, 123)
	})
}
