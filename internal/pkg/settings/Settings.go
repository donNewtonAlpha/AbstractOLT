package settings

var debug = false

/*
SetDebug - sets debug setting
*/
func SetDebug(logDebug bool) {
	debug = logDebug
}

/*
GetDebug returns debug setting
*/
func GetDebug() bool {
	return debug
}
