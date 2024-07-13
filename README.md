# turtle
Try a slow life like turtle.

## Installation
```shell
$ brew install hokita/tap/turtle
```

## Usage
### Create config file
Create a file named `turtle.yaml` in the root directory of the project or another appropriate location.

### Edit turtle.yaml
Here is sample.
```yaml
# https://pokeapi.co/
endpoints:
  ditto:
    url: "https://pokeapi.co/api/v2/pokemon/ditto"
  tyranitar:
    url: "https://pokeapi.co/api/v2/pokemon/tyranitar"
```

### Call turtle
Load the `turtle.yaml` file from the directory where the turtle command is executed.
```shell
$ turtle ditto
{"abilities":...} # Return the API response
```
