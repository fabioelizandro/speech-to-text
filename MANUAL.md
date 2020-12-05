# Manual

Here you will find all necessary commands to transcribe a audio file

## Converting audio files to flac
```shell
ffmpeg -i my-original-audio.m4a -ac 1 -f flac my-converted-audio.flac
```

## Extracting audio files from videos
```shell
ffmpeg -i my-video.mp4 -ac 1 -f flac -vn music.mp3
```

## Uploading file to Google

⚠️ NOTE ⚠️ : It's important that you keep your files with meaningful names to avoid name collisions. <br>
I suggest a format like `interview-$subject-$date.flac` where `$subject` is the interviewee name<br>
and `$date` is when it happened.

<hr>

Now, that you have a flac file at hand you can upload it to your [Google Cloud Storage](https://console.cloud.google.com/storage/browser).

Once the upload is complete, copy the file URL, it will look like `gs://.../my-coverted-audio.flac`.

## Transcript
```shell
transcript -f gs://.../my-coverted-audio.flac
```

For saving the transcript to a file do:

```shell
transcript -f gs://.../my-coverted-audio.flac > my-transcript.txt
```
