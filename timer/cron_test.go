package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestStartCronTimer(t *testing.T) {
	StartCronTimer("*/1 * * * *", func() {
		fmt.Println(time.Now().String())
	})
}
