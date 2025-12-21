# Go Escape Analysis Guide

## Table of Contents
- [What is Escape Analysis?](#what-is-escape-analysis)
- [How to Execute Escape Analysis](#how-to-execute-escape-analysis)
- [Understanding the Output](#understanding-the-output)
- [Analysis Techniques](#analysis-techniques)
- [Best Practices](#best-practices)
- [UI Tools and Visualization](#ui-tools-and-visualization)
- [Common Patterns and Pitfalls](#common-patterns-and-pitfalls)

## What is Escape Analysis?

Escape analysis is a compile-time optimization technique used by the Go compiler to determine whether a variable can be allocated on the stack or must be allocated on the heap. Variables that "escape" to the heap require garbage collection, while stack-allocated variables are automatically freed when the function returns.

**Benefits of understanding escape analysis:**
- Reduced garbage collection pressure
- Improved application performance
- Lower memory allocations
- Better cache locality

## How to Execute Escape Analysis

### Basic Command

```bash
go build -gcflags='-m' ./...
```

### Multiple Levels of Detail

```bash
# Level 1 - Basic escape analysis
go build -gcflags='-m' ./...

# Level 2 - More detailed information
go build -gcflags='-m -m' ./...

# Level 3 - Very detailed (including function inlining)
go build -gcflags='-m -m -m' ./...

# Level 4 - Maximum detail
go build -gcflags='-m -m -m -m' ./...
```

### For Specific Package or File

```bash
# Specific package
go build -gcflags='-m' ./internal/mypackage

# Specific file
go build -gcflags='-m' ./internal/mypackage/file.go

# Current directory
go build -gcflags='-m' .
```

### Disable Optimizations (for debugging)

```bash
# Disable all optimizations
go build -gcflags='-m -N -l' ./...
# -N: disable optimizations
# -l: disable inlining
```

### Save Output to File

```bash
go build -gcflags='-m -m' ./... 2> escape-analysis.txt
```

## Understanding the Output

### Common Messages

| Message | Meaning |
|---------|---------|
| `moved to heap: <var>` | Variable escapes and is allocated on the heap |
| `<var> escapes to heap` | Variable or its reference escapes to heap |
| `can inline <func>` | Function is small enough to be inlined |
| `inlining call to <func>` | Function call is being inlined |
| `leaking param: <param>` | Parameter escapes through return or assignment |
| `<var> does not escape` | Variable stays on the stack |
| `too complex for escape analysis` | Compiler cannot determine, defaults to heap |

### Example Output Analysis

```go
// Code
func createUser(name string) *User {
    user := &User{Name: name}
    return user
}

// Output
./main.go:10:6: moved to heap: user
./main.go:10:6: leaking param: name
```

**Interpretation:**
- `user` escapes because its pointer is returned
- `name` parameter leaks because it's stored in the returned struct

## Analysis Techniques

### 1. Iterative Analysis

```bash
# Run analysis
go build -gcflags='-m -m' ./... 2>&1 | grep "escapes to heap"

# Focus on specific types
go build -gcflags='-m -m' ./... 2>&1 | grep "User.*escapes"

# Count escapes
go build -gcflags='-m' ./... 2>&1 | grep -c "escapes to heap"
```

### 2. Benchmark Before and After

```bash
# Run benchmarks with memory stats
go test -bench=. -benchmem ./...

# Profile memory allocations
go test -bench=. -memprofile=mem.out ./...
go tool pprof -alloc_space mem.out
```

### 3. Compare with Assembly

```bash
# Generate assembly output
go build -gcflags='-S' ./... > assembly.txt

# Look for heap allocations (runtime.newobject calls)
grep "runtime.newobject" assembly.txt
```

### 4. Use Build Tags for Analysis

```go
//go:build analysis
// +build analysis

package main

// Add debug code for analysis builds
```

```bash
go build -tags=analysis -gcflags='-m -m' ./...
```

## Best Practices

### 1. Profile Before Optimizing

- Always measure with real workloads
- Use `go test -bench` and `pprof`
- Focus on hot paths identified by profiling

### 2. Common Optimization Techniques

#### Return Values Instead of Pointers

```go
// ❌ Causes escape
func NewConfig() *Config {
    cfg := Config{Size: 100}
    return &cfg
}

// ✅ No escape
func NewConfig() Config {
    return Config{Size: 100}
}
```

#### Use Preallocated Buffers

```go
// ❌ Allocates on each call
func processData(data []byte) []byte {
    result := make([]byte, len(data))
    // process...
    return result
}

// ✅ Reuse buffer
type Processor struct {
    buffer []byte
}

func (p *Processor) processData(data []byte) []byte {
    if cap(p.buffer) < len(data) {
        p.buffer = make([]byte, len(data))
    }
    result := p.buffer[:len(data)]
    // process...
    return result
}
```

#### Avoid Interface Conversions in Hot Paths

```go
// ❌ Interface conversion may cause escape
func process(v interface{}) {
    // ...
}

// ✅ Use generic types (Go 1.18+)
func process[T any](v T) {
    // ...
}
```

#### Limit Slice Growth

```go
// ❌ May cause multiple allocations
func appendMany() []int {
    var result []int
    for i := 0; i < 1000; i++ {
        result = append(result, i)
    }
    return result
}

// ✅ Preallocate capacity
func appendMany() []int {
    result := make([]int, 0, 1000)
    for i := 0; i < 1000; i++ {
        result = append(result, i)
    }
    return result
}
```

### 3. Accept Escaping When Necessary

Not all heap allocations are bad:
- Long-lived objects should be on heap
- Large objects (>10KB) are better on heap
- Shared objects across goroutines need heap allocation
- Don't sacrifice code clarity for micro-optimizations

### 4. Testing Strategy

```go
func TestNoEscapes(t *testing.T) {
    testing.AllocsPerRun(100, func() {
        // Code that should not allocate
        result := fastFunction()
        _ = result
    })
}
```

## UI Tools and Visualization

### 1. **Go Escape Analyzer (VSCode Extension)**

```
Extension ID: golang.go
```

**Features:**
- Inline escape analysis hints
- CodeLens showing allocation information
- Integration with Go language server

**Setup:**
```json
// settings.json
{
    "go.buildFlags": ["-gcflags='-m'"],
    "go.lintTool": "golangci-lint"
}
```

### 2. **GoLand/IntelliJ IDEA**

**Features:**
- Built-in escape analysis
- Memory allocation inspections
- Performance hints

**Usage:**
- Right-click → Run with Profiler
- Analyze → Inspect Code → Performance Issues

### 3. **go-torch (Flame Graphs)**

```bash
# Install
go install github.com/uber/go-torch@latest

# Generate flame graph from profile
go test -bench=. -cpuprofile=cpu.out
go-torch cpu.out
```

### 4. **pprof Web UI**

```bash
# Memory profile
go test -bench=. -memprofile=mem.out
go tool pprof -http=:8080 mem.out

# CPU profile
go test -bench=. -cpuprofile=cpu.out
go tool pprof -http=:8080 cpu.out
```

**Web UI Features:**
- Interactive flame graphs
- Call graphs
- Source code annotation
- Allocation breakdown

### 5. **escape-analysis-vis (Third-party)**

```bash
# Install
go install github.com/loov/goda@latest

# Visualize dependencies and potential escapes
goda graph "reach(./...)" | dot -Tsvg -o graph.svg
```

### 6. **Custom Analysis Scripts**

```bash
#!/bin/bash
# analyze-escapes.sh

echo "Running escape analysis..."
go build -gcflags='-m -m' ./... 2>&1 | \
    grep -E "moved to heap|escapes to heap" | \
    sort | uniq -c | sort -nr > escapes-summary.txt

echo "Top escape locations:"
head -20 escapes-summary.txt
```

### 7. **Continuous Integration Integration**

```yaml
# .github/workflows/escape-analysis.yml
name: Escape Analysis

on: [pull_request]

jobs:
  analyze:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run Escape Analysis
        run: |
          go build -gcflags='-m' ./... 2>&1 | tee escape-output.txt
          
      - name: Check for New Escapes
        run: |
          # Compare with baseline
          ./scripts/check-escapes.sh escape-output.txt
```

## Common Patterns and Pitfalls

### Patterns That Cause Escapes

1. **Returning Pointer to Local Variable**
```go
func bad() *int {
    x := 42
    return &x  // x escapes
}
```

2. **Storing Pointer in Interface**
```go
func bad() interface{} {
    x := 42
    return &x  // x escapes
}
```

3. **Sending Pointer Through Channel**
```go
func bad(ch chan *int) {
    x := 42
    ch <- &x  // x escapes
}
```

4. **Closure Capturing Variables**
```go
func bad() func() int {
    x := 42
    return func() int {
        return x  // x escapes
    }
}
```

5. **Assignment to Global**
```go
var global *int

func bad() {
    x := 42
    global = &x  // x escapes
}
```

### Patterns That Prevent Escapes

1. **Pass by Value for Small Structs**
```go
type Point struct { X, Y int }

func good(p Point) Point {
    return Point{p.X + 1, p.Y + 1}
}
```

2. **Use Stack Arrays for Known Sizes**
```go
func good() [100]byte {
    var buf [100]byte
    // process...
    return buf
}
```

3. **Inline Small Functions**
```go
//go:inline
func add(a, b int) int {
    return a + b
}
```

## Monitoring and Reporting

### Create Dashboard Metrics

```go
import (
    "runtime"
    "time"
)

func reportAllocations() {
    var m runtime.MemStats
    ticker := time.NewTicker(10 * time.Second)
    
    for range ticker.C {
        runtime.ReadMemStats(&m)
        log.Printf("Allocs: %d, HeapAlloc: %d MB", 
            m.Alloc, m.HeapAlloc/1024/1024)
    }
}
```

### Benchmark Comparison

```bash
# Before changes
go test -bench=. -benchmem > old.txt

# After changes
go test -bench=. -benchmem > new.txt

# Compare
benchstat old.txt new.txt
```

## Resources

- [Go Compiler Directives](https://pkg.go.dev/cmd/compile)
- [Dave Cheney - High Performance Go](https://dave.cheney.net/high-performance-go)
- [Golang Escape Analysis Documentation](https://go.dev/doc/diagnostics)
- [Go Performance Book](https://github.com/dgryski/go-perfbook)

## Quick Reference Card

```bash
# Basic analysis
go build -gcflags='-m' ./...

# Detailed analysis
go build -gcflags='-m -m' ./... 2>&1 | less

# With benchmarks
go test -bench=. -benchmem -gcflags='-m'

# Memory profile
go test -memprofile=mem.out -bench=.
go tool pprof -http=:8080 mem.out

# Disable optimizations
go build -gcflags='-m -N -l'

# Check specific file
go build -gcflags='-m' path/to/file.go
```

---

**Last Updated:** December 2025
**Go Version:** 1.21+
