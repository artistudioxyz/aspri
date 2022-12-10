<p align="center"><img src="logo.png"></p>

<p align="center">
    <img src="https://img.shields.io/github/last-commit/artistudioxyz/aspri" alt="Last Commit">
    <img src="https://img.shields.io/github/languages/code-size/artistudioxyz/aspri" alt="Code Size">
    <img src="https://img.shields.io/github/go-mod/go-version/artistudioxyz/aspri" alt="Go Mod Version">
    <img src="https://img.shields.io/github/v/tag/artistudioxyz/aspri" alt="Latest Tag">
    <img src="https://img.shields.io/github/stars/artistudioxyz/aspri?style=social" alt="Stars">
</p>

# ASPRI (Asisten Pribadi)

a Collection of scripts and libraries to speed up sotware development process

<p align="center"><img src="screenshot.png"></p>

## 📝 Installation
- Run : `go install`
- Run : `aspri --version`

Note :
- Please add your go install to system PATH [Learn More](https://go.dev/doc/tutorial/compile-install)

## 📟 Commands
[Docker](library/docker.go) :
- Stop and Remove Container : `--docker -prune -id {identifier}` 
- Compose restart (down & up) : `--docker-compose-restart`

[Git](library/git.go) :
- Commit and Push : `--git -push -m {message}`

[Quotes](library/quotes.go) :
- Quote of the day : `--quote-of-the-day`

[Miscellaneous](library/miscellaneous.go) :
- Search and Replace in Directory : `--search-replace-directory --path {workdir} -from {text} -to {text}`

[WordPress](library/wordpress.go) :
- Build WP Plugin : `--wp-plugin-build --path {workdir} -from {namespace} -to {namespace}`