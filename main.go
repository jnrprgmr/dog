package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andreykaipov/goobs"
	"github.com/jnrprgmr/dog/internal/rest/handlers"
	"github.com/jnrprgmr/dog/pkg/obs"
	"github.com/jnrprgmr/dog/pkg/twitch"
	"github.com/nicklaw5/helix"
)

func main() {
	obsCli, err := goobs.New("localhost:4455", goobs.WithPassword("test123"))
	if err != nil {
		log.Fatal(err)
	}
	defer obsCli.Disconnect()
	obs := obs.New(obsCli)
	client_id := os.Getenv("CLIENT_ID")
	client_secret := os.Getenv("CLIENT_SECRET")
	twitchCli, err := helix.NewClient(&helix.Options{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURI:  "http://localhost:8080/twitch/auth",
	})
	if err != nil {
		panic("error making twitch client: " + err.Error())
	}
	twitch := twitch.New(twitchCli)
	fmt.Println(twitch)
	h := handlers.New(twitch)
	http.HandleFunc("/twitch", h.TwitchHandler)
	http.HandleFunc("/twitch/update", h.TwitchUpdateHandler)
	http.HandleFunc("/twitch/auth", h.TwitchAuthHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
	//twitch.ChangeStreamTitle("jnrprgmr", "Making Bots in Golang")
	// resp, err := twitchCli.EditChannelInformation(&helix.EditChannelInformationParams{
	// 	BroadcasterID:       "123456",
	// 	GameID:              "456789",
	// 	BroadcasterLanguage: "en",
	// 	Title:               "Your stream title",
	// 	Delay:               0,
	// })
	// if err != nil {
	// 	// handle error
	// }

	// game, err := twitchCli.GetGames(&helix.GamesParams{
	// 	Names: []string{"Dota 2", "Software And Game Development"},
	// })
	// if err != nil {
	// 	panic("another one")
	// }

	// fmt.Printf("%+v\n", game)

	// user, err := twitchCli.GetUsers(&helix.UsersParams{
	// 	Logins: []string{"jnrprgmr"},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v\n", user)
	obs.SetTask("Change Twitch Title")
}
