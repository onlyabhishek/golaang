package main

import (
     "flag"
     "fmt"
 )

 func main() {
     // Define a string variable to hold the flag value
     var name string

     // Bind the flag to the variable
	 flag.StringVar(&name, "name", "World", "a name to say hello to")

   // Parse the command-line flags
    flag.Parse()

    // Use the variable
     fmt.Printf("Hello, %s!\n", name)
}
