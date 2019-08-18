# yada - Yet Another Downloader App

`yada` is a sample `golang` app that can download a list of URLs locally, either serially or concurrently.

## Getting started

To get the code, run :
```
go get https://github.com/Abhishekvrshny/yada
```
Get to the root of the project :
```
cd $GOPATH/src/github.com/Abhishekvrshny/yada
```
To run the code :
```
go run api/main.go
```
## Check if the app is up and running
`yada` runs of port 8081 by default. To change this, edit this in :
```
api/main.go
```
Hit the `health` check API :
```bash
$ curl -I http://localhost:8081/health
HTTP/1.1 200 OK
Date: Sun, 18 Aug 2019 11:08:55 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8
```
## `yada` user guide
### Serial Download
***Request***
```bash
$ curl -X POST http://localhost:8081/downloads -H 'Content-Type: application/json' -d '{
        "type":"serial",
        "urls":[
"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG"
        ]
}'
```
***Response***
```json
{"id":"a9c2439e-d37b-54ef-1fd8-20eed578c33d"}
```
### Concurrent Download
***Request***
```bash
$ curl -X POST http://localhost:8081/downloads -H 'Content-Type: application/json' -d '{
        "type":"concurrent",
        "urls":[
"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG"
        ]
}'
```
***Response***
```json
{"id":"a9c2439e-d37b-54ef-1fd8-20eed578c33d"}
```
### Check Download Status
***Request***
```
$ curl http://localhost:8081/downloads/a9c2439e-d37b-54ef-1fd8-20eed578c33d
```
***Response***
```json
{"id":"a9c2439e-d37b-54ef-1fd8-20eed578c33d",
"start_time":"2019-08-18 16:45:28.973436 +0530 IST m=+2183.303782741",
"end_time":"2019-08-18 16:45:34.447567 +0530 IST m=+2188.777882798",
"status":"SUCCESSFUL",
"download_type":"CONCURRENT",
"files":{
"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg":"/tmp/a9c2439e-d37b-54ef-1fd8-20eed578c33d/cf81fbfa-7697-12f3-189a-72fbba9b6556",
"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG":"/tmp/a9c2439e-d37b-54ef-1fd8-20eed578c33d/30ed4f5e-08f5-9220-483e-8559da75d02b"
}
}
```
### Browse Downloaded Files
***Request***
```bash
$ curl -L http://localhost:8081/files
```
***Response***
```html
<pre>
<a href="84b62179-c1f7-814e-cf5d-c852b46ecf91/">84b62179-c1f7-814e-cf5d-c852b46ecf91/</a>
<a href="a9c2439e-d37b-54ef-1fd8-20eed578c33d/">a9c2439e-d37b-54ef-1fd8-20eed578c33d/</a>
</pre>
```
