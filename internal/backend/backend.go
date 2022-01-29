package backend

import (
	"context"
	"io"
	"strings"

	"gocloud.dev/blob"
)

type CacheBackend interface {
	PutArtifact(hash string, reader io.Reader) error
	FetchArtifact(hash string) (io.Reader, error)
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

func (backend *cacheBackend) PutArtifact(hash string, reader io.Reader) error {
	writeCtx, cancelWrite := context.WithCancel(context.TODO())
	defer cancelWrite()

	w, err := backend.bucket.NewWriter(writeCtx, hash, nil)
	defer w.Close()

	if _, err := io.Copy(w, reader); err != nil {
		cancelWrite()
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (backend *cacheBackend) FetchArtifact(hash string) (io.Reader, error) {
	readCtx, cancelRead := context.WithCancel(context.TODO())
	defer cancelRead()

	r, err := backend.bucket.NewReader(readCtx, hash, nil)
	defer r.Close()
	if err != nil {
		cancelRead()
		return nil, err
	}

	return r, nil
}
