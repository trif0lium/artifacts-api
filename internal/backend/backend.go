package backend

type CacheBackend interface {
	PutArtifact(hash string, bucketFolder string) error
	FetchArtifact(hash string, bucketFolder string) error
}

type cacheBackend struct{}

func NewCacheBackend() CacheBackend {
	return &cacheBackend{}
}

func (backend *cacheBackend) PutArtifact(hash string, bucketFolder string) error {
	return nil
}

func (backend *cacheBackend) FetchArtifact(hash string, bucketFolder string) error {
	return nil
}
