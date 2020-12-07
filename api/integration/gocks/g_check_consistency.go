package gocks

import (
	. "github.com/loveandpeople-DAG/goClient/api"
	. "github.com/loveandpeople-DAG/goClient/api/integration/samples"
	. "github.com/loveandpeople-DAG/goClient/trinary"
	"gopkg.in/h2non/gock.v1"
	"strings"
)

func init() {
	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(CheckConsistencyCommand{Command: Command{CheckConsistencyCmd}, Tails: DefaultHashes()}).
		Reply(200).
		JSON(CheckConsistencyResponse{State: true, Info: ""})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(CheckConsistencyCommand{
			Command: Command{CheckConsistencyCmd},
			Tails:   append(DefaultHashes(), strings.Repeat("C", 81)),
		}).
		Reply(200).
		JSON(CheckConsistencyResponse{State: false, Info: "test response"})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(CheckConsistencyCommand{
			Command: Command{CheckConsistencyCmd},
			Tails:   Hashes{Bundle[0].Hash},
		}).
		Reply(200).
		JSON(CheckConsistencyResponse{State: true, Info: ""})
}
