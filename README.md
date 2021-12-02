# FizzBuzz

This is just a curiosity based on the following StackExchange: 
https://codegolf.stackexchange.com/questions/215216/high-throughput-fizz-buzz

Example valid output:

```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
```

## Hardware

2.6 GHz 6-Core Intel i7
32 GB RAM


## Results

### Python 3.9.7

```bash
 python3 main.py | pv > /dev/null             
^C73MiB 0:00:09 [20.1MiB/s] [                          <=>                                                                                                                   ]
```

### Go 1.17.2 Naive

```bash
 go run main.go | pv > /dev/null
^C.9MiB 0:00:05 [5.74MiB/s] [              <=>                                                                                                                   ]
signal: interrupt
```

## Go 1.17.2 Buffered

```bash
go run main.go -mode buffered | pv > /dev/null
^C52GiB 0:00:11 [ 740MiB/s] [                                <=>                                                                                                                   ]
signal: interrupt
```
