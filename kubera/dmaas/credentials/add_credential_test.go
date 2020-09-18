package credentials

import (
	"os"
	"testing"
)

func TestCredentialAddition(t *testing.T) {

	var (
		awsAccessID       = os.Getenv("AWS_ACCESS_ID")
		awsSecretKey      = os.Getenv("AWS_SECRET_KEY")
		cloudianAccessID  = os.Getenv("CLOUDIAN_ACCESS_ID")
		cloudianSecretKey = os.Getenv("CLOUDIAN_SECRET_KEY")
		minioAccessID     = os.Getenv("MINIO_ACCESS_ID")
		minioSecretKey    = os.Getenv("MINIO_SECRET_KEY")
	)

	var tests = map[string]struct {
		CloudProvider string
		Name          string
		AccessID      string
		SecretKey     string
		IsError       bool
	}{
		"aws bucket with all details": {
			CloudProvider: "aws",
			Name:          "aws-dmaas-test",
			AccessID:      awsAccessID,
			SecretKey:     awsSecretKey,
		},
		"aws bucket with duplicate name": {
			CloudProvider: "aws",
			Name:          "aws-dmaas-test",
			AccessID:      awsAccessID,
			SecretKey:     awsSecretKey,
		},
		"aws bucket with name length less than 6 character": {
			CloudProvider: "aws",
			Name:          "awsdm",
			AccessID:      awsAccessID,
			SecretKey:     awsSecretKey,
		},
		"aws bucket with name length more than 25 character": {
			CloudProvider: "aws",
			Name:          "aws-dmaas-test-backup-restore-cred",
			AccessID:      awsAccessID,
			SecretKey:     awsSecretKey,
		},
		"aws bucket without name": {
			CloudProvider: "aws",
			AccessID:      awsAccessID,
			SecretKey:     awsSecretKey,
			IsError:       true,
		},
		"aws bucket without accessID ": {
			CloudProvider: "aws",
			Name:          "aws-cred-neg",
			SecretKey:     awsSecretKey,
			IsError:       true,
		},
		"aws bucket without secretKey": {
			CloudProvider: "aws",
			Name:          "aws-cred-neg",
			AccessID:      awsAccessID,
			IsError:       true,
		},
		"aws bucket without name, accessID & secretKey": {
			CloudProvider: "aws",
			IsError:       true,
		},
		"cloudian bucket with all details": {
			CloudProvider: "cloudian",
			Name:          "cloudian-dmaas-test",
			AccessID:      cloudianAccessID,
			SecretKey:     cloudianSecretKey,
		},
		"cloudian bucket with duplicate name": {
			CloudProvider: "cloudian",
			Name:          "cloudian-dmaas-test",
			AccessID:      cloudianAccessID,
			SecretKey:     cloudianSecretKey,
		},
		"cloudian bucket with name length less than 6 character": {
			CloudProvider: "cloudian",
			Name:          "cldm",
			AccessID:      cloudianAccessID,
			SecretKey:     cloudianSecretKey,
		},
		"cloudian bucket with name length more than 25 character": {
			CloudProvider: "cloudian",
			Name:          "cloudian-dmaas-test-backup-restore-cred",
			AccessID:      cloudianAccessID,
			SecretKey:     cloudianSecretKey,
		},
		"cloudian bucket without name": {
			CloudProvider: "cloudian",
			AccessID:      cloudianAccessID,
			SecretKey:     cloudianSecretKey,
			IsError:       true,
		},
		"cloudian bucket without accessID": {
			CloudProvider: "cloudian",
			Name:          "cloudian-cred-neg",
			SecretKey:     cloudianSecretKey,
			IsError:       true,
		},
		"cloudian bucket without secretKey": {
			CloudProvider: "cloudian",
			Name:          "cloudian-cred-neg",
			AccessID:      cloudianAccessID,
			IsError:       true,
		},
		"cloudian bucket without name, accessID & secretKey": {
			CloudProvider: "cloudian",
			IsError:       true,
		},
		"minio bucket with all details": {
			CloudProvider: "minio",
			Name:          "minio-dmaas-test",
			AccessID:      minioAccessID,
			SecretKey:     minioSecretKey,
		},
		"minio bucket with duplicate name": {
			CloudProvider: "minio",
			Name:          "minio-dmaas-test",
			AccessID:      minioAccessID,
			SecretKey:     minioSecretKey,
		},
		"minio bucket with name length less than 6 character": {
			CloudProvider: "minio",
			Name:          "minio",
			AccessID:      minioAccessID,
			SecretKey:     minioSecretKey,
		},
		"minio bucket with name length more than 25 character": {
			CloudProvider: "minio",
			Name:          "minio-dmaas-test-backup-restore-cred",
			AccessID:      minioAccessID,
			SecretKey:     minioSecretKey,
		},
		"minio bucket without name": {
			CloudProvider: "minio",
			AccessID:      minioAccessID,
			SecretKey:     minioSecretKey,
			IsError:       true,
		},
		"minio bucket without accessID": {
			CloudProvider: "minio",
			Name:          "minio-cred-neg",
			SecretKey:     minioSecretKey,
			IsError:       true,
		},
		"minio bucket without secretKey": {
			CloudProvider: "minio",
			Name:          "minio-cred-neg",
			AccessID:      minioAccessID,
			IsError:       true,
		},
		"minio bucket without name, accessID & secretKey": {
			CloudProvider: "minio",
			IsError:       true,
		},
	}
	for name, mock := range tests {
		name := name // Pin it
		mock := mock // Pin it
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			cred := NewCredentialsAdder(CredentialAdding{
				CloudName: mock.CloudProvider,
				CredName:  mock.Name,
				Username:  mock.AccessID,
				Password:  mock.SecretKey,
			})
			err := cred.AddCloudCredential()
			if !mock.IsError && err != nil {
				t.Fatalf("Failed to add cloud credential with error \nerror=%s ",
					err)
			}
			if mock.IsError && err == nil {
				t.Fatalf("Expected error got none: \nerror=%s", err)
			}
		})
	}
}
