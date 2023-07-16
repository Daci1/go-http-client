# Go HTTP Client

The Go HTTP Client is a command-line tool developed in Go that allows you to make HTTP requests to a specified URL. It utilizes the `flags` package to parse command-line arguments and leverages `net/http` package to execute the HTTP calls.

## Flags

- **URL (-u):** Specifies the target URL to send the HTTP request.
- **Verb (-v):** Specifies the HTTP method verb to be used (e.g., GET, POST, PUT, DELETE). GET is used by default.
- **Body (-b):** Provides the request body content for POST or PUT requests. Empty body is used by default.
- **Headers (-h):** Allows you to include additional HTTP headers in the request.
- **Help (-help):** Displays a list of all available command-line flags.
- **Pretty (-pretty):** Formats the JSON response in a pretty printed format.

## Installation

To use the Go HTTP Client, ensure you have Go installed on your system and run the following command:

```bash
go get github.com/Daci1/go-http-client
```

## Usage
After installing the Go HTTP Client, you can run it from the command line using the following command:

```bash
go-http-client -u <url> [flags]
```

Here are some examples of how to use the command-line flags:

* Make a GET request:
```bash
go-http-client -u https://api.example.com/resource
```
* Make a POST request with a request body:
```bash
go-http-client -u https://api.example.com/resource -v POST -b '{\"name\":\"John\",\"age\":30}'
```

## Future Work
- Option to Use JSON Files for Body
- Option to Use JSON Files for Headers
- Export to Different File Formats

## Conclusion
In conclusion, the development of the Go HTTP Client was undertaken with the primary goal of gaining familiarity with the Go programming language. By embarking on this project, I have successfully delved into various aspects of Go, such as utilizing the net/http package for making HTTP requests and leveraging the flags package for parsing command-line arguments.

To further enhance my understanding of Go and embrace good programming practices, I would appreciate any feedback and recommendations from other programmers.