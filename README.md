```
     ____  _________________   ____         ________
    / __ \/ ____/ ___/_  __/  /  _/___     / ____/ /_  ____ _____  _____
   / /_/ / __/  \__ \ / /     / // __ \   / /   / __ \/ __ / __ \/ ___/
  / _, _/ /___ ___/ // /    _/ // / / /  / /___/ / / / /_/ / /_/ (__  )
 /_/ |_/_____//____//_/    /___/_/ /_/   \____/_/ /_/\__,_/\____/____/

```

### Add unreliability to any HTTP service

## Installation

Run the below command

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
