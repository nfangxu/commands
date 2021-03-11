package starter

import (
    "context"
    "fmt"
    "github.com/spf13/cobra"
    "os"
    "time"
)

const starter = "https://github.com/DoNewsCode/core-starter.git"

// Cmd represents the new command.
var Cmd = &cobra.Command{
    Use:   "create [repository: https://github.com/DoNewsCode/core-starter.git] [name]",
    Short: "Create a project",
    Long:  "Create a project using the repository template. Example: fangx create https://github.com/DoNewsCode/core-starter.git helloworld",
    Run:   run,
}

func run(cmd *cobra.Command, args []string) {
    var name,resp string
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()
    if len(args) == 0 {
        fmt.Fprintf(os.Stderr, "\033[31mERROR: project name is required.\033[m Example: fangx create helloworld\n")
        return
    } else if len(args) == 1 {
        name = args[0]
        resp = starter
    } else {
        name = args[1]
        resp = args[0]
    }

    p := &Starter{Name: name, Resp: resp}
    if err := p.New(ctx, wd); err != nil {
        fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
        return
    }
}
