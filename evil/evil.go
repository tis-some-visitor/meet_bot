package evil

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/lib/pq"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//returns true if the Class of a *pq.Error (if accessable) is 08 (Connection Exception) or 40 (Transaction Rollback)
func ShouldRetryPostgres(err error) bool {

	if err, ok := err.(*pq.Error); ok {

		if err.Code.Class() == "08" || err.Code.Class() == "40" {

			return true
		}
	}

	return false
}

//retries the function 10 times without delays
func Retry(f func(int64) (string, error), in int64) (string, error) {

	var err error
	attempts := 10

	for i := 0; i < attempts; i++ {

		if i > 0 {

			log.Println("retrying after error: ", err)

		}

		val, err := f(in)

		if err == nil {

			return val, nil

		}
	}

	return "", fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func RetryPostgres(f func() error) (err error) {

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

		err = f()

		if err == nil {

			return nil

		} else if ShouldRetryPostgres(err) == false {

			return err
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

//writes to log message "Problem in [function], \nquery: [query]; \nerror: [err]"
func ProblemIn(function, query string, err error) {

	log.Printf("Problem in %s, \nquery: %s; \nerror: %v", function, query, err)

}
