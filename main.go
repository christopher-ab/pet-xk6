package pet_xk6

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/christopher-ab/pet-xk6/constant"
	"github.com/christopher-ab/pet-xk6/model"
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
			return
		}
	}

	filePath := fmt.Sprintf("/tmp/%s", os.Getenv(constant.EnvKeyS3UserFileName))
	file, err := os.Create(filePath)

	if err != nil {
		err = errors.New(fmt.Sprintf("error create file: %s", err.Error()))
		return
	}

	defer file.Close()
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv(constant.EnvKeyAwsRegion)),
	})

	if err != nil {
		err = errors.New(fmt.Sprintf("error init session: %s", err.Error()))
		return
	}

	downloader := s3manager.NewDownloader(awsSession)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv(constant.EnvKeyS3Bucket)),
			Key:    aws.String(os.Getenv(constant.EnvKeyS3UserFileName)),
		})

	if err != nil {
		err = errors.New(fmt.Sprintf("error downloading file from s3: %s", err.Error()))
		return
	}

	userFile, err := os.Open(filePath)
	if err != nil {
		err = errors.New(fmt.Sprintf("error opening user csv: %s", err.Error()))
		return
	}
	defer userFile.Close()
	csvReader := csv.NewReader(userFile)
	data, err := csvReader.ReadAll()
	if err != nil {
		err = errors.New(fmt.Sprintf("error reading user csv: %s", err.Error()))
		return
	}

	users := make([]model.UserLogin, len(data))
	for idx, user := range data {
		users[idx] = model.UserLogin{Email: user[0]}
	}
	jsonUsers, err := json.MarshalIndent(users, "", "  ")

	res = string(jsonUsers)
	return
}
