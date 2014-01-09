func get(argstr string) (err error) {
    var ob interface{}

    parts := splitArg(argstr, " ", 2)
    if sh.bucket == nil {
        err = fmt.Errorf("Not connected to bucket")
    } else if len(parts) < 1 {
        err = fmt.Errorf("Need argument to get")
    } else if err = sh.bucket.Get(parts[0], &ob); err == nil {
        var s string
        s, err = prettyPrint(ob, "")
        fmt.Fprintf(sh.w, "%v\n", s)
    }
    return
}


