// Copyright © 2024- Luka Ivanović
// This code is licensed under the terms of the MIT Licence (see LICENCE for details).

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Playlist struct {
	Entries []struct {
		URL     string `json:"url"`
		Title   string `json:"title"`
		Channel string `json:"channel"`
	} `json:"entries"`
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: ,generate-doit-links <url> [<url>...]")
		os.Exit(1)
	}
	for _, arg := range args {
		if !strings.HasPrefix(arg, "https://") {
			fmt.Fprintln(os.Stderr, "Only 'https' links are supported.")
			continue
		}
		cmd := exec.Command("yt-dlp", "-J", "--flat-playlist", arg)
		data, err := cmd.Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		var playlist Playlist
		err = json.Unmarshal(data, &playlist)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		for _, entry := range playlist.Entries {
			fmt.Printf("%s - %s %s\n", entry.Channel, entry.Title, entry.URL)
		}
		fmt.Println()
	}
}
