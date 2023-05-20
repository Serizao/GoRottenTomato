package main

import (
	"github.com/Serizao/GoRottenTomato/module"
	"os"
)

func main()  {
	module.Parse(os.Args[1:])
}


