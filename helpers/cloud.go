package helpers

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func DownloadFiles() error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	it := client.Bucket("thangtd1").Objects(ctx, &storage.Query{
		Prefix: "OCRs/",
	})

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return err
		}

		if strings.Contains(attrs.Name, ".json") {
			rc, err := client.Bucket("thangtd1").Object(attrs.Name).NewReader(ctx)
			if err != nil {
				return err
			}
			defer rc.Close()

			f, err := os.Create("./database/data/" + attrs.Name[5:])
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err = f.ReadFrom(rc); err != nil {
				return err
			}

			if err := f.Sync(); err != nil {
				return err
			}

			if err := f.Close(); err != nil {
				return err
			}
		}
	}

	return nil

}
