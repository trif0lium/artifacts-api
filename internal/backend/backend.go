package backend

import "gocloud.dev/blob"

type CacheBackend interface {
	PutArtifact(hash string) error
	FetchArtifact(hash string) error
}

type cacheBackend struct {
	bucket *blob.Bucket
}

func NewCacheBackend(bucket *blob.Bucket, bucketFolder string) CacheBackend {
	prefixedBucket := blob.PrefixedBucket(bucket, bucketFolder)
	return &cacheBackend{bucket: prefixedBucket}
}

func (backend *cacheBackend) PutArtifact(hash string) error {
	return nil
}

func (backend *cacheBackend) FetchArtifact(hash string) error {
	return nil
}
