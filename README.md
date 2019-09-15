# Bagel (for Mattermost)

Bagel is a Mattermost tool that, given a channel, will randomly match the people in it into pairs and tell them to socialize over a coffee or something.
## Installation

1. Install Go
2. Create a directory for your go code (if necessary)
3. Set your `GOPATH` environment variable to point to your Go directory
4. Run `go get github.com/RobotDisco/mattermost-bagel`

## Usage

(This is very dev-oriented at the moment, sorry)

1. Go to your `$GOPATH/src/github.com/RobotDisco/mattermost-bagel` directory.
2. Create a file (we call ours `env.sh` that looks like the following:
```
export BAGEL_MATTERMOST_URL="https://<your mattermost server url>"

export BAGEL_USERNAME="<some mattermost username>"
export BAGEL_PASSWORD="<that username's password>"

export BAGEL_TEAM_NAME="<the name of some mattermost team name>"
export BAGEL_CHANNEL_NAME="<the name of some mattermost channel under that team>"
```
3. source that file ... in BASH or SH or ZSH this is `source env.sh` for example.
4. `go run main.go`

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
This software is licensed under the [MIT](https://choosealicense.com/licenses/mit/) software license.
