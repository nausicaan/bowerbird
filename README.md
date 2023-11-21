# Bowerbird

Bowerbird is a WordPress plugin update install tool. It simplifies the process of pulling free updates from [WordPress Packagist](https://repo.packagist.org) or premium updates from a subscription, then pushing them to Git and adding version tags to a private repository like Satis, if needed. Named after a very industrious creature who excels at building.

![Bird](bowerbird.webp)

## Prerequisites

- Login information to download the update package. -- ***premium content only*** --

- Googles' [Go language](https://go.dev) installed to enable building executables from source code.

- An *interactions* folder in `~/Documents/` shared between [Silkworm](https://github.com/nausicaan/bowerbird.git) and Bowerbird.

## Build

From the root folder containing the `go` files, use the command that matches your environment:

### Windows & Mac:

``` console
go build -o [name] .
```

### Linux:

``` console
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

Navigate to the folder containing your ***composer.json*** file and run:

``` console
[program] [flag]
```

## Examples

Currently there are three supported scenarios available, using flags to specify each.

### Premium (Purchased third party plugins):

``` console
bowerbird -p
```

### WordPress Packagist (Plugins currently available via WPackagist or Satis):

``` console
bowerbird -w
```

### Release (In-house Production ready content):

``` console
bowerbird -r
```

Flags `-r` and `-w` can accept multiple updates, chain together as many as you like!

## License

Code is distributed under [The Unlicense](https://github.com/nausicaan/free/blob/main/LICENSE.md) and is part of the Public Domain.
