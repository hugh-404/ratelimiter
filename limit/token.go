package limit

import "time"

type Token struct{}

type TokenGenerator struct {
	TPS     int32 // token per second
	channel chan Token
}

func NewTokenGenerator(tps int32) *TokenGenerator {
	if tps <= 0 {
		return nil
	}
	return &TokenGenerator{
		TPS:     tps,
		channel: make(chan Token),
	}
}

func (generator *TokenGenerator) StartGenerate() {
	timer := time.NewTicker(time.Second / time.Duration(generator.TPS))
	go func() {
		for {
			select {
			case <-timer.C:
				generator.channel <- struct{}{}
			}
		}
	}()
}

func (generator *TokenGenerator) GenerateToken() Token {
	select {
	case <-generator.channel:
		return Token{}
	}
}
