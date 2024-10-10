package main

import (
	"fmt"

	"github.com/bazel/golang/test/utils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	defer func() {
		if recover() != nil {
			log.Error().Msg("Well, I tried to divide by Zero")
		}
	}()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	fmt.Println("Some Magical Division happening here: , I have no idea how the lib works to be frank")
	utils.DivideNumbers(98237498234, 2134)

	fmt.Println("Some More Magical Division happening here: ")
	utils.DivideNumbers(238, 0)
}