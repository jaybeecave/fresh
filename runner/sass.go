package runner

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/wellington/go-libsass"
)

func serveSass(sassFilepath string) {
	sassLog("processing " + sassFilepath)
	// wd, _ := os.Getwd()
	// requestFile := filepath.Join(wd, reader.URL.Path)
	// sourceFile := strings.TrimSuffix(requestFile, ".css") + ".scss"

	if path.Ext(sassFilepath) == ".scss" {
		cssFilepath := strings.TrimSuffix(sassFilepath, ".scss") + ".css"
		cssFilepath = strings.Replace(cssFilepath, "/sass/", "/stylesheets/", 1)
		// sourceFile := filepath.Join(wd, reader.URL.Path)

		// if path.Ext(reader.URL.Path) != ".css" {
		// 	return
		// }

		fmt.Println(sassFilepath)

		fi, err := os.Open(sassFilepath)
		if err != nil {
			panic(err)
		}
		defer fi.Close()
		out, err := os.Create(cssFilepath)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		comp, err := libsass.New(out, fi,
			libsass.Path(sassFilepath),
			libsass.OutputStyle(libsass.COMPRESSED_STYLE),
			libsass.IncludePaths([]string{filepath.Join(wd)}),
			// libsass.LineComments(true),
			// libsass.Comments(true),
			libsass.SourceMap(true, cssFilepath+".map"),
			// libsass.OutputStyle(1),
		)
		if err != nil {
			log.Printf("ERROR: libsass error: %#v", err.Error())
			panic(err)
		}

		if err = comp.Run(); err != nil {
			log.Printf("ERROR: libsass compile error: %#v", err.Error())
			panic(err)
		}
		out.Close()
	}
}

// // Compile sass files using wellington
// func buildSass() {
// 	const sassdir = "assets/stylesheets/"

// 	// open input sass/scss file to be compiled
// 	fi, err := os.Open(sassdir + "test.scss")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer fi.Close()

// 	// create output css file
// 	fo, err := os.Create(sassdir + "style.css")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer fo.Close()

// 	// options for compilation
// 	p := libsass.IncludePaths([]string{sassdir})
// 	s := libsass.OutputStyle(libsass.COMPRESSED_STYLE)
// 	m := libsass.SourceMap(true, sassdir+"style.css.map")

// 	// create a new compiler with options
// 	comp, err := libsass.New(fo, fi, p, s, m)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// start compile
// 	if err := comp.Run(); err != nil {
// 		panic(err)
// 	}
// }
