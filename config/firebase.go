package config

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	firebase_storage "firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

type Firebase struct {
	Storage *firebase_storage.Client
}

func InitFirebase() *Firebase {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIAL_PATH"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	firebaseStorage, err := app.Storage(context.Background())
	if err != nil {
		panic(err)
	}

	return &Firebase{
		Storage: firebaseStorage,
	}
}

func (f *Firebase) UploadFile(ctx context.Context, file []byte, fileName string) (string, error) {
	bucket, err := f.Storage.Bucket(os.Getenv("FIREBASE_BUCKET"))
	if err != nil {
		return "", err
	}

	wc := bucket.Object(fileName).NewWriter(ctx)
	if _, err = wc.Write(file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	_, err = bucket.Object(fileName).Attrs(ctx)
	if err != nil {
		return "", err
	}

	link := fmt.Sprintf(os.Getenv("FIREBASE_IMAGE_URL"), fileName)
	return link, nil
}
