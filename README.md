# yt2abs

yt2abs is a cli tool that converts .mp3 files (from youtube) into an .m4b audiobook with chapters and metadata from audible to be used in audiobookshelf

> [!IMPORTANT]  
> [ffmpeg](https://ffmpeg.org/) is required to use yt2abs

## usage:

### file mode

to convert a single audio file to a .m4b audiobook

#### full auto

⚠️ Make sure you use the ASIN from audible.com

```
yt2abs -a B07KKMNZCH
```

full auto mode uses the default file names

```
folder/
├── audiobook.mp3
└── chapters.txt
```

`chapters.txt` example

⚠️ Make sure to format all timestamps in the format "H:MM:SS" and add a chapter named "End" with the full lenght of the file

```
0:00:00 Introduction
0:00:69 Part 1: First Part
0:04:20 1. First Chapter
0:13:37 Chapter without prefix
6:94:20 End
```

### folder mode

folder mode can merge multiple chapter files to one audiobook

```
yt2abs -a B017V4IM1G -f .
```

Folder and file structure

```
Harry Potter and the Sorcerer's Stone/
├── 00 - Introduction.mp3
├── 01 - The Boy Who Lived.mp3
├── 02 - The Vanishing Glass.mp3
└── 03 - The Letters From No One.mp3
```

## build

```
go build -o yt2abs .
sudo mv yt2abs /usr/local/bin/
```
