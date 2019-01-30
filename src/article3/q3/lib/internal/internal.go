package internal

import (
	"io"
	"fmt"
)

func Hello(w io.Writer, name string)  {
	fmt.Fprintf(w, "Hello , %s!\n",name)
}
