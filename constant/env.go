package constant

const (
	EnvKeyAwsRegion          = "AWS_REGION"
	EnvKeyAwsAccessKeyId     = "AWS_ACCESS_KEY_ID"
	EnvKeyAwsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	EnvKeyAwsSessionToken    = "AWS_SESSION_TOKEN"

	EnvKeyS3Bucket       = "S3_BUCKET"
	EnvKeyS3UserFileName = "S3_USER_FILE_NAME"
)

var EnvAwsConfig = map[string]string{
	EnvKeyAwsRegion:          "AWS_REGION",
	EnvKeyAwsAccessKeyId:     "AWS_ACCESS_KEY_ID",
	EnvKeyAwsSecretAccessKey: "AWS_SECRET_ACCESS_KEY",
	EnvKeyAwsSessionToken:    "AWS_SESSION_TOKEN",
	EnvKeyS3Bucket:           "S3_BUCKET",
	EnvKeyS3UserFileName:     "S3_USER_FILE_NAME",
}
