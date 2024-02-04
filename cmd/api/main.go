package main

import (
	"fmt"
	"image/png"
	"image/color/palette"
	"log"
	"net/http"
	"os"
	"strconv"

	"genart/formula"
)

func sendError(w http.ResponseWriter, msg string) {
	w.WriteHeader(400)
	w.Write([]byte(msg))
	return
}

func formulaHandler(w http.ResponseWriter, r *http.Request) {
	formulaName := r.PathValue("formulaName")
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

	newImage := formula.New()
	newImage.Init(formulaName, size, constant, colors, palette.WebSafe)
	newImage.Draw()

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, newImage.Image)
}

func main() {
	listenStr, found := os.LookupEnv("GENART_LISTEN")
	if !found {
		listenStr = ":8090"
	}

	mux := http.NewServeMux()
	mux.HandleFunc(
		"GET /{formulaName}/size/{size}/constant/{constant}/colors/{colors}/",
		formulaHandler,
	)

	fmt.Printf("Listening on %s\n", listenStr)
	log.Fatal(http.ListenAndServe(listenStr, mux))
}
