package godori_tcpnet

import (
	"fmt"
	"os"
)

const (
	LOG_LV_TRACE int = iota
	LOG_LV_DEBUG
	LOG_LV_INFO
	LOG_LV_WARN
	LOG_LV_ERR
	LOG_LV_FATAL
)

var logLvStr = [6]string{"trace", "debug", "info", "warn", "err", "fatal"}

var (
	OutputLog = _emptyExportLog
)

func logTrace(userID string, sessionID uint64, msg string) {
	OutputLog(LOG_LV_TRACE, userID, sessionID, msg)
}
func logDebug(userID string, sessionID uint64, msg string) {
	OutputLog(LOG_LV_DEBUG, userID, sessionID, msg)
}
func logInfo(userID string, sessionID uint64, msg string) {
	OutputLog(LOG_LV_INFO, userID, sessionID, msg)
}
func logWarn(userID string, sessionID uint64, msg string) {
	OutputLog(LOG_LV_WARN, userID, sessionID, msg)
}
func logErr(userID string, sessionID uint64, msg string) {
	OutputLog(LOG_LV_ERR, userID, sessionID, msg)
}
func logFatal(userID string, sessionID uint64, msg string) {
	OutputLog(LOG_LV_FATAL, userID, sessionID, msg)
}

func _emptyExportLog(level int, userID string, sessionID uint64, msg string) {
	if level < _logLv {
		return
	}

	fmt.Fprintf(os.Stdout, "[%s] %s\n", logLvStr[level], msg)
}

var _logLv int
