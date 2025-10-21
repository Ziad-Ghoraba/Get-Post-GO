package main

import (
	"get_post_go_lang/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3030", nil)
}
