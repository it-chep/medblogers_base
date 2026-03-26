package s3

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"medblogers_base/internal/modules/freelancers/domain/doctor"
)

type fields struct {
	BucketName    string
	Client        *stubS3Client
	PresignClient S3PresignClient
}

func p(t *testing.T) fields {
	return fields{
		BucketName: "test",
		Client:     &stubS3Client{},
	}
}

type stubS3Client struct {
	listObjects func(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

func (s *stubS3Client) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	return s.listObjects(ctx, params, optFns...)
}

func (s *stubS3Client) PutObject(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return nil, nil
}

func (s *stubS3Client) GetObject(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return nil, nil
}

func (s *stubS3Client) DeleteObject(context.Context, *s3.DeleteObjectInput, ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	return nil, nil
}

func TestGateway_GetDoctorsPhotos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		setupMock   func(*stubS3Client)
		expected    map[doctor.S3Key]string
		expectedErr string
	}{
		{
			name: "успешное получение всех фотографий докторов",
			setupMock: func(m *stubS3Client) {
				m.listObjects = func(_ context.Context, params *s3.ListObjectsV2Input, _ ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
					assert.Equal(t, &s3.ListObjectsV2Input{
						Bucket: aws.String("test"),
						Prefix: aws.String("images/user_"),
					}, params)

					return &s3.ListObjectsV2Output{
						Contents: []types.Object{
							{Key: aws.String("images/user_maxim_photo1.png")},
							{Key: aws.String("images/user_marina_photo2.jpg")},
							{Key: aws.String("images/user_alina_photo3.jpeg")},
						},
					}, nil
				}
			},
			expected: map[doctor.S3Key]string{
				doctor.S3Key("images/user_maxim_photo1.png"):  "https://storage.yandexcloud.net/test/images/user_maxim_photo1.png",
				doctor.S3Key("images/user_marina_photo2.jpg"): "https://storage.yandexcloud.net/test/images/user_marina_photo2.jpg",
				doctor.S3Key("images/user_alina_photo3.jpeg"): "https://storage.yandexcloud.net/test/images/user_alina_photo3.jpeg",
			},
		},
		{
			name: "пустая директория",
			setupMock: func(m *stubS3Client) {
				m.listObjects = func(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
					return &s3.ListObjectsV2Output{Contents: []types.Object{}}, nil
				}
			},
			expected: map[doctor.S3Key]string{},
		},
		{
			name: "ошибка при запросе к S3",
			setupMock: func(m *stubS3Client) {
				m.listObjects = func(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
					return nil, errors.New("s3 internal error")
				}
			},
			expectedErr: "failed to list objects",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			deps := p(t)
			tt.setupMock(deps.Client)

			gw := NewGateway(deps.BucketName, deps.BucketName, deps.BucketName, deps.BucketName, deps.Client, deps.PresignClient)
			photosMap, err := gw.GetDoctorsPhotos(context.Background())

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
