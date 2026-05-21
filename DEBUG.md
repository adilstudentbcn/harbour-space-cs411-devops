
## Hypotheses
1. Incompatible glibc version (Most Likely): The Go binary was dynamically linked against a newer version of the `glibc` library on the Jenkins build machine, but the customer's older Ubuntu 18.04 VM only has an older version installed.

2. Missing shared library: The binary is attempting to dynamically link, but the customer's system doesn't have the standard C library in the path it expects.



## Verifications
To prove this, I would run `ldd ./main` on the customer's machine. This command lists all the dynamic libraries the application requires to run. If it points to missing or older versions of `libc.so.6`, my hypothesis is correct.


## The Fix
I would rebuild the application using this command: 
`CGO_ENABLED=0 go build -o main main.go`

`CGO_ENABLED=0` disables the use of C code and forces the Go compiler to create a 100% statically linked binary. It packs every single instruction it needs directly into the file, meaning it will never ask the customer's OS for `glibc`.


## The Lesson
Go binaries dynamically link to standard system libraries like `glibc` by default, so you must explicitly compile them as statically linked if you want true portability across different Linux versions.
