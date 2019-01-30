package lib12

import "fmt"

func GetElement(container interface{}) (elem string, err error) {
	switch t := container.(type) {
	case []string :
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", container)
		return

	}
	return
}
