package s3

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Option func(*S3Config)

// Credentials -.
func Credentials(accessKeyId string, secretAccessKey string) Option {
	return func(c *S3Config) {
		c.credentials = credentials.NewStaticCredentials(accessKeyId, secretAccessKey, "")
	}
}

// Region -.
func Region(region string) Option {
	return func(c *S3Config) {
		c.region = region
	}
}

// Endpoint -.
func Endpoint(endpoint string) Option {
	return func(c *S3Config) {
		c.endpoint = endpoint
	}
}
