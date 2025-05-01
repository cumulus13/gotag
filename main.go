package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	_ "image/jpeg"
	_ "image/png"

	"github.com/dhowden/tag"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gotag file.mp3")
		return
	}
	filePath := os.Args[1]

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer f.Close()

	meta, err := tag.ReadFrom(f)
	if err != nil {
		fmt.Println("Failed to read metadata:", err)
		return
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("MP3 Tag Viewer")
	myWindow.Resize(fyne.NewSize(400, 600)) // Atur ukuran window di sini

	// Set icon window dari file PNG
	iconData, err := ioutil.ReadFile("icon.png")
	if err == nil {
		fmt.Println("Icon loaded, size:", len(iconData))
		resource := fyne.NewStaticResource("gotag.png", iconData)
		myWindow.SetIcon(resource)
	} else {
		fmt.Println("Failed to load icon.png:", err)
		myWindow.SetIcon(theme.FyneLogo())
	}

	// --- Tab 1: Info ---
	track, total := meta.Track()
	disc, totalDisc := meta.Disc()
	infoItems := []string{
		fmt.Sprintf("Title: %s", meta.Title()),
		fmt.Sprintf("Artist: %s", meta.Artist()),
		fmt.Sprintf("Album: %s", meta.Album()),
		fmt.Sprintf("Genre: %s", meta.Genre()),
		fmt.Sprintf("Year: %d", meta.Year()),
		fmt.Sprintf("Track: %d/%d", track, total),
		fmt.Sprintf("Disc: %d/%d", disc, totalDisc),
		fmt.Sprintf("Album Artist: %s", meta.AlbumArtist()),
		// fmt.Sprintf("Year: %s", meta.OriginalArtist()),
		fmt.Sprintf("Composer: %s", meta.Composer()),
		// fmt.Sprintf("ISRC: %s", meta.ISRC()),
		// fmt.Sprintf("Copyright: %s", meta.Copyright()),
		// fmt.Sprintf("Publisher: %s", meta.Publisher()),
		// fmt.Sprintf("URL: %s", meta.URL()),
		// fmt.Sprintf("Encoded by: %s", meta.Encoded()),
		fmt.Sprintf("Filetype: %s", meta.FileType()),
		fmt.Sprintf("Format: %s", meta.Format()),
		fmt.Sprintf("Comment: %s", meta.Comment()),
		// fmt.Sprintf("RAW: %s", meta.Raw()),
	}
	var infoLabels []fyne.CanvasObject
	for _, item := range infoItems {
		txt := canvas.NewText(item, color.NRGBA{R: 0x00, G: 0xFF, B: 0xFF, A: 0xFF})
		txt.TextSize = 16
		infoLabels = append(infoLabels, txt)
	}
	infoContainer := container.NewVBox(infoLabels...)

	// --- Tab 2: Cover & Lyrics ---
	var cover fyne.CanvasObject = widget.NewLabel("No cover found.")
	if pic := meta.Picture(); pic != nil {
		fmt.Println("Picture MIME type:", pic.MIMEType)
		fmt.Println("Picture data size:", len(pic.Data))
		img, _, err := image.Decode(bytes.NewReader(pic.Data))
		if err == nil {
			imageWidget := canvas.NewImageFromImage(img)
			imageWidget.FillMode = canvas.ImageFillContain
			imageWidget.SetMinSize(fyne.NewSize(200, 200)) // Atur ukuran minimal gambar
			cover = container.NewCenter(imageWidget)
		} else {
			fmt.Println("Cover decode error:", err)
		}
	} else {
		fmt.Println("No picture found in tag")
	}

	var lyrics fyne.CanvasObject
	if l := meta.Lyrics(); l != "" {
		lyricsLines := strings.Split(l, "\n")
		var lyricLabels []fyne.CanvasObject
		for _, line := range lyricsLines {
			txt := canvas.NewText(line, color.NRGBA{R: 0x00, G: 0xFF, B: 0xFF, A: 0xFF}) // #00FFFF
			txt.TextSize = 16
			lyricLabels = append(lyricLabels, txt)
		}
		lyrics = container.NewScroll(container.NewVBox(lyricLabels...))
	} else {
		lyrics = widget.NewLabel("No lyrics found.")
	}

	tabs := container.NewAppTabs(
		container.NewTabItem(
			"Tag Info",
			container.NewBorder(cover, nil, nil, nil, container.NewScroll(infoContainer)),
		),
		container.NewTabItem(
			"Cover & Lyrics",
			container.NewBorder(cover, nil, nil, nil, lyrics),
		),
	)

	// Tambahkan shortcut ESC dan Q untuk keluar
	myWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if key.Name == fyne.KeyEscape || key.Name == fyne.KeyQ {
			myApp.Quit()
		}
	})

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
