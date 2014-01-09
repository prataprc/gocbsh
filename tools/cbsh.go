package main

import (
    "flag"
    "fmt"
    "github.com/couchbaselabs/clog"
    "github.com/couchbaselabs/go-couchbase"
    "github.com/prataprc/liner"
    "github.com/prataprc/cbsh"
    "io"
    "net/url"
    "os"
    "os/signal"
    "syscall"
)

var sh struct {
    prompt string
    w      io.Writer
    u      *url.URL
    client couchbase.Client
    pool   couchbase.Pool
    bucket *couchbase.Bucket
}

func argParse() {
    var err error

    poolname := flag.String("pool", "default",
        "pool to connect to (defaults to `default`)")
    bucketname := flag.String("bucket", "default",
        "bucket to connect to (defaults to `default`)")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr,
            "%v [flags] http://user:pass@host:8091/\n\nFlags:\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(64)
    }
    flag.Parse()
    if flag.NArg() < 1 {
        flag.Usage()
    }
    if sh.u, err = url.Parse(flag.Arg(0)); err != nil {
        clog.Fatal(err)
    } else {
        if sh.client, err = couchbase.Connect(sh.u.String()); err != nil {
            clog.Fatal(err)
        }
        if err = pool(*poolname); err != nil {
            clog.Fatal(err)
        }
        if err = bucket(*bucketname); err != nil {
            clog.Fatal(err)
        }
    }
    sh.prompt = sh.u.Host + "/" + *poolname + "/" + *bucketname
}

func main() {
    argParse()
    handleInteractiveMode()
}

func handleInteractiveMode() {
    var err error
    var line string

    homedir := homeDir()
    liner := liner.NewLiner()
    defer liner.Close()

    sh.w = os.Stdout
    LoadHistory(liner, homedir)
    go signalCatcher(liner)

    for {
        if line, err = liner.Prompt(sh.prompt + "> "); err != nil {
            clog.Error(err)
            break
        } else if line == "" {
            continue
        } else {
            UpdateHistory(liner, homedir, line)
            if err = interpret(line); err != nil {
                clog.Error(err)
            }
        }
    }
}

/**
 *  Attempt to clean up after ctrl-C otherwise
 *  terminal is left in bad shape
 */
func signalCatcher(liner *liner.State) {
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT)
    <-ch
    liner.Close()
    os.Exit(0)
}

/*
 * Get user's home directory
 */
func homeDir() string {
    hdir := os.Getenv("HOME") // try to find a HOME environment variable
    if hdir == "" {           // then try USERPROFILE for Windows
        hdir = os.Getenv("USERPROFILE")
        if hdir == "" {
            fmt.Printf("Unable to determine home directory, history file disabled\n")
        }
    }
    return hdir
}
