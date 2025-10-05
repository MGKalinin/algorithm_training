package main

import (
	"fmt"
	"strings"
)

// Дан слайс из миллиона строк, нужно получить строку, склееную из этих строк.
//
//	Как правильно это сделать, какие есть подводные камни?
func main() {
	originalSlice := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://stackoverflow.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.reddit.com",
		"https://www.wikipedia.org",
		"https://www.medium.com",
		"https://www.netflix.com",
		"https://www.spotify.com",
		"https://www.dropbox.com",
		"https://www.slack.com",
		"https://www.notion.so",
		"https://www.figma.com",
		"https://www.trello.com",
		"https://www.airtable.com",
	}

	var builder strings.Builder
	for _, i := range originalSlice {
		builder.WriteString(i)
	}

	result := builder.String()
	fmt.Println(result)
	// правильно сделать через string.Builder
	/*подводные камни - аллокация памяти-строка это массив байт-при о добавлении строки в строку,
	если capacity не позволяет добавить-увеличение capacity в 2 раза до размера 1024, далее на одну четверть*/
}
