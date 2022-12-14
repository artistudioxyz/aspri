## 🔥 Development
Development notes
- Install from source
    - Re-initiate go.mod : `rm go.mod && go mod init aspri`
    - Go Install : `go install`
    - Run : `aspri --help`
- Deployment to registry : `GOPROXY=proxy.golang.org go list -m github.com/artistudioxyz/aspri@{version}`