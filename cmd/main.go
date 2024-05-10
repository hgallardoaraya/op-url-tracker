package main

import (
	"bufio"
	"fmt"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching url: ", err)
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body: ", err)
		return "", err
	}

	return string(data), nil
}

// checks for a specific className in html
func containsClassName(html string, className string) bool {
	if !strings.Contains(html, className) {
		return false
	}
	return true
}

func sendTgMessage(token string, text string, strChatId string) error {
	chatId, err := strconv.ParseInt(strChatId, 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing CHAT_ID: %v", err)
	}

	bot, err := tgBotApi.NewBotAPI(token) // Reemplaza "TOKEN" con tu token de bot
	if err != nil {
		return fmt.Errorf("error creating bot: %s", err)
	}

	bot.Debug = true

	// sets message
	msg := tgBotApi.NewMessage(chatId, text)

	// sends message through Telegram
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}

	return nil
}

func searchClassnameInHtml(config map[string]string) {
	url := config["URL"]
	className := config["CLASSNAME"]
	token := config["TOKEN"]
	text := config["MESSAGE"]
	chatId := config["CHAT_ID"]

	html, err := getHTML(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if containsClassName(html, className) {
		err := sendTgMessage(token, text, chatId)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}

func writeStringWrapper(file *os.File, str string) {
	_, err := file.WriteString(str)
	if err != nil {
		fmt.Println("error writing to file: ", err)
		os.Exit(1)
	}
}

func writeDefaultConfig(file *os.File) {
	writeStringWrapper(file, "URL=https://www.montasycomicsnyc.com/mtg-events/mtg-modern-horizons-3-launch-weekend-fri-614-sun-616\n")
	writeStringWrapper(file, "CLASSNAME=product-grid__item-overlay\n")
	writeStringWrapper(file, "TOKEN=6939137646:AAHumX9pZKc2r6bbHDl0k-4H0LMBwmnRcrM\n")
	writeStringWrapper(file, "CHAT_ID=6669733397\n")
	writeStringWrapper(file, "MESSAGE=An event has appeared")
}

// reads "config" file and puts its key/values in a map
func readConfig(configPath string) map[string]string {
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("error opening file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var config = make(map[string]string)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		if len(line) != 2 {
			fmt.Println("error parsing line: ", line)
			os.Exit(1)
		}
		key := line[0]
		value := line[1]
		config[key] = value
	}

	if err := scanner.Err(); err != nil {
		fmt.Print("error scanning file: ", err)
		os.Exit(1)
	}

	return config
}

func getConfigPath(configFileName string) string {
	// Obtiene la ruta del ejecutable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("error obteniendo la ruta del ejecutable:", err)
		os.Exit(1)
	}
	execDir := filepath.Dir(exePath)
	return filepath.Join(execDir, configFileName)
}

func getConfig(configPath string) map[string]string {
	var config map[string]string
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		fmt.Println("configuration file \"config\" does not exist, creating default configuration...")
		file, err := os.Create(configPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		writeDefaultConfig(file)
		config = readConfig(configPath)
	} else if err != nil {
		fmt.Println("error reading configuration file")
		os.Exit(1)
	} else {
		config = readConfig(configPath)
	}
	return config
}

func main() {
	configFileName := "config"
	configPath := getConfigPath(configFileName)
	config := getConfig(configPath)

	if config == nil {
		fmt.Println("error reading configuration file")
		os.Exit(1)
	}

	searchClassnameInHtml(config)
}
