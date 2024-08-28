package options

const (
	// default config file name
	DefBotConfFile string = "tg_bot.conf"
	// default log file name
	DefBotLogFile string = "tg_bot.log"

	// default logger level
	DefBotLogLvl string = "Error"
	// default logger name
	DefBotLogName string = "TGBOT"
	// default log file size in megabytes
	DefBotLogSizeMb uint = 0
	// default log file size in days
	DefBotLogSizeDay uint = 0
	// default token
	DefBotToken string = "0000000000:AAES000000000000000000-ae0000000000"
	// default token regular expression mask
	DefBotTokenRegExp string = `\A\d{10}:\w{22}-\w{12}\z`

	// default description for config file
	DefBotConfDescr string = "The configuration file for the \"Telegram Bot\" process"
)

// Set default values to configuration structure
func (obj *TgBotObj) SetDefaultConf() {
	obj.conf.SrvHost = DefSrvHost
	obj.conf.SrvPort = DefSrvPort
	obj.conf.Token = DefBotToken
	obj.conf.LogLevel = DefBotLogLvl
	obj.conf.LogFile = DefBotLogFile
	obj.conf.ConfFile = DefBotConfFile
}
