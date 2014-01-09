func list(argstr string) (err error) {
    var s string

    parts := splitArg(argstr, " ", 2)
    switch parts[0] {
    case "nodes":
        nodes := make([]string, 0)
        for _, node := range sh.bucket.Nodes() {
            nodes = append(nodes, node.Hostname)
        }
        s, err = prettyPrint(nodes, "")
        fmt.Fprintf(sh.w, "%v\n", s)
    case "pools":
        pools := make([]string, 0)
        for _, restPool := range sh.client.Info.Pools {
            pools = append(pools, restPool.Name)
        }
        s, err = prettyPrint(pools, "")
        fmt.Fprintf(sh.w, "%v\n", s)
    case "buckets":
        buckets := make([]string, 0)
        for k, _ := range sh.pool.BucketMap {
            buckets = append(buckets, k)
        }
        s, err = prettyPrint(buckets, "")
        fmt.Fprintf(sh.w, "%v\n", s)
        fmt.Fprintf(sh.w, "BucketURL\n")
        s, err = prettyPrint(sh.pool.BucketURL, "")
        fmt.Fprintf(sh.w, "%v\n", s)
    }
    return
}

