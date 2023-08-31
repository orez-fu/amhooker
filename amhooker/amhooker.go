package amhooker

import (
	inconstants "amhooker/amhooker/constants"
	"amhooker/amhooker/handlers"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type AMHookerApp interface {
	Start() error
}

type amhookerApp struct {
	port            int
	debugMode       string
	alertConfigpath string
}

func (a *amhookerApp) Start() error {

	log.Println("start 1")
	telegramWebhookHandler := handlers.NewTelegramWebhook(a.alertConfigpath)

	log.Println("start 2")
	router := gin.Default()
	log.Println("start 3")
	router.GET("/ping", telegramWebhookHandler.Ping)
	router.POST("/webhook/telegram", telegramWebhookHandler.SendAlert)

	if err := router.Run(fmt.Sprintf(":%d", a.port)); err != nil {
		log.Println("err")
		return err
	}
	return nil
}

// init create basement
func init() {
	if _, err := os.Stat(inconstants.DEFAULT_TEMPLATE_PATH); err != nil {
		// try to create folder and default template
		if err := os.MkdirAll(inconstants.DEFAULT_TEMPLATE_PATH, os.ModePerm); err != nil {
			log.Fatalf("[ERROR] can not create default templates directory: %v", err)
		}
		defaultTemplatePath := filepath.Join(inconstants.DEFAULT_TEMPLATE_PATH, "default.tmpl")
		if _, err := os.Stat(defaultTemplatePath); err != nil {
			// Missing default template file, try to create it
			f, err := os.Create(defaultTemplatePath)
			if err != nil {
				log.Fatalf("[ERROR] can not create default templates file: %v", err)
			}
			_, err = f.WriteString(defaultTmplContent)
			if err != nil {
				log.Fatalf("[ERROR] can not write content to default templates file: %v", err)
			}
		}
	}
	if _, err := os.Stat(inconstants.EXTRA_TEMPLATE_PATH); err != nil {
		if err := os.MkdirAll(inconstants.EXTRA_TEMPLATE_PATH, os.ModePerm); err != nil {
			log.Fatalf("[ERROR] can not create extra templates directory: %v", err)
		}
	}
}

func NewAMHookerApp(
	port int,
	debugMode string,
	alertConfigPath string,
) AMHookerApp {

	return &amhookerApp{
		port:            port,
		debugMode:       debugMode,
		alertConfigpath: alertConfigPath,
	}
}

var (
	defaultTmplContent string = `
{{ define "telegram.content.vds" }}{{ range . }}
---
🪪 <b>{{ .Labels.alertname }}</b>
{{- if .Annotations.summary }}
📝 {{ .Annotations.summary }}{{ end }}
{{- if .Annotations.description }}
📖 {{ .Annotations.description }}{{ end }}
🏷 Labels:
{{ range $key, $val := .Labels -}}
<i>{{ $key }}</i> = <code>{{ $val }}</code>
{{ end }}
{{- end }}
{{- end }}

{{ if gt (len .Alerts.Firing) 0 }}
🔥 Alerts Firing 🔥
{{ template "telegram.content.vds" .Alerts.Firing }}
{{ end }}
{{ if gt (len .Alerts.Resolved) 0 }}
✅ Alerts Resolved ✅
{{ template "telegram.content.vds" .Alerts.Resolved }}
{{ end }}
	`
)
