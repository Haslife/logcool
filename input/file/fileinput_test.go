package fileinput

import (
	"fmt"
	"testing"

	"logcool/utils"
)

func Test_DefaultInputConfig(t *testing.T) {
	DefaultInputConfig()
}

func Test_InitHandler(t *testing.T) {
	config := utils.ConfigRaw{}
	co, err := InitHandler(&config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(co)
}

func Test_Start(t *testing.T) {
	// config := DefaultInputConfig()
	// config.Start()
}
