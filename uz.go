package main

import (
    "fmt"
    "log"
    "archive/zip"
    "io"
    "os"
    "path/filepath"
)

func Input(question string) string {
    fmt.Print(question)
    var answer string
    fmt.Scanln(&answer)
    return answer
}

func UnZip(fn, dist string) error {
    r, err := zip.OpenReader(fn)
    if err != nil {
        return err
    }
    defer r.Close()
    for _, f := range r.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()
        path := filepath.Join(dist, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            tmpf, err := os.Create(path)
            if err != nil {
                return err
            }
            defer tmpf.Close()
            _, err = io.Copy(tmpf, rc)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func main() {
    var inp string
    var err error

    if len(os.Args) < 2 {
        inp = Input("File Name: ")
    } else {
        inp = os.Args[1]
    }
    dst := filepath.Join(filepath.Dir(inp), "out")
    err = UnZip(inp, dst)
    if err != nil {
        log.Fatal(err)
    }
}
