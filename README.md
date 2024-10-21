# Unique IP Counter

The application processes large IPv4 datasets from a file and calculates the number of unique IPs efficiently handling all possible IPv4 addresses.


## How to use

Place your "ip_addresses" file in the project folder.
Expected file format is 1 IP per line: 

```
145.67.23.4
8.34.5.23
89.54.3.124
89.54.3.124
3.45.71.5
```


Run the app:
```
go run main.go
```
Run Tests:
```
go test -v
```


## Performance Summary

Test Environment:
CPU: Intel i7-8565U
RAM: 16 GB
OS: Windows 10

Processing a 120GB file with 1,000,000,000 unique IPs
Memory used: 784 MB
Processing time: 4,065 sec (~1 hour 7 min)


## Optimizations and Improvements

The application was designed to gain maximum performance while consuming as little memory as possible when processing huge files. Therefore, some refactoring and possible error handling were omitted on purpose.

No 3rd party libriaries were used on purpose as this is one of the task requirements.

The current version of the application consumes the same amount of memory regardless of the input file's size. But, as it consumes less than 800Mb of RAM it is considered an optimal approach suitable for most environments. 
