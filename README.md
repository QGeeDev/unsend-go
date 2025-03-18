# unsend-go

Go SDK for [unsend.dev](https://unsend.dev)

## Supported versions
- Unsend 1.4.x
- Go 1.22.x

## About this project
This was built to be an SDK that can be used with both the cloud hosted and self hosted versions of Unsend

## How to use

- Run `go get github.com/QGeeDev/unsend-go` to install module to your project

### Cloud Hosted

- Generate an API key for Unsend by following [this guide](https://app.unsend.dev/dev-settings/api-keys)
- Set environment variable `UNSEND_API_KEY` to the value of the key generated
- Create an Unsend client in your Go project as shown in [examples](/examples/)

### Self-hosted
- Generate an API key for Unsend by following [this guide](https://app.unsend.dev/dev-settings/api-keys)
- Set environment variable `UNSEND_API_KEY` to the value of the key generated
- Set environment variable `UNSEND_BASE_URL` to FQDN of Unsend instance INCLUDING the `/api` at the end
  - example: `https://unsend.test.com/api`
- Create an Unsend client in your Go project as shown in [examples](/examples/)

## Environment variables
| Variable Name     | Required | Default                      |
|-------------------|----------|------------------------------|
| `UNSEND_API_KEY`  | `YES`    | N/A                          |
| `UNSEND_BASE_URL` | `NO`     | `https://app.unsend.dev/api` |
