package envorsecretm

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type Config struct {
	ProjectId string
}

// Secret Manager
func (c *Config) Get(name string) string {
	if value := os.Getenv(name); value != "" {
		log.Printf("Getting value %s from Environment value\n", name)
		return value
	}
	log.Printf("Getting value %s from Secret Manager\n", name)
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v\n", err)
	}
	//var secret *secretmanager.Secret
	pathname := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", c.ProjectId, name)
	log.Println("pathname:" + pathname)
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: pathname,
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	value := string(result.Payload.Data)
	log.Println(value)
	return value
}
