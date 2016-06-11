go-ansicfile
============

Package for operating ANSI-C's `FILE*` with Go.

- `ansicfile.FilePtr` is binary compatible to `FILE*` of msvcrt.dll .
- `ansicfile.FilePtr` satisfies io.Write/io.Reader.

        package main

        import (
            "bufio"
            "fmt"
            "os"

            "github.com/zetamatta/go-ansicfile"
        )

        func main() {
            fp, err := ansicfile.Open("test", "w")
            if err != nil {
                fmt.Fprintln(os.Stderr, err.Error())
                return
            }
            for i := 0; i < 5; i++ {
                fmt.Fprintf(fp, "test: %d\n", i)
            }
            fp.Close()

            fp, err = ansicfile.Open("test", "r")
            if err != nil {
                fmt.Fprintln(os.Stderr, err.Error())
                return
            }
            scanner := bufio.NewScanner(fp)
            for scanner.Scan() {
                fmt.Println(scanner.Text())
            }
            fp.Close()
        }
