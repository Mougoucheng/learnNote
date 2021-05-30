## json.Marshal规则

```go
type example struct {
    // Field appears in JSON as key "myName".
    Field int `json:"myName"` //Json翻译为:"myName"
    
    // Field appears in JSON as key "myName" and
    // the field is omitted from the object if its value is empty,
    // as defined above.
    Field int `json:"myName,omitempty"` //Json翻译为:"myName或者零值省略"
    
    // Field appears in JSON as key "Field" (the default), but
    // the field is skipped if empty.
    // Note the leading comma.
    Field int `json:",omitempty"` //Json翻译为:"Field或者零值省略"
    
    // Field is ignored by this package.
    Field int `json:"-"` //Json翻译一直省略
    
    // Field appears in JSON as key "-".
    Field int `json:"-,"` //Json翻译为:"-"
}
```