func pool(argstr string) (err error) {
    var s string
    parts := splitArg(argstr, " ", 2)
    if len(parts) < 1 {
        s, err = prettyPrint(sh.pool, "")
        fmt.Fprintf(sh.w, "%v\n", s)
    } else if sh.u == nil {
        err = fmt.Errorf("Not connected to any server")
    } else if sh.pool, err = sh.client.GetPool(parts[0]); err == nil {
        sh.prompt = sh.prompt + "/" + parts[0]
        if len(parts) > 1 {
            return bucket(parts[1])
        }
    }
    return
}


