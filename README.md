# yt2abs

yt2abs is a cli tool that converts .mp3 files (from youtube) into an .m4b audiobook with chapters and metadata from audible to be used in audiobookshelf

> [!IMPORTANT]  
> [ffmpeg](https://ffmpeg.org/) is required to use yt2abs

## installation:

```
brew tap simonwehle/tools
brew install yt2abs
```

## usage:

make sure you use the ASIN from [audible.com](https://audible.com)

### file mode

to convert a single audio file to a .m4b audiobook

```
yt2abs -a B07KKMNZCH
```

full auto mode uses the default file names

```
folder/
├── audiobook.mp3
└── chapters.txt
```

Make sure to format all timestamps in the format "H:MM:SS" and add a chapter named "End" with the full lenght of the file.

`chapters.txt` example:

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

folder and file structure

```
Harry Potter and the Sorcerer's Stone/
├── 00 - Introduction.mp3
├── 01 - The Boy Who Lived.mp3
├── 02 - The Vanishing Glass.mp3
└── 03 - The Letters From No One.mp3
```

### manual metadata

When no ASIN is supplied, yt2abs looks for `metadata.yml` beside the input
audio file, or inside the input folder. A missing file keeps the usual title
fallback behavior. If both `-a` and `metadata.yml` are present, Audible data
takes precedence. A manual title takes precedence over `-t`.

Use `-m` to load metadata from a specific file, for example `yt2abs -i book.mp3
-m book.yml`.

```yaml
title: My Audiobook
subtitle: Optional Subtitle
release_date: "2024-01-15"
publisher_name: Example Publisher
publisher_summary: Description text
authors:
  - name: Author Name
narrators:
  - name: Narrator Name
product_images:
  "500": ./cover.jpg
category_ladders:
  - root: Fiction
    ladder:
      - id: fiction
        name: Fiction
```

The `product_images."500"` value may be an HTTP/HTTPS URL, an absolute path,
or a path relative to `metadata.yml`.

## build

```
go build -o yt2abs .
sudo mv yt2abs /usr/local/bin/
```

uninstall

```
sudo rm /usr/local/bin/yt2abs
rm -rf ~/.yt2abs
```
