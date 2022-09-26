# Gitag


gitag is a command line tool built with go(lang) to help automate git tag release

## Build


```bash
GOOS=linux GOARCH=amd64 go build -o gitag
chmod +x gitag && mv /usr/local/bin/gitag
```

## Usage


```bash
# Move to Git Directory, Make sure have git tagged example # 1.0.1 
# Commit Message Fix: Minor: Major:
gitag
# Example Fix: xxxxx return 1.0.2
```
