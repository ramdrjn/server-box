package serverbox

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func Initialize(){
	fmt.Println("Server box initialized")
	logger = log.New(os.Stdout, "serverbox: ", 0)
	logger.Println("Server box initialized")
}
