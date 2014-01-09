func pp(argstr string) (err error) {
    var s string

    parts := splitArg(argstr, " ", 2)
    switch parts[0] {
    case "bucket":
        parts := splitArg(parts[1], ".", 2)
        if len(parts) > 1 {
            s, err = prettyPrint(sh.pool.BucketMap[parts[0]], parts[1])
        } else {
            s, err = prettyPrint(sh.pool.BucketMap[parts[0]], "")
        }
        fmt.Fprintf(sh.w, "%v\n", s)
    case "pools":
        s, err = prettyPrint(sh.client.Info, "")
        fmt.Fprintf(sh.w, "%v\n", s)
    }
    return
}


