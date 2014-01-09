package commands

import (
    "github.com/prataprc/cbsh"
)

type CmdConnect struct {
    cbsh.Command
}

func connect(argstr string) (err error) {
    parts := splitArg(argstr, " ", 2)
    if len(parts) < 1 {
        err = fmt.Errorf("Need argument to connect")
    } else if sh.u, err = url.Parse(parts[0]); err != nil {
        return
    } else if sh.client, err = couchbase.Connect(sh.u.String()); err == nil {
        sh.prompt = sh.u.Host
        if len(parts) > 1 {
            return pool(parts[1])
        }
    }
    return
}

func init() {
    "connect": CmdConnect
}
