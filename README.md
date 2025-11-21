### go-face-detect
Real-time face detection using Cascade filter from video input device

# Finding video devices
```ls -l /dev/video*```

The output would look something like:
```
crw-rw----+ 1 root video 81, 0 Nov 20 01:31 /dev/video0
crw-rw----+ 1 root video 81, 1 Nov 20 01:31 /dev/video1
crw-rw----+ 1 root video 81, 2 Nov 20 01:31 /dev/video2
crw-rw----+ 1 root video 81, 3 Nov 20 01:31 /dev/video3
```
where 0-3 would be the device ID

# Build
```go mod tidy```
```go build```

# Run
```./go-face-detect 0 haarcascade_frontalface_default.xml```