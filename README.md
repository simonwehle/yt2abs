# audiobook-converter

```
ffmpeg \
  -i input.mp3 \
  -i cover.jpg \
  -i chapters.txt \
  -map 0:a \
  -map 1 \
  -map_metadata 2 \
  -c:a aac \
  -b:a 64k \
  -vn \
  -movflags +faststart \
  output.m4b
```