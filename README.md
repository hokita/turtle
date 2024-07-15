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
endpoints:
    sample1:
        method: "GET"
        url: "http://example.com/sample1"
    sample2:
        method: "POST"
        url: "http://example.com/sample2"
        body: |
            {
              "name": "Karashi",
              "type": "Degu"
            }
```

### Call turtle

Load the `turtle.yaml` file from the directory where the turtle command is executed.

```shell
$ turtle sample1
{"name":...} # Return the API response
```
