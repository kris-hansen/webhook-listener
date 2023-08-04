# Webhook Listener

# Webhook Listener

Shh... listen to the webhooks.. 

This project is a simple webhook listener that receives incoming webhook requests and logs the payload to the console in a pretty printed way. The way that I generally use this is to map [ngrok](https://ngrok.com) to map a URL back to my local and then use this webhook listener to inspect non production webhook payloads.

## Getting Started

To get started with this project, clone the repository and install the dependencies:

`go mod tidy`

Then run the project:
`go run main.go`

Or, to build the project:
`go build` which will build the binary `webhook-listener`

It is configured to listen on port `8080` and will print the output to the console. You can also compile this and pipe the output to a file, e.g.,: `./webhook-listener >> webhook_log.txt`

