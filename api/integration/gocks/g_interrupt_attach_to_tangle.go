package gocks

import (
	. "github.com/loveandpeople-DAG/goClient/api"
	"gopkg.in/h2non/gock.v1"
)

func init() {
	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(InterruptAttachToTangleCommand{Command: Command{InterruptAttachToTangleCmd}}).
		Reply(200)
}
