# gizmo

## Hardware Device Registry Demo

```
.
├── cmd
│   └── aus
│       └── aus.go
├── LICENSE
├── pkg
│   ├── device
│   │   └── device.go
│   ├── devices
│   │   ├── devices.go
│   │   ├── v0001p0001v0001
│   │   │   └── device.go
│   │   ├── v0001p0003v0001
│   │   │   └── device.go
│   │   └── v106Bp1570v00F1
│   │       └── device.go
│   └── registry
│       └── registry.go
└── README.md
```

## Registry Usage

```
$ ./aus 
recognized devices:
[>] vendor=0001	productID=0003	version=0001	HITACHI Red Lightsaber
[>] vendor=106B	productID=1570	version=00F1	SAMSUNG Double-Ended Lightsaber
[>] vendor=0001	productID=0001	version=0001	HITACHI Blue Lightsaber
[*] found preferred lightsaber: SAMSUNG Double-Ended Lightsaber
[*] reading from device...
returned awesome double-ended lightsaber stuff from device vendor=106B productID=1570 version=F1 from input: Linux version 4.18.0-25-generic (buildd@lcy01-amd64-025) (gcc version 8.3.0 (Ubuntu 8.3.0-6ubuntu1~18.10.1)) #26-Ubuntu SMP Mon Jun 24 09:32:08 UTC 2019

[*] device wrote 263 bytes.
$
```
