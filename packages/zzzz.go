package packages

import (
	"fmt"
	"time"
)

func init() {
	//completion.Wait()
	fmt.Printf("InitCount: %d, InitTime: %s\n", *initCount, time.Since(initStart).String())
}
