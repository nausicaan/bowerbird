# Bowerbird

Bowerbird is a WordPress plugin update install tool. It simplifies the process of pulling free updates from [WordPress Packagist](https://repo.packagist.org) or premium updates from a subscription, then pushing them to Git and adding version tags to a private repository like Satis, if needed.

## Prerequisites

- Login information to download the update package. -- ***premium content only*** --

- Googles' [Go language](https://go.dev) installed to enable building executables from source code.

## Build

From the root folder containing *main.go*, use the command that matches your environment:

### Windows & Mac:

```console
go build -o [name] main.go
```

### Linux:

```console
GOOS=linux GOARCH=amd64 go build -o [name] main.go
```

## Run

Navigate to the folder containing your ***composer.json*** file and run:

```console
[program] [flag] [vendor/plugin]:[version] [ticket#]
```

## Examples

Currently there are three supported scenarios available, using flags to specify each.

### Premium:

```console
bowerbird -p bcgov-plugin/events-virtual:1.13.4 795
```

### WPackagist:

```console
bowerbird -m wpackagist-plugin/mailpoet:4.6.1 821
```

### Release:

```console
bowerbird -r bcgov-plugin/bcgov-inline-comments:1.9.0 820
```

Flags `-r` and `-m` can accept multiple updates, chain together as many as you like!

## License

Code is distributed under [The Unlicense](https://github.com/nausicaan/free/blob/main/LICENSE.md) and is part of the Public Domain.
