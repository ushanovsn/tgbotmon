package server

import (
	"fmt"
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ushanovsn/golanglogger"
)

// TG Bot configuration structure
type botParam struct {
	TgToken string
	confFile string
	logLvl golanglogger.LoggingLevel
}

func InitBot() (botParam, error) {
	// TG Bot configuration data
	var params botParam
	// error receiving logger level from flags
	var lvlErr error

	// check flags at start
	flags := getFlags()

	// get log level (flag received or default value)
	params.logLvl, lvlErr = golanglogger.LoggingLevelValue(flags.logLevel)

	// start logger with flags init values
	log := golanglogger.NewSync(params.logLvl, flags.logFile)

	// this IF one who can change config file
	if flags.confFile != "" {
		params.confFile = flags.confFile
	} else {
		params.confFile = defConfFile
	}

	log.Out(fmt.Sprintf("Start to load configuration from \"%s\" file", params.confFile))


	if lvlErr != nil {
		// need to load level from config file

	}



	return params, nil
}

// returns default values for project configuration
func loadDefaultConfig() map[string][]string {
	return map[string][]string{
		"log_file_path":      {"tg_bot_log.log", "string"},
		"telegram_bot_token": {"0000000000:AAES000000000000000000-ae0000000000", "regexp", `\A\d{10}:\w{22}-\w{12}\z`},
		"log_split_size_b":   {"0", "int"},
		"log_split_day_time": {"0", "int"},
		"log_level":          {"Error", "string"},
	}
}

// main initializing and configure function
func ConfigAndInit(fPath string) bool {
	log = botlog.New(botlog.Error)
	log.OutInfo("* Starting ConfigAndInit()...")

	var fullConf map[string][]string
	var lostConf map[string][]string

	// check the file
	if _, err := os.Stat(fPath); err == nil {
		// file exist, now get config data
		fullConf, lostConf = readConfig(fPath)
	} else if errors.Is(err, os.ErrNotExist) {
		// file is not exist, need create it
		log.OutWarning("Config file is not exist")
		if !createConfigFile(fPath) {
			log.OutError("Config file can't be created")
			return false
		}
		fullConf = loadDefaultConfig()
		lostConf = loadDefaultConfig()
	} else {
		log.OutError("Error while checking config file: " + err.Error())
		// stopping process
		return false
	}

	// check lost config data in config file and add it to file
	if len(lostConf) > 0 {
		if !addConfig(fPath, lostConf) {
			log.OutError("Can't add parameters in config file")
			return false
		}
	}

	log.OutDebug("Starting apply Logger config")

	sizeLogFile, _ := strconv.Atoi(fullConf["log_split_size_b"][0])
	durationLogFile, _ := strconv.Atoi(fullConf["log_split_day_time"][0])
	log.SetFile(fullConf["log_file_path"][0], sizeLogFile, durationLogFile)

	lvl, _ := botlog.LoggingLevelValue(fullConf["log_level"][0])
	log.SetLevel(lvl)

	log.OutDebug("Applying permanent config data")

	CurConf.TgToken = fullConf["telegram_bot_token"][0]

	return true
}

// stop all process of tg bot
func StopBot() {
	log.Out("Logger stopping")
	log.StopLog()
}

// read config from file
func readConfig(filePath string) (map[string][]string, map[string][]string) {
	// load Defaults for config
	fConf := loadDefaultConfig()
	lostConf := loadDefaultConfig()

	log.OutInfo("Check config file...")
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.OutError("Cant open config file, err: " + err.Error())
		return fConf, lostConf
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	log.OutInfo("Reading config file...")

	// reading the file and apply configs
	for sc.Scan() {
		l := strings.TrimSpace(sc.Text())
		if l != "" && l[0:1] == "#" {
			continue
		}

		values := strings.Fields(l)
		if len(values) < 2 {
			continue
		}

		if def_val, ok := fConf[values[0]]; ok {
			param := strings.TrimSpace(strings.Replace(l, values[0], "", 1))

			// check and apply config data
			switch def_val[1] {
			case "regexp":
				if v, err := regexp.MatchString(def_val[2], param); err == nil && v {
					fConf[values[0]][0] = param
				} else if err == nil {
					log.OutWarning("Param " + values[0] + " is not matched! Loaded by default. Skipped bad value: " + param)
				} else {
					log.OutWarning("Param " + values[0] + " is wrong. Loaded by default. Skipped bad value: " + param + ". Error: " + err.Error())
				}
			case "int":
				if _, err := strconv.Atoi(param); err == nil {
					fConf[values[0]][0] = param
				} else {
					log.OutWarning("Param " + values[0] + " is wrong. Loaded by default. Skipped bad value: " + param + ". Error: " + err.Error())
				}
			case "bool":
				if _, err := strconv.ParseBool(param); err == nil {
					fConf[values[0]][0] = param
				} else {
					log.OutWarning("Param " + values[0] + " is wrong. Loaded by default. Skipped bad value: " + param + ". Error: " + err.Error())
				}
			case "float":
				if _, err := strconv.ParseFloat(param, 64); err == nil {
					fConf[values[0]][0] = param
				} else {
					log.OutWarning("Param " + values[0] + " is wrong. Loaded by default. Skipped bad value: " + param + ". Error: " + err.Error())
				}
			default:
				// is just string
				fConf[values[0]][0] = param
			}

			delete(lostConf, values[0])
		} else {
			continue
		}
	}

	// error if not EOF
	if err := sc.Err(); err != nil {
		log.OutError("Error while reading the config file, err: " + err.Error())
	}

	return fConf, lostConf
}

// add params to config file
func addConfig(filePath string, lostData map[string][]string) bool {
	// check file
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.OutError("Error while open config file for append data. Err: " + err.Error())
		return false
	}
	defer f.Close()

	log.OutDebug("Config file oppened for append new parameters")

	for i, d := range lostData {
		if _, err = f.WriteString(i + " \t" + d[0] + "\n"); err != nil {
			log.OutError("Error while write append data in config file. Err: " + err.Error())
		}
	}

	return true
}

// creating config file w header text
func createConfigFile(filePath string) bool {
	fRes := true
	headerConf := `# This config file for "telegram bot" project` + "\n\n"

	log.OutInfo("Creating config file...")

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.OutError("Error while creating config file. Err: " + err.Error())
		fRes = false
	}
	defer f.Close()

	if _, err = f.WriteString(headerConf); err != nil {
		log.OutError("Error while write data in config file. Err: " + err.Error())
		fRes = false
	}

	return fRes
}
