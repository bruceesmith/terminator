[![Go Reference][goreference_badge]][goreference_link]
[![Go Report Card][goreportcard_badge]][goreportcard_link]
 
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# terminator

```go
import "github.com/bruceesmith/terminator"
```

Package terminator permits orderly stopping / shutdown of a group of goroutines via methods which mimic stop of a [sync.WaitGroup](<https://pkg.go.dev/sync/#WaitGroup>).There is a default Terminator accessible through top level functions \(Add, Done, Wait and so on\) that call the corresponding Terminator methods

## Index

- [func Add\(delta int\)](<#Add>)
- [func Done\(\)](<#Done>)
- [func SetDefault\(t \*Terminator\)](<#SetDefault>)
- [func ShutDown\(\) \<\-chan struct\{\}](<#ShutDown>)
- [func ShuttingDown\(\) bool](<#ShuttingDown>)
- [func Stop\(\)](<#Stop>)
- [func Wait\(\)](<#Wait>)
- [type Terminator](<#Terminator>)
  - [func Default\(\) \*Terminator](<#Default>)
  - [func New\(\) \*Terminator](<#New>)
  - [func \(t \*Terminator\) Add\(delta int\)](<#Terminator.Add>)
  - [func \(t \*Terminator\) Done\(\)](<#Terminator.Done>)
  - [func \(t \*Terminator\) ShutDown\(\) \<\-chan struct\{\}](<#Terminator.ShutDown>)
  - [func \(t \*Terminator\) ShuttingDown\(\) bool](<#Terminator.ShuttingDown>)
  - [func \(t \*Terminator\) Stop\(\)](<#Terminator.Stop>)
  - [func \(t \*Terminator\) Wait\(\)](<#Terminator.Wait>)


<a name="Add"></a>
## func Add

```go
func Add(delta int)
```

Add adds delta to the count of goroutines in the group

<a name="Done"></a>
## func Done

```go
func Done()
```

Done decrements the count of goroutines in the group by one

<a name="SetDefault"></a>
## func SetDefault

```go
func SetDefault(t *Terminator)
```

SetDefault sets the default Terminator

<a name="ShutDown"></a>
## func ShutDown

```go
func ShutDown() <-chan struct{}
```

ShutDown allows code to wait for a shut down signal

<a name="ShuttingDown"></a>
## func ShuttingDown

```go
func ShuttingDown() bool
```

ShuttingDown returns true if shutdown is in progress

<a name="Stop"></a>
## func Stop

```go
func Stop()
```

Stop signals that all goroutines in the group should safely exit

<a name="Wait"></a>
## func Wait

```go
func Wait()
```

Wait blocks until every goroutines in the group has called Done\(\)

<a name="Terminator"></a>
## type Terminator

Terminator manages groups of goroutines

```go
type Terminator struct {
    // contains filtered or unexported fields
}
```

<a name="Default"></a>
### func Default

```go
func Default() *Terminator
```

Default returns the default [Terminator](<#Terminator>).

<a name="New"></a>
### func New

```go
func New() *Terminator
```

New returns a Terminator

<a name="Terminator.Add"></a>
### func \(\*Terminator\) Add

```go
func (t *Terminator) Add(delta int)
```

Add adds delta to the count of goroutines in the group

<a name="Terminator.Done"></a>
### func \(\*Terminator\) Done

```go
func (t *Terminator) Done()
```

Done decrements the count of goroutines in the group by one

<a name="Terminator.ShutDown"></a>
### func \(\*Terminator\) ShutDown

```go
func (t *Terminator) ShutDown() <-chan struct{}
```

ShutDown allows code to wait for a shut down signal

<a name="Terminator.ShuttingDown"></a>
### func \(\*Terminator\) ShuttingDown

```go
func (t *Terminator) ShuttingDown() bool
```

ShuttingDown returns true if shutdown is in Default\(\).progress

<a name="Terminator.Stop"></a>
### func \(\*Terminator\) Stop

```go
func (t *Terminator) Stop()
```

Stop signals that all goroutines in the group should safely exit

<a name="Terminator.Wait"></a>
### func \(\*Terminator\) Wait

```go
func (t *Terminator) Wait()
```

Wait blocks until every goroutines in the group has called Done\(\)

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
 
[goreference_badge]: https://pkg.go.dev/badge/github.com/bruceesmith/terminator/v3.svg
[goreference_link]: https://pkg.go.dev/github.com/bruceesmith/terminator
[goreportcard_badge]: https://goreportcard.com/badge/github.com/bruceesmith/terminator
[goreportcard_link]: https://goreportcard.com/report/github.com/bruceesmith/terminator
