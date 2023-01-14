# WordPress Updater

Simplify the process of pulling free updates from [WordPress Packagist](https://wpackagist.org) or premium updates from a subscription, pushing them to Git and adding version tags to Satis, if applicable.

## Prerequisites

- Login information to download the update package. -- *premium content* --

- Googles' [Go language](https://go.dev) installed to enable building executables from source code.

## Build

From the root folder containing *main.go*, use the command that matches your environment:

### Mac:

```bash
go build -o <build_location>/<program_name> main.go
```

### Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o <build_location>/<program_name> main.go
```

## Run

Navigate to the folder containing your *composer.json* file and run:

```bash
<build_location>/<program_name> <full_plugin_name>:<update_version> <jira_ticket_number>
```

## Examples

### Premium:

```bash
~/Documents/programs/updater bcgov-plugin/gravityforms:2.6.8.4 759
```

### Free:

```bash
~/Documents/programs/updater wpackagist-plugin/spotlight-social-photo-feeds:1.4.2 762
```

Multiple updates, chain together as many as you like!

```bash
~/Documents/programs/updater wpackagist-plugin/redis-cache:2.2.3 761 wpackagist-plugin/spotlight-social-photo-feeds:1.4.2 762
```

## License

Code is distributed under [The Unlicense](https://github.com/nausicaan/free/blob/main/LICENSE.md) and is part of the Public Domain.
