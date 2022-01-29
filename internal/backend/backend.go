package backend

import (
	"context"
	"strings"

	"gocloud.dev/blob"
)

type CacheBackend interface {
	PutArtifact(hash string) error
	FetchArtifact(hash string) error
}

type cacheBackend struct {
	bucket *blob.Bucket
}

func NewCacheBackend(bucket *blob.Bucket, bucketFolder string) CacheBackend {
	prefix := bucketFolder
	if !strings.HasSuffix(prefix, "/") {
		prefix = prefix + "/"
	}

	prefixedBucket := blob.PrefixedBucket(bucket, prefix)
	return &cacheBackend{bucket: prefixedBucket}
}

func (backend *cacheBackend) PutArtifact(hash string) error {
	writeCtx, cancelWrite := context.WithCancel(context.TODO())
	defer cancelWrite()

	w, err := backend.bucket.NewWriter(writeCtx, hash, nil)
	defer w.Close()

	if err != nil {
		return err
	}

	return nil
}

func (backend *cacheBackend) FetchArtifact(hash string) error {
	return nil
}
