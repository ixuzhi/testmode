package testmode

import "fmt"

func Hi(name string) string {
	return fmt.Sprintf("hi,testmode:%s,%s,version tag", name, version())

}
