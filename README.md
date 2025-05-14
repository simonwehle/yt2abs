# yt2abs

> [!WARNING]  
> This project uses ffmpeg make sure you have it installed

yt2abs is a cli tool that converts .mp3 files (from youtube) into an .m4b audiobook with chapters and metadata from audible to be used in audiobookshelf

### usage:

```
yt2abs --asin B07KKPR52P
```

default file names:

```
folder/
├── audiobook.mp3
└── chapters.txt
```

`chapters.txt` example

```
0:00:00 Introduction
0:00:69 Part 1: First Part
0:04:20 1. First Chapter
0:13:37 Chapter without prefix
6:94:20 End
```

### build

```
go build -o yt2abs .
sudo mv yt2abs /usr/local/bin/
```
