package renderers

import (
	"bytes"
	"html/template"
	"log"
)

func RenderHTML(data map[string]string) (bs []byte) {
	/* Probably want:
	* - SPA (react etc), packaged into the container
	* - gorilla mux SPA example
	* - JSON handlers for this struct (make it a struct and JSON serialse it) so it can be read by the SPA
	 */

	var b bytes.Buffer
	t, err := template.ParseFiles("data/templates/html.tpl")
	if err != nil {
		log.Fatalf("Failed to parse template html.tpl: %v", err)
	}
	t.Execute(&b, data)
	bs = b.Bytes()

	return
}
