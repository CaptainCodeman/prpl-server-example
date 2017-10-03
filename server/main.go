package app

import (
	"os"
	"net/http"

	"github.com/captaincodeman/prpl-server-go"
)

func init() {
	version := os.Getenv("STATIC_VERSION")
	m, _ := prpl.New(
		prpl.WithVersion(version),
		prpl.WithRoot("./static"),
		prpl.WithConfigFile("./static/polymer.json"),
		prpl.WithRoutes(prpl.Routes{
			"/":      "src/my-view1.html",
			"/view1": "src/my-view1.html",
			"/view2": "src/my-view2.html",
			"/view3": "src/my-view3.html",
		}),
	)

	http.Handle("/", m)
}
