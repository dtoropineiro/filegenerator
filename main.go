package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Debe proporcionar el tamaño del archivo en MB")
        return
    }
    sizeMB, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("El tamaño del archivo debe ser un número entero")
        return
    }
    filesize := int64(sizeMB) * 1024 * 1024

    filename := fmt.Sprintf("file-%dmb.txt", sizeMB)
    f, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    data := make([]byte, 1024)
    for i := int64(0); i < filesize; i += int64(len(data)) {
        rand.Read(data)
        f.Write(data)
    }

    fmt.Printf("Archivo generado: %s\n", filename)
}
