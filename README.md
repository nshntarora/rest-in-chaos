```
     ____  _________________   ____         ________
    / __ \/ ____/ ___/_  __/  /  _/___     / ____/ /_  ____ _____  _____
   / /_/ / __/  \__ \ / /     / // __ \   / /   / __ \/ __ / __ \/ ___/
  / _, _/ /___ ___/ // /    _/ // / / /  / /___/ / / / /_/ / /_/ (__  )
 /_/ |_/_____//____//_/    /___/_/ /_/   \____/_/ /_/\__,_/\____/____/

```

### Add unreliability to any HTTP service

## Why?

I mostly do frontend at work, and I've always ignored building error states. Error states are a second class citizen in building UIs. One of the reasons is the developers building it don't get error states often from their APIs. What if we change that?

## How?

rest-in-chaos is a command line utility that spawns proxy server, a really shitty one at that. Every request made to the rest-in-chaos server is either proxied to the passed URL, or returns an error response. It's completely random. It's chaos.

## Installation

Run the following command

```
go get -u github.com/nshntarora/rest-in-chaos
```

The binary will now be available in the `bin` directory in your `GOPATH`

## Usage

```
rest-in-chaos <<YOUR URL>>
```

The server will be running on the port **24267**

## Example

```
rest-in-chaos http://localhost:3000
```

Any request to localhost:24267 will be proxied to localhost:3000, with the exception of some requests chosen at random.
