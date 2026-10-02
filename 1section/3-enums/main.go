package main

import "fmt"

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type LogLevel int

const (
	LogTrace LogLevel = iota
	LogDebug
	LogInfo
	LogWarn
	LogError
	LogFatal
)

func main() {

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	printLogLevel(LogTrace)
	printLogLevel(LogInfo)
	printLogLevel(LogWarn)
	printLogLevel(LogFatal)
	printLogLevel(10)

}

var levelNames = []string{"Trace", "Debug", "Info", "Warn", "Error", "Fatal"}

func (l LogLevel) String() string {
	if l < LogTrace || l > LogFatal {
		return "Unknown"
	}

	return levelNames[l]

}

func printLogLevel(level LogLevel) {
	fmt.Printf("LogLevel: %d, %s \n", level, level.String())
}
