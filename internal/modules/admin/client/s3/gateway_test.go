package s3

import (
	"context"
	"medblogers_base/internal/modules/doctors/client/s3/mocks"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fields struct {
	BucketName    string
	Client        *mocks.MockS3Client
	PresignClient *mocks.MockS3PresignClient
}

func p(t *testing.T) fields {
	return fields{
		BucketName:    "test",
		Client:        mocks.NewMockS3Client(gomock.NewController(t)),
		PresignClient: mocks.NewMockS3PresignClient(gomock.NewController(t)),
	}
}

func TestGateway_GetUserPhotos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		setupMock   func(*mocks.MockS3Client)
		expected    map[string]string
		expectedErr string
	}{
		{
			name: "успешное получение всех фотографий пользователей",
			setupMock: func(m *mocks.MockS3Client) {
				m.EXPECT().ListObjectsV2(gomock.Any(), &s3.ListObjectsV2Input{
					Bucket: aws.String("test"),
					Prefix: aws.String("images/user_"),
				}).Return(&s3.ListObjectsV2Output{
					Contents: []types.Object{
						{Key: aws.String("images/user_maxim_photo1.png")},
						{Key: aws.String("images/user_marina_photo2.jpg")},
						{Key: aws.String("images/user_alina_photo3.jpeg")},
					},
				}, nil)
			},
			expected: map[string]string{
				"maxim":  "https://storage.yandexcloud.net/test/images/user_maxim_photo1.png",
				"marina": "https://storage.yandexcloud.net/test/images/user_marina_photo2.jpg",
				"alina":  "https://storage.yandexcloud.net/test/images/user_alina_photo3.jpeg",
			},
		},
		{
			name: "пустая директория",
			setupMock: func(m *mocks.MockS3Client) {
				m.EXPECT().ListObjectsV2(gomock.Any(), gomock.Any()).
					Return(&s3.ListObjectsV2Output{Contents: []types.Object{}}, nil)
			},
			expected: map[string]string{},
		},
		{
			name: "ошибка при запросе к S3",
			setupMock: func(m *mocks.MockS3Client) {
				m.EXPECT().ListObjectsV2(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("s3 internal error"))
			},
			expectedErr: "failed to list objects",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			deps := p(t)
			tt.setupMock(deps.Client)

			gw := NewGateway(deps.BucketName, deps.Client, deps.PresignClient)
			photosMap, err := gw.GetUserPhotos(context.Background())

			if tt.expectedErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, photosMap)
		})
	}
}
