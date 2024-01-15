package main

import (
	"image/png"
	"image/color/palette"
	"net/http"
	"strconv"

	"genart/ccxy"
)

func sendError(w http.ResponseWriter, msg string) {
	w.WriteHeader(400)
	w.Write([]byte(msg))
	return
}

func ccxyHandler(w http.ResponseWriter, r *http.Request) {
	size, err := strconv.Atoi(r.PathValue("size"))
	if err != nil {
		sendError(w, "Invalid size value")
		return
	}

	constant, err := strconv.Atoi(r.PathValue("constant"))
	if err != nil {
		sendError(w, "Invalid constant value")
		return
	}

	colors, err := strconv.Atoi(r.PathValue("colors"))
	if err != nil {
		sendError(w, "Invalid colors value")
		return
	}

	newImage := ccxy.CcxyStruct{}
	newImage.Init(size, constant, colors, palette.WebSafe)
	newImage.Draw()

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, newImage.Image)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(
		"GET /ccxy/size/{size}/constant/{constant}/colors/{colors}/",
		ccxyHandler,
	)

	http.ListenAndServe("127.0.0.1:8090", mux)
}
