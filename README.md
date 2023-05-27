# HLS Validator

Overcomplicated Go application used to validate HLS manifests using 
[Apple's HLS tools](https://developer.apple.com/documentation/http-live-streaming/using-apple-s-http-live-streaming-hls-tools) 
that generates validated json files using the `mediastreamvalidator`, and thereafter generates html reports with the use
of `hlsreport`. The application can validate a manifest n times if for some reason the manifest changes on each request.

## Requirements
The user must install the aforementioned
[Apple's HLS tools](https://developer.apple.com/documentation/http-live-streaming/using-apple-s-http-live-streaming-hls-tools)
, and have go `1.20` or above installed.

## Flags
```
-m, --manifest   
  manifest url to verify with HLS spec. 
-s, --save-valid        
  save reports for valid HLS manifests. 
  Validated jsons will always be generated.
-t, --times         
  number of times to validate url.
```

## Usage
```bash
hlsvalidator [-m manifest uri] [-t time] [-s]
```

## Example
```bash
hlsvalidator -m https://example.com/playlist.m3u8 -t 10 -s
```
This will result in a new `output` dir with a `validated` dir containing json reports of the validated manifest, and a 
`reports` dir containing HTML translated reports of the previously generated json reports.

## Installation
In order to build the application the following command can be used:
```bash
make build
```
This will generate a binary in the `bin/` folder.

If you wish to install this application the following command can be run:
```bash
make install
```


