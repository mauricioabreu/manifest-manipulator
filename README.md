# Manifest Manipulator

Manifest Manipulator is a powerful tool designed to simplify the modification of video manifests. With this tool, you can easily filter video manifests based on bandwidth and frame rate, as well as set the desired manifest as the first one.

## Why?

Video content comes with its fair share of complexities. Certain video manifests may encounter compatibility issues with various devices due to factors such as frame rate, bitrate, or other tags that directly impact smooth playback.

<img src="img/manifest_manipulator.png" width="500" height="400">

The scenario depicted in the image above perfectly illustrates a common problem: devices incapable of playing videos at 60fps. To address this challenge, we introduce Manifest Manipulator — a solution designed to rewrite video manifests on the fly, just before delivering them to users.

## Features

* **Bandwidth Filtering**: Filter video manifests based on the desired minimum and maximum bandwidth. This allows you to select the appropriate video quality based on available devices. e.g some TVs don't play a too low bandwidth.

* **Frame Rate Filtering**: Filter video manifests based on the desired frame rate. You can specify the desired frame rate to ensure smooth playback on devices with different capabilities.

* **Set First Manifest**: Set a specific video manifest as the first one in the modified manifest. This feature is useful when you want to prioritize a particular video quality or adapt the manifest for specific player requirements. Players start playing from the first manifest.

## Demo

To demonstrate the functionality of the Video Manifest Modifier, we've set up a demo consisting of an HTML page with a video player. The demo involves two servers: one server hosting the HTML page and another server acting as a proxy to modify the manifest downloaded from the video HLS (HTTP Live Streaming).

```mermaid
graph LR
  A[Player] -->|Requests video manifest| B(Proxy Server);
  B -->|Intercepts and modifies the manifest| C[Original Streaming];
  C -->|Applies modifications| B;
  B -->|Returns modified manifest| A;

```

Steps to reproduce:

* Go do `demo` folder
* Run `go run main.go`
* Browse http://localhost:8080
* Play it!

When you access the video manifest through the proxy server, you might notice that only three quality options are available in the player. This limitation is a result of the proxy server manipulating the manifest by setting the maximum bandwidth to **3000000**.

<img src="img/limited.png" width="600" height="400">

To understand the changes, you can compare the original video manifest fetched directly from the server with the modified manifest obtained through the proxy server. Here's an example using _curl_ commands:

```console
$ curl https://cph-p2p-msl.akamaized.net/hls/live/2000341/test/master.m3u8
```

```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-STREAM-INF:BANDWIDTH=550172,RESOLUTION=256x106
level_0.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=1650064,RESOLUTION=640x266
level_1.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=2749539,RESOLUTION=1280x534
level_2.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=4947980,RESOLUTION=1920x800
level_3.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=8247438,RESOLUTION=1920x800
level_4.m3u8
```

However, when the manifest is fetched through the proxy server, you will observe a reduced number of quality options:

```console
$ curl http://localhost:9090/hls/live/2000341/test/master.m3u8
```

The modified manifest, obtained from the proxy server, limits the quality options as follows:

```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-STREAM-INF:BANDWIDTH=550172,RESOLUTION=256x106
level_0.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=1650064,RESOLUTION=640x266
level_1.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=2749539,RESOLUTION=1280x534
level_2.m3u8
```

In the modified manifest, the **BANDWIDTH** attribute for the remaining three quality options has been omitted, effectively limiting the available options in the video player.

By adjusting the maximum bandwidth value in the proxy server's configuration, you can control the number and quality of options presented to the viewer, ensuring an optimized viewing experience based on your specific requirements - like offering 4k only to premium users.
