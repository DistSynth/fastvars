# fastvars

_fastvars_ is [fasttemplate](https://github.com/valyala/fasttemplate) wrapper with dictionary substitusion ability.

## Usage

```go
fv, err := NewFastVars()
if err != nil {
    log.Fatal("NewFastVars() failed")
}
fv.Append(map[string]interface{}{
	"IP":   "127.0.0.1",
	"PORT": 80,
	"URL":  "http://#{IP}:#{PORT}/",
})
fmt.Println(fv.Get("URL"))
//returns http://127.0.0.1:80/
```
