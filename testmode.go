package testmode

import (
	"errors"
	"fmt"
)

func Hi(name string,lang string) (string,error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi:%s!",name),nil
	case "ch":
		return fmt.Sprintf("Hi:%s,%s!",name,lang),nil
	default:
		return "",errors.New("unknow lang:name")
	}


}
