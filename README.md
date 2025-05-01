# GoTag

**GoTag** is a simple desktop application for viewing MP3 tag information and cover art.

## Features

- Display MP3 tag info (title, artist, album, genre, year, track, etc.)
- Show cover art (if available)
- Show lyrics (if available)
- ESC or Q shortcut to quit the app
- Modern look with cyan font color (`#00FFFF`)
- Custom window icon (use `icon.png`)

## Build Instructions

1. **Install Go and dependencies:**
    ```sh
    go get fyne.io/fyne/v2
    go get github.com/dhowden/tag
    ```

2. **Place your `icon.png` file in the same folder as the source code.**

3. **Build the app:**
    ```sh
    go build -o gotag.exe
    ```

## Usage

```sh
gotag.exe "path/to/your/file.mp3"
```

Example:
```sh
gotag.exe "C:\TEMP\01. Fallen.mp3"
```

## UI Structure

- **Tab "Tag Info"**: Cover art at the top, tag info below.
- **Tab "Cover & Lyrics"**: Cover art at the top, lyrics below (if available).

## Notes

- The window icon will only appear if `icon.png` is found and valid.
- If cover art does not appear, make sure your MP3 file has embedded cover art in JPEG/PNG format.
- Quit shortcut (`ESC`/`Q`) only works when the window is focused.

## License

MIT

---

## Author
[Hadi Cahyadi](mailto:cumulus13@gmail.com)
    
## Coffie
[![Buy Me a Coffee](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/cumulus13)

[![Donate via Ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/cumulus13)

 [Support me on Patreon](https://www.patreon.com/cumulus13)