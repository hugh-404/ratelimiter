package limit

type Bucket struct {
	size      int32
	channel   chan Token
	generator *TokenGenerator
}

func NewBucket(size int32, tokenGenerator *TokenGenerator) *Bucket {
	if size <= 0 || tokenGenerator == nil {
		return nil
	}
	return &Bucket{
		size:      size,
		channel:   make(chan Token, size),
		generator: tokenGenerator,
	}
}

func (bucket *Bucket) StartLimit() {
	bucket.generator.StartGenerate()
	go func() {
		for {
			bucket.channel <- bucket.generator.GenerateToken()
		}
	}()
}

func (bucket *Bucket) FetchToken() bool {
	// TODO 允许排队等待一段时间来获取令牌
	select {
	case <-bucket.channel:
		return true
	default:
		return false
	}
}
