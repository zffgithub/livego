# livego
live streaming server write in pure go, simple efficient and can run in any platform.

<a target='_blank' rel='nofollow' href='https://app.codesponsor.io/link/3bvxELAxnq8r4wheFyRkED8U/gwuhaolin/livego'>
  <img alt='Sponsor' width='888' height='68' src='https://app.codesponsor.io/embed/3bvxELAxnq8r4wheFyRkED8U/gwuhaolin/livego.svg' />
</a>

## Support
#### Transport protocol
- [x] RTMP
- [x] AMF
- [x] HLS
- [x] HTTP-FLV
#### File container
- [x] FLV
- [x] TS
#### AV coder
- [x] H264
- [x] AAC
- [x] MP3

## Install
### Download Bin
[releases](https://github.com/gwuhaolin/livego/releases)

### Docker
TODO

### Install System Service
TODO

### Build From Source code
1. run `git clone https://github.com/gwuhaolin/livego.git`
2. cd to livego dir then run `go build`

## Use
2. run  `livego` to start livego server
3. push `RTMP` stream to `rtmp://localhost:1935/live/movie`, eg use `ffmpeg -re -i demo.flv -c copy -f flv rtmp://localhost:1935/live/movie`
4. play live stream form:
    - `RTMP`:`rtmp://localhost:1935/live/movie`
    - `FLV`:`http://127.0.0.1:7001/live/movie.flv`
    - `HLS`:`http://127.0.0.1:7002/live/movie.m3u8`
    
    
## Roadmap
1. support config file - 20%
2. add unit tests and continuous integration - 0%
3. improve stable and reliable for production use - 0%
4. rewrite docs - 0%
