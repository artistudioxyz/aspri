<p align="center"><img src="screenshot.png"></p>

<p align="center">
    <img src="https://img.shields.io/github/last-commit/artistudioxyz/aspri" alt="Last Commit">
    <img src="https://img.shields.io/github/languages/code-size/artistudioxyz/aspri" alt="Code Size">
    <img src="https://img.shields.io/github/go-mod/go-version/artistudioxyz/aspri" alt="Go Mod Version">
    <img src="https://img.shields.io/github/v/tag/artistudioxyz/aspri" alt="Latest Tag">
    <img src="https://github.com/artistudioxyz/aspri/actions/workflows/workflow.yml/badge.svg" alt="Build Status">
    <img src="https://img.shields.io/github/stars/artistudioxyz/aspri?style=social" alt="Stars">
</p>

# ASPRI (Asisten Pribadi)

a Collection of scripts and libraries to speed up sotware development process

## 📝 Installation
- Run : `go get github.com/artistudioxyz/aspri`
- Run : `aspri --help`

Note :
- Please add your go install to system PATH [Learn More](https://go.dev/doc/tutorial/compile-install)

## 📟 Commands
[Docker](library/docker.go) :
- Stop and Remove Container : `--docker -prune -id {identifier}` 
- Compose restart (down & up) : `--docker-compose-restart`

[Git](library/git.go) :
- Commit and Push : `--git -m {message}`

[Quotes](library/quotes.go) :
- Quote of the day : `--quote-of-the-day`

[Miscellaneous](library/miscellaneous.go) :
- Search and Replace in Directory or File : `--search-replace --path {dir or file} -from {text} -to {text}`

[WordPress](wordpress/wordpress.go) :
- Refactor : `--wp-refactor --path {workdir} -from {namespace} -to {namespace}`
- WP Plugin Build Check : `--wp-plugin-build-check --path {workdir}`
  - Build WP Plugin (Require Path) : `--wp-plugin-build --path {workdir} --type {wordpress|github}`
- WP Theme Build Check : `--wp-theme-build-check --path {workdir}`
  - Build WP Plugin (Require Path) : `--wp-theme-build --path {workdir} --type {wordpress|github}`

## 🔥 Development
Development notes
- Install from source
  - Re-initiate go.mod : `rm go.mod && go mod init aspri`
  - Go Install : `go install`
  - Run : `aspri --help`
- Deployment to registry : `GOPROXY=proxy.golang.org go list -m github.com/artistudioxyz/aspri@v1.2.0`