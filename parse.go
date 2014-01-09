package cbsh

import (
    "encoding/json"
    "fmt"
    "github.com/couchbaselabs/go-couchbase"
    "net/url"
    "reflect"
    "strings"
    "text/scanner"
)

type CommandHandler interface {
    Parse()
}

var commands = map[string]CommandHandler{}
    "connect": connect,
    "pp":      pp,
    "list":    list,
    "pool":    pool,
    "bucket":  bucket,
    "vbmap":   vbmap,
    "nodes":   nodes,
    "get":     get,
}

func interpret(line string) (err error) {
    parts := splitArg(line, " ", 2)
    argstr := ""
    if len(parts) > 1 {
        argstr = parts[1]
    }
    if err = validCommand(parts[0]); err != nil {
        return err
    }
    return commands[parts[0]](argstr)
}

func parse(line string) ([]string, error) {
    var s scanner.Scanner

    parts := make([]string, 0, 10)
    s.Init(strings.NewReader(line))
    for s.Scan() != scanner.EOF { // do something with tok
        parts = append(parts, s.TokenText())
    }
    return parts, nil
}

func validCommand(cmdname string) (err error) {
    for k, _ := range commands {
        if cmdname == k {
            return nil
        }
    }
    return fmt.Errorf("Invalid command %q", cmdname)
}

func prettyPrint(obj interface{}, attr string) (s string, err error) {
    var v, bs []byte
    var mobj map[string]interface{}
    var sobj []interface{}

    if v, err = json.Marshal(obj); err == nil {
        switch reflect.TypeOf(obj).Kind() {
        case reflect.Slice:
            json.Unmarshal(v, &sobj)
            obj = sobj
        case reflect.Map, reflect.Struct:
            json.Unmarshal(v, &mobj)
            obj = mobj
            if attr != "" {
                obj = mobj[attr]
            }
        default:
            err = fmt.Errorf("Neither slice nor map")
        }
        bs, err = json.MarshalIndent(obj, "", "  ")
        s = string(bs)
    }
    return
}

func splitArg(argstr string, sep string, count int) []string {
    parts := strings.SplitN(strings.Trim(argstr, " "), sep, count)
    args := make([]string, 0)
    for _, s := range parts {
        s1 := strings.Trim(s, " ")
        if s1 != "" {
            args = append(args, s1)
        }
    }
    return args
}
