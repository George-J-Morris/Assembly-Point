package internal

import (
	"fmt"
	"os"

	"github.com/bep/godartsass/v2"
)

func TranspileBootstrapCss(src string, out string) {

	scssOpts := godartsass.Options{DartSassEmbeddedFilename: "./third_party/dart-sass/sass"}

	transpiler, err := godartsass.Start(scssOpts)
	if err != nil {
		fmt.Println(err)
	}

	myscsss, _ := os.ReadFile(src)
	scssString := string(myscsss)

	bsDir := "third_party/bootstrap/scss"

	result, err := transpiler.Execute(godartsass.Args{Source: scssString, IncludePaths: []string{bsDir}})
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create(fmt.Sprintf("%v/bootstrap.css", out))
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(result.CSS)

	transpiler.Close()
}
