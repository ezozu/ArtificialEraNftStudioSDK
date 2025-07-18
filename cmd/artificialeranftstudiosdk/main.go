// cmd/artificialeranftstudiosdk/main.go
package main

import (
"flag"
"log"
"os"

"artificialeranftstudiosdk/internal/artificialeranftstudiosdk"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeranftstudiosdk.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
