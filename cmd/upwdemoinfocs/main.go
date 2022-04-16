package main

import (
	"fmt"
	"os"

	ex "github.com/markus-wa/demoinfocs-golang/v2/examples"
	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

// Run like this: go run cmd/mine/main.go -demo ./default.dem
func main() {
	f, err := os.Open(ex.DemoPathFromArgs())
	checkError(err)

	defer f.Close()

	p := demoinfocs.NewParser(f)
	defer p.Close()

	fmt.Printf("-- Players damaged because of fall --\n\n\n")
	fmt.Println("user_id, name, last_alive_pos")

	// Register handler on generic game events for manually select event by name
	p.RegisterEventHandler(func(e events.GenericGameEvent) {
		if e.Name == "player_falldamage" {
			userid := int(e.Data["userid"].GetValShort())
			player := p.GameState().Participants().AllByUserID()[userid]

			// for simplicity, print with csv form instead of using actual csv formatter
			fmt.Printf("%d,\"%s\",\"%s\"\n", player.UserID, player.Name, player.LastAlivePosition)
		}
	})

	// Parse to end
	err = p.ParseToEnd()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
