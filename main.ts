const link = "<your-link>" // Deno.args[1] or something

const command = new Deno.Command("yt-dlp", {
	args: [
		"-J",
		"--flat-playlist",
		link,
	],
});

const result = await command.output();

const text = new TextDecoder().decode(result.stdout);

const data = JSON.parse(text);

if (data.playlist_count!==data.entries.length){
	throw new Error("FIXME");
}

console.log(`-- ARTIST - ${data.title} - ${data.release_year}`)
data.entries.forEach((entry: { channel: any; title: any; url: any; }) => {
	console.log(`${entry.channel} - ${entry.title} ${entry.url}`)
});

