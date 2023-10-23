# github-jwt-generator
Generate Bearer tokens to act as a GitHub App against the GitHub API

### Run

```
go mod tidy

export GH_APP_ID="#####"
export PRIV_KEY_PATH="$HOME/Downloads/abc-123.2083-10-06.private-key.pem"

go run main.go
```

### Reference

https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-as-a-github-app
