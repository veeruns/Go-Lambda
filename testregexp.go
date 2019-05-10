package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scn := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("Enter Lines:")
        var lines []string
        for scn.Scan() {
            line := scn.Text()
            if len(line) == 1 {
                // Group Separator (GS ^]): ctrl-]
                if line[0] == '\x1D' {
                    break
                }
            }
            lines = append(lines, line)
        }

        if len(lines) > 0 {
            fmt.Println()
            fmt.Println("Result:")
            for _, line := range lines {
                fmt.Println(line)
            }
            fmt.Println()
        }

        if err := scn.Err(); err != nil {
            fmt.Fprintln(os.Stderr, err)
            break
        }
        if len(lines) == 0 {
            break
        }
    }
}
