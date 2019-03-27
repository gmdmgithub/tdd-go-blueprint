# Sample TTD driven approach to API responding

## Problem to solve

```go
func handle(w http.ResponseWriter, r *http.Request){
    if err:=SomeAction(); err != nil{
        respond.Err(w,r,err)
    }
    respond.Success(w,r,http.StatusOK,"Happy answer!")
    
    .... 
}
```

## Add is - testing helper for Go

```
go get github.com/cheekybits/is
```