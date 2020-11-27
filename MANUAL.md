# Manual

Here you will find all necessary commands to transcribe a audio file

## Converting audio files to flac
```
ffmpeg -i my-original-audio.m4a -f flac my-converted-audio.flac
```

## Extracting audio files from videos
```
ffmpeg -i my-video.mp4 -f flac -vn music.mp3
```

## Uploading file to Google

Now, that you have a flac file at hand you can upload it to your [Google Cloud Storage](https://console.cloud.google.com/storage/browser). 

Once the upload is complete, copy the file URL, it will look like `gs://.../my-coverted-audio.flac`.

## Transcript
```
transcript -f gs://.../my-coverted-audio.flac
```

In case you want to send this to a file just do:

```
transcript -f gs://.../my-coverted-audio.flac > my-transcript.txt
```
