/*
Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
Contributors to this project, hereby assign copyright in this code to the project,
to be licensed under the same terms as the rest of the code.
*/
package main

import (
	"fmt"
	"github.com/rivo/tview"
)

// https://gist.github.com/rivo/2893c6740a6c651f685b9766d1898084
func main() {
	// Start the application.
	app = tview.NewApplication()
	// finder(os.Args[1])
	if err := app.Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
	}
}
