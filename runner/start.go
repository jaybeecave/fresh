package runner

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	startChannel chan string
	stopChannel  chan bool
	mainLog      logFunc
	watcherLog   logFunc
	sassLog      logFunc
	runnerLog    logFunc
	buildLog     logFunc
	appLog       logFunc
)

func flushEvents() {
	for {
		select {
		case eventName := <-startChannel:
			mainLog("receiving event %s", eventName)
		default:
			return
		}
	}
}

func start() {
	buildDelay := time.Duration(settings.BuildDelay) * time.Millisecond

	started := false

	go func() {
		for {
			eventName := <-startChannel
			filePath := getFilePath(eventName)
			fileName := getFileName(filePath)

			mainLog("receiving first event %s", eventName)

			mainLog("sleeping for %d milliseconds", buildDelay/time.Millisecond)
			time.Sleep(buildDelay)

			mainLog("flushing events")
			flushEvents()

			mainLog("Started! (%d Goroutines)", runtime.NumGoroutine())
			err := removeBuildErrorsLog()
			if err != nil {
				mainLog(err.Error())
			}

			// if its sass do it here
			sassLog("fullpath:" + filePath)
			sassLog("filename:" + fileName)
			if filepath.Ext(filePath) == ".scss" {
				errorMessage, ok := buildSass(filePath)
				if ok {
					sassLog("processed:" + fileName)
				} else {
					sassLog(errorMessage)
				}
				continue
			}

			errorMessage, ok := build()

			if !ok {
				createBuildErrorsLog(errorMessage)
			} else {
				if started {
					stopChannel <- true
				}
				if run() {
					started = true
				}
			}
		}
	}()
}

func init() {
	startChannel = make(chan string, 1000)
	stopChannel = make(chan bool)
}

func initLogFuncs() {
	mainLog = newLogFunc("main")
	watcherLog = newLogFunc("watcher")
	sassLog = newLogFunc("sass")
	runnerLog = newLogFunc("runner")
	buildLog = newLogFunc("build")
	appLog = newLogFunc("app")
}

// Start watches for file changes in the root directory.
// After each file system event it builds and (re)starts the application.
func Start(confFile, buildArgs *string, runArgs []string, buildPath, outputBinary, tmpPath *string, watchList, excludeList Multiflag) {
	os.Setenv("DEV_RUNNER", "1")
	initLimit()
	initLogFuncs()
	err := initSettings(confFile, buildArgs, runArgs, buildPath, outputBinary, tmpPath, watchList, excludeList)
	if err != nil {
		logger.Fatalf("Failed to start: %v", err)
		return
	}
	initFolders()
	watch()
	start()
	startChannel <- "/"

	select {}
}

func getFilePath(eventName string) (path string) {
	path = strings.Split(eventName, "\": MODIFY")[0]
	path = strings.Replace(path, `"`, "", 1)
	return path
}
func getFileName(path string) (fileName string) {
	parts := strings.Split(path, `/`)
	lastIndex := len(parts) - 1
	return parts[lastIndex]
}
