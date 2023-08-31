# Development

**Alert manager webhook payload example**

```json
{
  "receiver": "webhook",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "Test",
        "dc": "eu-west-1",
        "instance": "localhost:9090",
        "job": "prometheus24"
      },
      "annotations": {
        "description": "some description"
      },
      "startsAt": "2018-08-03T09:52:26.739266876+02:00",
      "endsAt": "0001-01-01T00:00:00Z",
      "generatorURL": "http://example.com:9090/graph?g0.expr=go_memstats_alloc_bytes+%3E+0\u0026g0.tab=1"                                                                                  
    }
  ],
  "groupLabels": {
    "alertname": "Test",
    "job": "prometheus24"
  },
  "commonLabels": {
    "alertname": "Test",
    "dc": "eu-west-1",
    "instance": "localhost:9090",
    "job": "prometheus24"
  },
  "commonAnnotations": {
    "description": "some description"
  },
  "externalURL": "http://example.com:9093",
  "version": "4",
  "groupKey": "{}:{alertname=\"Test\", job=\"prometheus24\"}"
}
```

### Transform alert format

In `amhooker/handlers/telegram_webhook.go`, 
Before send group of alert, need to transform from origin webhook payload to `AlertBodyTransform` model.

Example of alert body transform object: 

```json

{
  "receiver": "webhook",
  "status": "firing",
  "alerts": {
    "firing": [
      {
        "status": "firing",
        "labels": {
          "alertname": "Test 1",
          "dc": "eu-west-1",
          "instance": "localhost:9090",
          "job": "prometheus24"
        },
        "annotations": {
          "summary": "some mary",
          "description": "some description"
        },
        "startsAt": "2018-08-03T09:52:26.739266876+02:00",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "http://example.com:9090/graph?g0.expr=go_memstats_alloc_bytes+%3E+0\u0026g0.tab=1"
      },
      {
        "status": "firing",
        "labels": {
          "alertname": "Test 2",
          "dc": "eu-west-1",
          "instance": "localhost:9090",
          "job": "prometheus24"
        },
        "annotations": {
          "summary": "some mary",
          "description": "some description"
        },
        "startsAt": "2018-08-03T09:52:26.739266876+02:00",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "http://example.com:9090/graph?g0.expr=go_memstats_alloc_bytes+%3E+0\u0026g0.tab=1"
      }
    ],
    "resolved": [
      {
        "status": "resolved",
        "labels": {
          "alertname": "Test X",
          "dc": "eu-west-1",
          "instance": "localhost:9090",
          "job": "prometheus24"
        },
        "annotations": {
          "summary": "some mary",
          "description": "some description"
        },
        "startsAt": "2018-08-03T09:52:26.739266876+02:00",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "http://example.com:9090/graph?g0.expr=go_memstats_alloc_bytes+%3E+0\u0026g0.tab=1"
      }
    ]
  },
  "groupLabels": {
    "alertname": "Test",
    "job": "prometheus24"
  },
  "commonLabels": {
    "alertname": "Test",
    "dc": "eu-west-1",
    "instance": "localhost:9090",
    "job": "prometheus24"
  },
  "commonAnnotations": {
    "description": "some description"
  },
  "externalURL": "http://example.com:9093",
  "version": "4",
  "groupKey": "{}:{alertname=\"Test\", job=\"prometheus24\"}"
}
```


## References

- [Ultimate config for golang production](https://benchkram.de/blog/dev/ultimate-config-for-golang-apps) + [Code base](https://github.com/benchkram/cli-utils/tree/main/base)
- [Go Telegram Bot](https://github.com/go-telegram/bot)
- [Go YAML in simple](https://marketsplash.com/tutorials/go/golang-yaml/)
- Go Template:
  - [Teamplte Introduction](https://gowebexamples.com/templates/)
  - [Example](https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go)
  - [Package docs](https://pkg.go.dev/text/template#hdr-Text_and_spaces)
