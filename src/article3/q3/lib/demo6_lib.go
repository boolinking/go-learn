package lib

import (
	"os"
	 in "article3/q3/lib/internal"
)

func Hello(name string)  {
	in.Hello(os.Stdout, name)
}
