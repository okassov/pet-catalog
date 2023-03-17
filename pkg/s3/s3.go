package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	_defaultRegion           = "us-east-1"
	_defaultEndpoint         = ""
	_defaultS3ForcePathStyle = false
)

type S3Config struct {
	region           string
	endpoint         string
	s3ForcePathStyle bool
	credentials      *credentials.Credentials
}

func New(url string) (*s3.S3, error) {

	s3Config := &S3Config{
		region:           _defaultRegion,
		endpoint:         _defaultEndpoint,
		s3ForcePathStyle: _defaultS3ForcePathStyle,
	}

	session, err := session.NewSession(&aws.Config{
		Region:           &s3Config.region,
		Endpoint:         &s3Config.endpoint,
		S3ForcePathStyle: &s3Config.s3ForcePathStyle,
		Credentials:      s3Config.credentials,
	})
	if err != nil {
		return nil, fmt.Errorf("s3 - NewS3 - session.NewSession: %w", err)
	}

	s3Client := s3.New(session)

	return s3Client, nil
}
