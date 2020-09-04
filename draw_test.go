package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkDraw(b *testing.B) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	base := filepath.Dir(ex) + "/assets/"
	boucheFile, err := os.Open(base + "bouche.png")
	if err != nil {
		panic(err)
	}
	defer boucheFile.Close()
	boucheImg, _, err := image.Decode(boucheFile)
	if err != nil {
		panic(err)
	}
	cheveuxFile, err := os.Open(base + "cheveux.png")
	if err != nil {
		panic(err)
	}
	defer cheveuxFile.Close()
	cheveuxImg, _, err := image.Decode(cheveuxFile)
	if err != nil {
		panic(err)
	}
	nezFile, err := os.Open(base + "nez.png")
	if err != nil {
		panic(err)
	}
	defer nezFile.Close()
	nezImg, _, err := image.Decode(nezFile)
	if err != nil {
		panic(err)
	}
	teteFile, err := os.Open(base + "tete.png")
	if err != nil {
		panic(err)
	}
	defer teteFile.Close()
	teteImg, _, err := image.Decode(teteFile)
	if err != nil {
		panic(err)
	}
	yeuxFile, err := os.Open(base + "yeux.png")
	if err != nil {
		panic(err)
	}
	defer yeuxFile.Close()
	yeuxImg, _, err := image.Decode(yeuxFile)
	if err != nil {
		panic(err)
	}
	for n := 0; n < b.N; n++ {
		rec := image.Rect(0, 0, 350, 400)
		finalImg := image.NewRGBA(rec)
		draw.Draw(finalImg, rec, teteImg, rec.Min, draw.Src)
		draw.Draw(finalImg, rec, boucheImg, rec.Min, draw.Over)
		draw.Draw(finalImg, rec, nezImg, rec.Min, draw.Over)
		draw.Draw(finalImg, rec, yeuxImg, rec.Min, draw.Over)
		draw.Draw(finalImg, rec, cheveuxImg, rec.Min, draw.Over)
		out, err := os.Create(fmt.Sprintf("%s../out/%d.png", base, n))
		if err != nil {
			panic(err)
		}
		defer out.Close()
		if err := png.Encode(out, finalImg); err != nil {
			panic(err)
		}
	}
}

func BenchmarkDrawParallel(b *testing.B) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	base := filepath.Dir(ex) + "/assets/"
	boucheFile, err := os.Open(base + "bouche.png")
	if err != nil {
		panic(err)
	}
	defer boucheFile.Close()
	boucheImg, _, err := image.Decode(boucheFile)
	if err != nil {
		panic(err)
	}
	cheveuxFile, err := os.Open(base + "cheveux.png")
	if err != nil {
		panic(err)
	}
	defer cheveuxFile.Close()
	cheveuxImg, _, err := image.Decode(cheveuxFile)
	if err != nil {
		panic(err)
	}
	nezFile, err := os.Open(base + "nez.png")
	if err != nil {
		panic(err)
	}
	defer nezFile.Close()
	nezImg, _, err := image.Decode(nezFile)
	if err != nil {
		panic(err)
	}
	teteFile, err := os.Open(base + "tete.png")
	if err != nil {
		panic(err)
	}
	defer teteFile.Close()
	teteImg, _, err := image.Decode(teteFile)
	if err != nil {
		panic(err)
	}
	yeuxFile, err := os.Open(base + "yeux.png")
	if err != nil {
		panic(err)
	}
	defer yeuxFile.Close()
	yeuxImg, _, err := image.Decode(yeuxFile)
	if err != nil {
		panic(err)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rec := image.Rect(0, 0, 350, 400)
			finalImg := image.NewRGBA(rec)
			draw.Draw(finalImg, rec, teteImg, rec.Min, draw.Src)
			draw.Draw(finalImg, rec, boucheImg, rec.Min, draw.Over)
			draw.Draw(finalImg, rec, nezImg, rec.Min, draw.Over)
			draw.Draw(finalImg, rec, yeuxImg, rec.Min, draw.Over)
			draw.Draw(finalImg, rec, cheveuxImg, rec.Min, draw.Over)
			out, err := os.Create(fmt.Sprintf("%s../out/%d.png", base, rand.Int()))
			if err != nil {
				panic(err)
			}
			defer out.Close()
			if err := png.Encode(out, finalImg); err != nil {
				panic(err)
			}
		}
	})
}
