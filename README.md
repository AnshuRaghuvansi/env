# env

env is a tiny package to handle environment variables in go application.

## Getting Started

Create a file with the name of environment(Ex. development.env, testing.env etc) in root directory.

Put the variables in the file like this

```
PORT=8080
PASSWORD=YSG&*^HDHS
```
and then import the package in main package and configure the env

```
import (
	"env/env"
)

func main() {

	env.Configure()
  
}

```

Export the enviroment before runing the application 
For example

```
export environment=development
```

then run the go application using 

```
go run main.go
```

now variable are available as env variable and can be access like this 

```
import (
	"env/env"
	"fmt"
	"os"
)

func main() {

	env.Configure()

	port := os.Getenv("PORT")
	pwd := os.Getenv("PASSWORD")

	fmt.Println(port)
	fmt.Println(pwd)
	
	To get specific type of value 
	
	prod, ok := GetBoolValue("IS_PRODUCTION")
	if ok {
		fmt.Println(prod)
	}
}

```
