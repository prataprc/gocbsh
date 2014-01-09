func bucket(argstr string) (err error) {
    parts := splitArg(argstr, " ", 2)
    if len(parts) < 1 {
        err = fmt.Errorf("Need argument to pool")
    } else if sh.u == nil {
        err = fmt.Errorf("Not connected to any server")
    } else if sh.bucket, err = sh.pool.GetBucket(parts[0]); err == nil {
        sh.prompt = sh.prompt + "/" + parts[0]
    }
    return
}


