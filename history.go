package cbsh

import (
    "bufio"
    "fmt"
    "os"
    "path"

    "github.com/sbinet/liner"
)

var HISTORY_FILE = "./cbsh_history"

func LoadHistory(liner *liner.State, dir string) {
    historyFile := path.Join(dir, HISTORY_FILE)
    if dir != "" {
        ReadHistoryFromFile(liner, historyFile)
    }
}

func UpdateHistory(liner *liner.State, dir, line string) {
    liner.AppendHistory(line)
    if dir != "" {
        WriteHistoryToFile(liner, dir+HISTORY_FILE)
    }
}

func WriteHistoryToFile(liner *liner.State, path string) error {
    f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        return err
    }

    defer f.Close()

    writer := bufio.NewWriter(f)
    if _, err = liner.WriteHistory(writer); err != nil {
        fmt.Printf("Error updating %v file: %v\n", HISTORY_FILE, err)
    } else {
        writer.Flush()
    }
    return nil
}

func ReadHistoryFromFile(liner *liner.State, path string) error {
    if f, err := os.Open(path); err != nil {
        return err
    } else {
        defer f.Close()
        reader := bufio.NewReader(f)
        liner.ReadHistory(reader)
    }
    return nil
}
