package main

import (
	"fmt"
	"github.com/gookit/cliapp"
)

func main() {
	app := cliapp.New(func(a *cliapp.App) {
		a.Name = "awesome kit"
		a.Version = "1.0.0"
		a.Description = "this is a tool for search awesome project."
		a.Logo.Text = fmt.Sprintf(`                                                    __   _ __
  ____ __      _____  _________  ____ ___  ___     / /__(_) /_
 / __ %s/ | /| / / _ \/ ___/ __ \/ __ %s__ \/ _ \   / //_/ / __/
/ /_/ /| |/ |/ /  __(__  ) /_/ / / / / / /  __/  / ,< / / /_
\__,_/ |__/|__/\___/____/\____/_/ /_/ /_/\___/  /_/|_/_/\__/`, "`", "`")
	})

	app.Run()
}

var repositories = map[string]string{
	// https://github.com/avelino/awesome-go
	// file: https://raw.githubusercontent.com/avelino/awesome-go/master/README.md
	"go": "avelino/awesome-go",
}
