func vbmap(argstr string) (err error) {
    if sh.bucket == nil {
        err = fmt.Errorf("Not connected to a bucket")
    } else {
        var s string
        s, err = prettyPrint(sh.bucket.VBServerMap(), "")
        fmt.Fprintf(sh.w, "%v\n", s)
    }
    return
}


