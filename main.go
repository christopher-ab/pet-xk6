package pet_xk6

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/christopher-ab/pet-xk6/constant"
	"go.k6.io/k6/js/modules"
	"os"
	"strings"
)

func init() {
	modules.Register("k6/x/pet-util", new(PETUtil))
}

type PETUtil struct{}

func (pet *PETUtil) Test() (res string, err error) {
	for _, env := range constant.EnvAwsConfig {
		if strings.TrimSpace(os.Getenv(env)) == "" {
			err = errors.New(fmt.Sprintf("invalid env for %s", env))
		}
		return
	}
	file, err := os.Create(fmt.Sprintf("/root/%s", constant.EnvKeyS3UserFileName))
	if err != nil {
		return
	}
	defer file.Close()
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv(constant.EnvKeyAwsRegion)),
	})
	if err != nil {
		return
	}
	downloader := s3manager.NewDownloader(awsSession)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(constant.EnvKeyS3Bucket),
			Key:    aws.String(constant.EnvKeyS3UserFileName),
		})
	if err != nil {
		return
	}

	res = "OK"
	return
}
