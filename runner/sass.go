package runner

import (
	"os"

	"github.com/wellington/go-libsass"
)

func serveSass() {
	// wd, _ := os.Getwd()
	// requestFile := filepath.Join(wd, reader.URL.Path)
	// sourceFile := strings.TrimSuffix(requestFile, ".css") + ".scss"

	// if path.Ext(reader.URL.Path) == ".scss" {
	// 	requestFile = strings.TrimSuffix(requestFile, ".scss") + ".css"
	// 	requestFile = strings.Replace(requestFile, "/sass/", "/stylesheets/", 1)
	// 	sourceFile = filepath.Join(wd, reader.URL.Path)

	// 	// if path.Ext(reader.URL.Path) != ".css" {
	// 	// 	return
	// 	// }

	// 	fmt.Println(sourceFile)

	// 	fi, err := os.Open(sourceFile)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer fi.Close()
	// 	out, err := os.Create(requestFile)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer out.Close()

	// 	comp, err := libsass.New(out, fi,
	// 		libsass.Path(sourceFile),
	// 		libsass.OutputStyle(libsass.COMPRESSED_STYLE),
	// 		libsass.IncludePaths([]string{filepath.Join(wd)}),
	// 		// libsass.LineComments(true),
	// 		// libsass.Comments(true),
	// 		libsass.SourceMap(true, requestFile+".map"),
	// 		// libsass.OutputStyle(1),
	// 	)
	// 	if err != nil {
	// 		log.Printf("ERROR: libsass error: %#v", err.Error())
	// 		http.Error(w, err.Error(), 500)
	// 		return
	// 	}

	// 	if err = comp.Run(); err != nil {
	// 		log.Printf("ERROR: libsass compile error: %#v", err.Error())
	// 		http.Error(w, err.Error(), 500)
	// 		return
	// 	}
	// 	out.Close()
	// }

	// // buffer := []byte("asdfhfhkj")
	// // buffer := bytes.NewBuffer(nil)
	// // io.Copy(buffer, out)

	// buffer, err := ioutil.ReadFile(requestFile)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// }

	// w.Header().Set("Content-Type", "text/css")
	// w.Write(buffer)
}

// Compile sass files using wellington
func buildSass() {
	const sassdir = "assets/stylesheets/"

	// open input sass/scss file to be compiled
	fi, err := os.Open(sassdir + "test.scss")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	// create output css file
	fo, err := os.Create(sassdir + "style.css")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	// options for compilation
	p := libsass.IncludePaths([]string{sassdir})
	s := libsass.OutputStyle(libsass.COMPRESSED_STYLE)
	m := libsass.SourceMap(true, sassdir+"style.css.map")

	// create a new compiler with options
	comp, err := libsass.New(fo, fi, p, s, m)
	if err != nil {
		panic(err)
	}

	// start compile
	if err := comp.Run(); err != nil {
		panic(err)
	}
}
