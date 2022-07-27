package dbfellow

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//retries to Exec() the given query 5 times with exponential backoff
func RetryExec(query string) (err error) {

	attempts := 5
	sleep := time.Duration(1)

	for i := 0; i < attempts; i++ {

		if i > 0 {

			log.Println("retrying after error: ", err)
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			sleep *= 2
		}

		_, err = Db.Exec(query)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
