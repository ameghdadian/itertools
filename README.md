# itertools
**itertools** provide tools to simplify working with Iterators(officially added to the language in Go 1.23 release). 

## Example

```go
    package main

    import (
        "fmt"
        "slices"
        "github.com/ameghdadian/itertools"
    )

    func main() {
        a := []int{1,2,3}
        b := []int{4,5,6}

        // Combine arbitrary number of slices into an iterator
        it := itertools.Concat(a,b)
        for v := range it {
            fmt.Println(v)
        }

        // Combine arbitrary number of iterators into a single iterator
        it1 := slices.Values(a)
        it2 := slices.Values(b)

        combined := itertools.ConcatIter(a, b)
        for v := range combined {
            fmt.Println(v)
        }

        // ...
    }
```

## License
This is project is licensed under the Apache Version 2.0 License. See [LICENSE](LICENSE) for details.