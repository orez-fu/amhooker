package handlers

import (
	cfgtelegram "amhooker/amhooker/configs/telegram"
	"amhooker/amhooker/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	inconstants "amhooker/amhooker/constants"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	"github.com/sanity-io/litter"
)

type TelegramWebhook interface {
	SendAlert(c *gin.Context)
	Ping(c *gin.Context)
}

type telegramWebhook struct {
	telegramBot *bot.Bot
}

func (tw *telegramWebhook) SendAlert(c *gin.Context) {

	var params AlertParams
	var alerts models.AlertBody

	if c.ShouldBindQuery(&params) != nil {
		c.JSON(400, gin.H{
			"code":    "TW4002",
			"message": "invalid query params",
		})
		return
	}
	log.Printf("Query params: ChatID=%s , TopicID=%d, Template=%s", params.ChatID, params.TopicID, params.Template)

	if c.BindJSON(&alerts) != nil {
		c.JSON(400, "")
		c.JSON(400, gin.H{
			"code":    "TW4004",
			"message": "parse body content from alert manager failed, need to update source to match compatibility",
		})
		return
	}
	log.Printf("Request body: %s", alerts)

	messageText, err := RenderMessageText(params.Template, &alerts)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "TW5002",
			"message": "render message text failed",
		})
		return
	}
	message := bot.SendMessageParams{
		ChatID:    params.ChatID,
		Text:      messageText,
		ParseMode: "html",
	}
	if params.TopicID > 0 {
		message.MessageThreadID = params.TopicID
	}

	if _, err := tw.telegramBot.SendMessage(context.Background(), &message); err != nil {
		log.Printf("Error sending alert message: %v", err)
		c.JSON(200, gin.H{
			"code":    "TW5004",
			"message": "error in sending alert message",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "00",
		"message": "send alert message success",
	})
}

func (tw *telegramWebhook) Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}

func NewTelegramWebhook(alertConfigPath string) TelegramWebhook {
	telegramBot := cfgtelegram.GetTelegramBot(alertConfigPath)
	return &telegramWebhook{
		telegramBot: telegramBot,
	}
}

type AlertParams struct {
	TopicID  int    `form:"topic_id"`
	ChatID   string `form:"chat_id"`
	Template string `form:"template"`
}

func RenderMessageText(templateName string, alerts *models.AlertBody) (string, error) {
	if templateName == "" {
		log.Printf("No template is specified. Set using default template")
		templateName = "default"
	}
	templatePath := filepath.Join(inconstants.EXTRA_TEMPLATE_PATH, fmt.Sprintf("%s.tmpl", templateName))
	templatePathOk := false
	if _, err := os.Stat(templatePath); err != nil {
		log.Printf("Cannot find or read template from extra templates directory")
	} else {
		templatePathOk = true
	}

	pwd, _ := os.Getwd()
	if !templatePathOk {
		log.Println("Try to find template in common templates directory")
		templatePath = filepath.Join(pwd, inconstants.DEFAULT_TEMPLATE_PATH, fmt.Sprintf("%s.tmpl", templateName))
		if _, err := os.Stat(templatePath); err != nil {
			log.Printf("Cannot find or read template from common templates directory")
		} else {
			templatePathOk = true
		}
	}
	if !templatePathOk {
		log.Printf("No template found for '%s'. Creat a raw message...", templateName)
		result, err := json.Marshal(alerts)
		if err != nil {
			log.Printf("[ERROR] convert raw message to json failed: %s", err.Error())
			return "", err
		}
		return string(result), nil
	} else {
		log.Printf("Template '%s' found at '%s'. Create a new message...", templateName, templatePath)
		alertTransform := TransformAlertFormat(*alerts)
		litter.Dump(alertTransform)
		tmpl := template.Must(template.ParseFiles(templatePath))
		var out bytes.Buffer
		err := tmpl.Execute(&out, alertTransform)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return "", err
		}
		log.Println(out.String())
		return out.String(), nil
	}
}

func TransformAlertFormat(alerts models.AlertBody) *models.AlertBodyTransform {
	firingAlerts := make([]*models.Alert, 0)
	resolvedAlerts := make([]*models.Alert, 0)
	for _, alert := range alerts.Alerts {
		if alert.Status == "firing" {
			firingAlerts = append(firingAlerts, &alert)
		} else if alert.Status == "resolved" {
			resolvedAlerts = append(resolvedAlerts, &alert)
		}
	}
	alertStandard := models.AlertBodyTransform{
		Alerts: models.AlertInTransform{
			Firing:   firingAlerts,
			Resolved: resolvedAlerts,
		},
		CommonLabels:      alerts.CommonLabels,
		CommonAnnotations: alerts.CommonAnnotations,
		ExternalURL:       alerts.ExternalURL,
		GroupKey:          alerts.GroupKey,
		GroupLabels:       alerts.GroupLabels,
		Receiver:          alerts.Receiver,
		Status:            alerts.Status,
		Version:           alerts.Version,
	}
	return &alertStandard
}
