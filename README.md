<p align="center"><img src="screenshot.png"></p>

<p align="center">
    <img src="https://img.shields.io/github/last-commit/artistudioxyz/aspri" alt="Last Commit">
    <img src="https://img.shields.io/github/languages/code-size/artistudioxyz/aspri" alt="Code Size">
    <img src="https://img.shields.io/github/go-mod/go-version/artistudioxyz/aspri" alt="Go Mod Version">
    <img src="https://img.shields.io/github/v/tag/artistudioxyz/aspri" alt="Latest Tag">
    <img src="https://github.com/artistudioxyz/aspri/actions/workflows/release.yml/badge.svg" alt="Build Status">
    <img src="https://img.shields.io/github/stars/artistudioxyz/aspri?style=social" alt="Stars">
</p>

# ASPRI (Asisten Pribadi)

💃 a Collection of scripts and libraries to speed up sotware development process

## 📝 Installation

You can install aspri by using one of these method

### Install from binary

- Download the latest binary from [release branch](https://github.com/artistudioxyz/aspri/tree/release/dist)
- Extract the binary to your `$PATH` directory
- Make sure the binary is executable

### Install using `go get`

- Run : `go get github.com/artistudioxyz/aspri`

### Install from source

- Re-initiate go.mod : `rm go.mod && go mod init aspri`
- Go Install : `go install`
  - to Generate binary for other OS : `./build.sh`

### Note

- Please add your Go install directory to your system's shell path. [Learn More](https://go.dev/doc/tutorial/compile-install)
- Tutorial in bahasa can be found [here](https://www.youtube.com/watch?v=oe5a-2OVUco)

## 📟 Commands

[Contribution](library/contribution.go) :

- Calculate Contribution : `--contribution --date-start {date} --date-end {date} --path {workdir}`

[ChatGPT](library/chatgpt.go) :

- Start Chat : `--chatgpt --api-key {API_KEY}`
  - The chat support multiple line, don't forget to end it with `~` to get an answer.
  - Get the api key from [here](https://beta.openai.com/account/api-keys)

[Docker](library/docker.go) :

- Stop and Remove Container : `--docker --prune -id {identifier}`
- Compose restart (down & up) : `--docker-compose --restart -f {filename}`

[File](library/file.go) :

- Asset minifier (.js and .css) : `--minify --path {workdir}`
- Count files containing text : `--file --count --path {workdir} --text {text} --exclude {dirname}`
- Directory Stats : `--dir --stats --path {workdir}`
- Extract Urls : `--extract-url --path {workdir} --url {url}`
- Find files older than : `--file --find --older-than --days {days} --regex {regex} --path {workdir} --dry-run`
- Find files younger than : `--file --find --younger-than --days {days} --regex {regex} --path {workdir} --dry-run`
- Find files between dates : `--file --find --between --start {start} --end {end} --regex {regex} --path {workdir} --dry-run`
- Remove Directories or Files Nested by Filenames :
  - Remove Directories `--dir --remove --dirname {dirname} --path {workdir}`
  - Remove Files `--file --remove --f {filename} --path {workdir}`
- Remove Files Nested Except Extensions : `--file --remove --ext {.php} --except {composer.json} --path {workdir}`
- Remove Files older than x days matching regex nested : `--file --remove --older-than --days {days} --regex {regex} --path {workdir} --dry-run`
- Remove Directory older than x days : `--dir --remove --older-than --days {days} --level {0} --path {workdir} --dry-run`
- Search and Replace :
  - in Directory : `--search-replace --path {dir} --from {text} --to {text}`
  - in File : `--search-replace -f {filename} --from {text} --to {text}`
- Standardize directory name : `--dir --standardize --path {workdir}`
- **Support Multiple Params**
  - Dirname : `--dirname {dirname}`
  - Filename : `-f {filename}`
  - Except : `--except {except}`
  - Extension : `--ext {ext}`

[Git](library/git.go) :

- Commit and Push : `--git -m {message}`
- Gone : `--git --gone`
- Reset to previous state (Ignore changes and Remove untracked files) : `--git --reset`
- Reset Cache : `--git --reset-cache`

[Markdown](library/markdown.go) :

- Extract markdown content by heading : `--md --path {workdir} --heading {heading}`
- Extract markdown headings : `--md --path {workdir} --heading {heading}`
- Remove Link from Markdown File : `--md --remove-link --path {workdir}`

[NoIP](library/noip.go) :

- Update Hostname : `--noip --update -u {username} -p {password} --hostname {hostname}`

[PHP](library/php.go) :

- List all class in directory nested : `--php --list-class --path {workdir}`
- List all function in directory nested : `--php --list-function --path {workdir}`
- List function call in directory nested : `--php --list-function-call --path {workdir} --functionname {functionname}`
- **Support Multiple Params**
  - FunctionName : `--functionname {FunctionName}`

[PHPCS](library/phpcs.go) :

- PHPCS Install Ruleset : `--phpcs --install`

[Rsync](library/rsync.go)

- Generate Rsync command based on [rsync.json](docs/rsync.json) : `--rsync`

[Miscellaneous](library/miscellaneous.go) :

- Self Update : `--self-update`

[Quotes](library/quotes.go) :

- Quote of the day : `--quote-of-the-day`

[Syncthing](library/syncthing.go) :

- Remove all conflicts files after certain days : `--syncthing --remove-conflicts --days {days} --dry-run`

[WordPress](wordpress/wordpress.go) :

- Refactor Dot Framework : `--wp-refactor --path {workdir} --from {namespace} --to {namespace} --type {plugin|theme}`
- WP Clean Project Files for Production : `--wp-clean --path {workdir} --type {wordpress|github}`
- WP Plugin Build Check : `--wp-plugin-build-check --path {workdir}`
  - Build WP Plugin : `--wp-plugin-build --path {workdir} --type {wordpress|github}`
  - Release WP Plugin : `--wp-plugin-release --path {workdir} --to {version}`
- WP Theme Build Check : `--wp-theme-build-check --path {workdir}`
  - Build WP Plugin : `--wp-theme-build --path {workdir} --type {wordpress|github}`
- WP Tag Trunk for Subversion (SVN) : `--wp-tag-trunk --path {workdir}`

[YouTube](library/youtube.go)

- Extract YouTube Video Data : `--youtube --extract --path {filepath}`

## ⚒️ Built with

- [Commitlint](https://commitlint.js.org)
- [Golang pflag](https://pkg.go.dev/github.com/spf13/pflag)
- [Husky](https://typicode.github.io/husky)
- [Release-It](https://www.npmjs.com/package/release-it)
  - [Conventional Changelog](https://github.com/release-it/conventional-changelog)

## ⭐️ Support & Contribution

- Help support me by giving a 🌟 or [donate][website]

[website]: https://agung2001.github.io
