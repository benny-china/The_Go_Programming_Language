package main
import (
    "net/http"
    "path"
    "os"
    "io"
    "fmt"
)

func fetch(url string)(filename string, n int64, err error) {
    resp,err := http.Get(url)
    if err!=nil {
        return "",0,err
    }
    defer resp.Body.Close()
    local := path.Base(resp.Request.URL.Path)
    if local == "/" {
        local = "index.html"    
    }
    f,err := os.Create(local)
    if err != nil {
        return "",0,err
    }

    n,err = io.Copy(f,resp.Body)
    if closeErr := f.Close(); err == nil {
        err = closeErr
    }
    return local,n,err    

}

func main() {
    file,n,err := fetch(os.Args[1])
    if err!= nil {
        fmt.Println(err)
    }else {
        fmt.Println(file,n)
    }
}
