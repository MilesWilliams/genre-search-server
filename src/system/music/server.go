package music

import (
	"log"
	"net/http"
	"os"

	"github.com/MilesWilliams/music/src/system/router"

	"github.com/rs/cors"
)

// StartServer func
func StartServer() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-App-Token", "X-CSRF-Token", "Authorization", "Service-Worker-Allowed'"},
	})

	r := router.NewRouter()

	r.Init()

	http.Handle("/", ErrorChain.Then(c.Handler(r.Router)))

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Println(err)
	}

	os.Exit(0)
}
