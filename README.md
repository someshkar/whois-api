# Whois API

A simple serverless REST Whois API, written in Go. Mainly exists because other Whois API services are either too expensive or have a strict rate limit.

## Usage

This API is deployed at https://whois-api.now.sh. However, self hosting is as simple as forking this repository and setting up a [Vercel](https://vercel.com) account.

### `POST /`

Use this format to `POST` data to the main endpoint:

```json
{
  "domain": "example.com"
}
```

### `GET /ping`

Use this endpoint to check if the service is running. It should respond with `pong`.

## Roadmap

- [ ] Write MultiHandler function which handles multiple domains concurrently using goroutines.