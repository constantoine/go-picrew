package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Assets
var (
	Cheveux []image.Image
	Peaux   []image.Image
	Yeux    []image.Image
	Shirts  []image.Image
	Bouches []image.Image
)

func loadAssets(base string, imgs *[]image.Image) error {
	files, err := ioutil.ReadDir(base)
	if err != nil {
		return err
	}
	for i := range files {
		file, err := os.Open(filepath.Join(base, files[i].Name()))
		if err != nil {
			return err
		}
		img, _, err := image.Decode(file)
		file.Close()
		if err != nil {
			fmt.Println(filepath.Join(base, files[i].Name()))
			return err
		}
		*imgs = append(*imgs, img)
	}
	return nil
}

func loadAllAssets(base string) error {
	if err := loadAssets(filepath.Join(base, "CHEVEUX"), &Cheveux); err != nil {
		return err
	}
	if err := loadAssets(filepath.Join(base, "PEAUX"), &Peaux); err != nil {
		return err
	}
	if err := loadAssets(filepath.Join(base, "YEUX"), &Yeux); err != nil {
		return err
	}
	if err := loadAssets(filepath.Join(base, "TEE-SHIRT"), &Shirts); err != nil {
		return err
	}
	if err := loadAssets(filepath.Join(base, "BOUCHES"), &Bouches); err != nil {
		return err
	}
	return nil
}

func generate(base string, filename string, cheveuxID, peauxID, yeuxID, shirtsID, bouchesID int) error {
	rec := image.Rect(0, 0, 512, 512)
	finalImg := image.NewRGBA(rec)
	draw.Draw(finalImg, rec, Peaux[peauxID], rec.Min, draw.Src)
	draw.Draw(finalImg, rec, Yeux[yeuxID], rec.Min, draw.Over)
	draw.Draw(finalImg, rec, Bouches[bouchesID], rec.Min, draw.Over)
	draw.Draw(finalImg, rec, Cheveux[cheveuxID], rec.Min, draw.Over)
	draw.Draw(finalImg, rec, Shirts[shirtsID], rec.Min, draw.Over)
	out, err := os.Create(filepath.Join(base, filename))
	if err != nil {
		return err
	}
	defer out.Close()
	if err := png.Encode(out, finalImg); err != nil {
		return err
	}
	return nil
}

func main() {
	base := "/Users/crebert/cursus42/side_project/src/picrew/assets/"
	if err := loadAllAssets(base); err != nil {
		panic(err)
	}
	fmt.Printf("%d %d %d %d %d\n", len(Peaux), len(Yeux), len(Bouches), len(Cheveux), len(Shirts))
	for peauxID := range Peaux {
		for yeuxID := range Yeux {
			for bouchesID := range Bouches {
				for cheveuxID := range Cheveux {
					for shirtsID := range Shirts {
						if err := generate(base, fmt.Sprintf("../out/%d-%d-%d-%d-%d.png", peauxID, yeuxID, bouchesID, cheveuxID, shirtsID), cheveuxID, peauxID, yeuxID, shirtsID, bouchesID); err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}
}
