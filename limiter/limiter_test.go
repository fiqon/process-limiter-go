package limiter_test

import (
	"log"
	"testing"

	"github.com/grand-x/process-limiter-go/limiter"
)

func TestLimiter(t *testing.T) {
	vet := [...]string{"Hello", "World", "!"}
	limit := limiter.New(2)

	for _, str := range vet {
		limit.Execute(foo, str)
	}

	for i := 0; i < 10; i++ {
		limit.Execute(bar, i)
	}

	limit.Wait()
}

func foo(str interface{}) {
	log.Println(str)
}

func bar(d int) {
	log.Println(d)
}
