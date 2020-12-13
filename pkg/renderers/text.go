package renderers

import (
	"bytes"
	"log"
	"text/template"
)

func RenderText(data map[string]string) (bs []byte) {
	var b bytes.Buffer
	t, err := template.ParseFiles("data/templates/text.tpl")
	if err != nil {
		log.Fatalf("Failed to parse template text.tpl: %v", err)
	}
	t.Execute(&b, data)
	bs = b.Bytes()

	return
}
