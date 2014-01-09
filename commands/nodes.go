func nodes(argstr string) (err error) {
    if sh.bucket == nil {
        err = fmt.Errorf("Not connected to a bucket")
    } else {
        var s string
        sh.bucket.Nodes()
        s, err = prettyPrint(sh.bucket.Nodes(), "")
        fmt.Fprintf(sh.w, "%v\n", s)
    }
    return
}



