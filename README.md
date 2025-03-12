<div align="center">
    <h1>itertools</h1>
    <a href="https://goreportcard.com/report/github.com/ameghdadian/itertools">
     <img src="https://goreportcard.com/badge/github.com/ameghdadian/itertools" height="20" alt="Go Report Card">
    </a>
    <a href="https://pkg.go.dev/github.com/ameghdadian/itertools">
        <img src="https://pkg.go.dev/badge/github.com/ameghdadian/itertools.svg" height="20" alt="Go Reference">
    </a>
    <a href="https://github.com/ameghdadian/itertools/actions/workflows/github-actions.yaml/badge.svg">
        <img src="https://github.com/ameghdadian/itertools/actions/workflows/github-actions.yaml/badge.svg" height="20" alt="CI">
    </a>
    <a href="#">
     <img src="https://img.shields.io/coverallsCoverage/github/ameghdadian/itertools" height="20" alt="Code Test Coverage">
    </a>

  <h3><em>itertools provides tools to simplify working with Iterators(officially introduced in Go 1.23 release).</em></h3>
  <h3>NOTE: this repository is archived. Use <a href="https://github.com/ameghdadian/iter">iter</a> instead.</h3>
</div>


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

        combined := itertools.ConcatIter(it1, it2)
        for v := range combined {
            fmt.Println(v)
        }

        // ...
    }
```

## License
This is project is licensed under the Apache Version 2.0 License. See [LICENSE](LICENSE) for details.
