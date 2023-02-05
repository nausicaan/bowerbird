# WordPress Update Install

Simplify the process of pulling free updates from [WordPress Packagist](https://repo.packagist.org) or premium updates from a subscription, then pushing them to Git and adding version tags to a private repository like Satis, if needed.

## Prerequisites

- Login information to download the update package. -- ***premium content only*** --

- Googles' [Go language](https://go.dev) installed to enable building executables from source code.

## Build

From the root folder containing *main.go*, use the command that matches your environment:

### Windows & Mac:

```console
go build -o <build_location>/<program_name> main.go
```

### Linux:

```console
GOOS=linux GOARCH=amd64 go build -o <build_location>/<program_name> main.go
```

## Run

Navigate to the folder containing your *composer.json* file and run:

```console
<build_location>/<program_name> <flag> <full_plugin_name>:<update_version> <jira_ticket_number>
```

## Examples

Currently there are three supported scenarios available, using flags to specify each.

### Premium:

```console
~/Documents/programs/upinstall -p bcgov-plugin/gravityforms:2.6.8.4 759
```

### WPackagist:

```console
~/Documents/programs/upinstall -w wpackagist-plugin/mailpoet:5.5.2 762
```

### Release:

```console
~/Documents/programs/upinstall -r bcgov-plugin/events-virtual:1.13.4 795
```

Flags `-r` and `-w` can accept multiple updates, chain together as many as you like!

## License

Code is distributed under [The Unlicense](https://github.com/nausicaan/free/blob/main/LICENSE.md) and is part of the Public Domain.
