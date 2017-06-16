package main

import (
	"fmt"
	"github.com/snahelou/firstGoProject/second"
	"os"
	log "github.com/Sirupsen/logrus"

)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02T15:04:05.999999-0700",
	})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

}
func main() {
	// Upper case = public func
	fmt.Print("Let's rock!\n")
	second.Print()

	myObj := second.TestType{
		TestInt: 10,
		TestString: "Ten",
	}

	myObj.Print()

	myObj.Loop()

	err:= second.CallWS()
	if err != nil {
		fmt.Println(err.Error())
	}
}
