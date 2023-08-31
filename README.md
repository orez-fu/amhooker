# Amhooker

Wehook application for Alert Manager.

Features: 
- Support send message for Telegram chat or thread/topic

## Expose APIs


`/webhook/telegram?topic_id=1?template_name=`

## Configuration

Firstly, create amhooker `config.yaml` with format:

| Param             | Type    | Description                                                              | Example                              |
|-------------------|---------|--------------------------------------------------------------------------|--------------------------------------|
| timeZone          | string  | Set time zone                                                            | "Asia/Singapore"                     |
| timeOutputFormat  | string  | Set time output format, provide a datetime as desired datetime           | "01/01/2004 11:11:11"                |
| splitMessageBytes | number  | Message size, if exceed this size, split origin message to small message | 4000                                 |
| telegram.enabled  | boolean | Set true to enable sending alert to telegram                             | true                                 |
| telegram.botToken | string  | Telegram bot token                                                       | "1234567890:WERDFAjdksf832SD324DAER" |

Example `amhooker.yaml`: 
```yaml
timeZone: "Asia/Singapore"
timeOutputFormat: "31/01/2004 12:12:12"
splitMessageBytes: 4000
telegram: 
  enabled: true 
  botToken: "1234567890:WERDFAjdksf832SD324DAER"
```

## Usage

### Binary

Download the binary file and run command amhooker and pass following arguments or setup environment:

```bash
amhooker [flags]

Flags:
      --config_file string   AMHooker manager config file (*require) (env AMHOOKER_CONFIG_FILE)
      --debug string         Type of debug mode: INFO | DEBUG | NONE . (env AMHOOKER_DEBUG_MODE) (default "INFO")
  -h, --help                 help for amhooker
      --port int             Running application port. (env AMHOOKER_PORT) (default 8866)
```

Example for setup command arguments approach with amhooker config file existed at "/etc/amhooker/amhooker.yaml":

```bash
# Add execute permission
chmod mod +x amhooker

# Start amhooker
amhooker --alert_config_path=/etc/amhooker/amhooker.yaml --debug=DEBUG --port=8181
```

Example for setup environment approach:

```bash
# Export environment variables
export AMHOOKER_CONFIG_FILE=/etc/amhooker/amhooker.yaml
export AMHOOKER_PORT=8181
export AMHOOKER_DEBUG_MODE=DEBUG

# Add execute permission
chmod mod +x amhooker

# Start amhooker
amhooker
```

### Docker


```bash
docker run --rm -v <amhooker_config_file>:<container_amhooker_config_file> -e AMHOOKER_CONFIG_FILE=<container_amhooker_config_file> orezfu/amhooker:<tag>
# docker run --rm -v /etc/amhooker/amhooker.yaml:/amhooker.yaml -e AMHOOKER_CONFIG_FILE=/amhooker.yaml orezfu/amhooker:v0.1.0
```

### Kubernetes Helm Chart

(Continue)

## Some issues:
- `SortedPairs` is not working