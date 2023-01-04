# Mini Application 💾

### Creating applications from the command line

🔹 In the mini application project, I used the Cli package to build a command line application in Golang. <br>

🔹 The first part of the application aims to locate the IP and the second a server of a given domain.

🔹 In the terminal the application is executed as follows:

```go
      go run main.go ip --host domain.com
      go run main.go servers --host domain.com
```

 🔹 In case of invalid domain:
```go
     flags := []cli.Flag{
		cli.StringFlag{
		Name:  "host",
		Value: "google.com", // If the domain is invalid, the value element will be used by default.
		},
	  }
``` 
