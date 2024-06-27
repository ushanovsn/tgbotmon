package server

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
	"reflect"

	"github.com/ushanovsn/tgbotmon/internal/options"
)



// main initializing and configure function
func getConfig(srv *options.ServerObj) bool {
	log := srv.GetLogger()
	fPath := srv.GetConfFile()

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
	} else {
		log.OutError("Error while checking config file: " + err.Error())
		return false
	}

	// check lost config data in config file and add it to file
	if len(lostConf) > 0 {
		if !addConfig(fPath, lostConf) {
			log.OutError("Can't add parameters in config file")
			return false
		}
	}

	log.OutInfo("Config file succesfully processed")

	return true
}


// read config from file
func readConfig(filePath string) (map[string][]string, map[string][]string) {

	type User struct {
		Name  string `mytag:"MyName"`
		Email string `mytag:"MyEmail"`
	}
	
	u := User{"Bob", "bob@mycompany.com"}
	t := reflect.TypeOf(u)
	
	for _, fieldName := range []string{"Name", "Email"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}
		fmt.Printf("\nField: User.%s\n", fieldName)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
	}

	/*

	*/

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
