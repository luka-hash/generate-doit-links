// Copyright © 2024- Luka Ivanović
// This code is licensed under the terms of the MIT Licence (see LICENCE for details).

package main

import (
	"encoding/json"
	"flag"
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
	additional_command := flag.String("command", "", "additional command to pass to yt-dlp")
	output := flag.String("output", "playlist_links", "output file to put the playlist links in")
	flag.Parse()
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: ,generate-doit-links [OPTIONS] <url> [<url>...]")
		os.Exit(1)
	}
	for _, url := range args {
		if !strings.HasPrefix(url, "https://") {
			fmt.Fprintln(os.Stderr, "Only 'https' links are supported.")
			continue
		}
		arguments := []string{"-J", "--flat-playlist"}
		if (*additional_command) != "" {
			arguments = append(arguments, strings.Fields(*additional_command)...)
		}
		arguments = append(arguments, url)
		cmd := exec.Command("yt-dlp", arguments...)
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
		f, err := os.OpenFile(*output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o0640)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			f = os.Stdout
		}
		defer f.Close()
		for _, entry := range playlist.Entries {
			fmt.Fprintf(f, "%s - %s %s\n", entry.Channel, entry.Title, entry.URL)
		}
		fmt.Fprint(f, "\n")
	}
}
