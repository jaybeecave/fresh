package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wellington/go-libsass"
)

func buildSass(sassFilepath string) (string, bool) {
	// get just the filename
	fileName := getFileName(sassFilepath)
	if fileName == "_variables.scss" {
		sassFilepath = strings.Replace(sassFilepath, "_variables.scss", "main.scss", -1)
		sassLog("processing main.scss in place of " + fileName)
	} else {
		sassLog("processing " + fileName)
	}
	wd, err := os.Getwd()
	if err != nil {
		return err.Error(), false
	}

	cssFilepath := strings.TrimSuffix(sassFilepath, ".scss") + ".css"
	cssFilepath = strings.Replace(cssFilepath, "/scss/", "/css/", 1)
	sassLog("scss input " + sassFilepath)
	sassLog("css output" + cssFilepath)

	fi, err := os.Open(sassFilepath)
	if err != nil {
		return err.Error(), false
	}
	defer fi.Close()
	out, err := os.Create(cssFilepath)
	if err != nil {
		return err.Error(), false
	}
	defer out.Close()

	comp, err := libsass.New(out, fi,
		libsass.Path(sassFilepath),
		libsass.OutputStyle(libsass.COMPRESSED_STYLE),
		libsass.IncludePaths([]string{filepath.Join(wd)}),
		// libsass.LineComments(true),
		// libsass.Comments(true),
		libsass.SourceMap(true, cssFilepath+".map", ""),
		// libsass.OutputStyle(1),
	)
	if err != nil {
		errorMessage := "ERROR: libsass error: " + err.Error()
		fmt.Println(sassFilepath)
		return errorMessage, false
	}

	if err = comp.Run(); err != nil {
		errorMessage := "ERROR: libsass compile error: " + err.Error()
		fmt.Println(sassFilepath)
		return errorMessage, false
	}
	return "", true
}
