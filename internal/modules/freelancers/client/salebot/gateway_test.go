package salebot

import (
	"context"
	indto "medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
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

		createDTO := indto.CreateRequest{
			Slug:       "test-doctor",
			Name:       "Test Doctor",
			TgUsername: "test_tg",
		}

		deps.http.EXPECT().Do(gomock.Any()).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       http.NoBody,
			}, nil)

		gateway.NotificatorCreateFreelancer(context.Background(), createDTO, 123)
	})

	t.Run("HttpError", func(t *testing.T) {
		t.Parallel()

		deps := p(t)
		gateway := NewGateway("http://example.com", deps.http)

		deps.http.EXPECT().Do(gomock.Any()).
			Return(nil, errors.New("http error"))

		gateway.NotificatorCreateFreelancer(context.Background(), indto.CreateRequest{
			Slug: "test-doctor",
			Name: "Test Doctor",
		}, 123)
	})
}
