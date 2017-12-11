package interfaces

import (
  "fmt"
  "io"
  "os"
)

func Copy(in io.ReadSeeker, out io.Writer) error {
  w := io.MultiWriter(out, os.Stdout)

  if _, err := io.Copy(w, in); err != nil {
    return err
  }
  in.Seek(0, 0)

  buff := make([]byte, 64)
  if _, err := io.CopyBuffer(w, in, buff); err != nil {
    return err
  }

  fmt.Println()

  return nil
}
