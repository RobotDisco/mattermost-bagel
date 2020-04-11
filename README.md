# Bagel (for Mattermost)

Bagel is a Mattermost tool that, given a channel, will randomly match the people in it into pairs and tell them to socialize over a coffee or something.

This is a command-line tool built in Go, which is configured entirely via environment variables, that will perform the above matching when it is run.

## How to create the necessary configuration to run this program
1. Copy the source file `env.sh.template` into a new executable file (we call ours `env.sh`) that looks like the following:
```
export BAGEL_MATTERMOST_URL="https://<your mattermost server url>"

export BAGEL_USERNAME="<some mattermost username>"
export BAGEL_PASSWORD="<that username's password>"

export BAGEL_TEAM_NAME="<the name of some mattermost team name>"
export BAGEL_CHANNEL_NAME="<the name of some mattermost channel under that team>"

export BAGEL_PERSISTENCE_METHOD="<'none' to not track pairing history in a database, 'sqlite' to store pairing history in a SQLite file.>"
export BAGEL_PERSISTENCE_RETRY_COUNT=<number of times to retry generating matches if we detect a conversation pair that was the same as from the last time this program was run>

export BAGEL_SQLITE_FILE="<the name of the file (including path) where you would like to create and load a SQLite database if configured to do so.>"
export BAGEL_SQLITE_STALE_BOUNDARY="<'N days', where N is an integer. This tells us the number of days in the past to check for previous matches between our newly generated pairs in order to not repeat them.>"
```
2. source that file ... in BASH or SH or ZSH this is `source env.sh` for example.

## How to create and run the binary.

1. [Install Go](https://golang.org/dl/)
2. From the repository directory, run `go build`
3. (Optionally) copy the `env.sh` generated above and `mattermost-bagel` files to wherever you would like them to live.
4. After sourcing `env.sh`, run `mattermost-bagel`

## How to run this code as it is being developed.

1. Generate and source the `env.sh` file as done above.
2. Edit the code.
3. Run `go test ./...` to run all unit tests.
4. Run `go run main.go` to run your latest code.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
This software is licensed under the [MIT](https://choosealicense.com/licenses/mit/) software license.
