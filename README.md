# Images To PDF

I needed to put all screenshots from French workbook to a one PDF file, so someone could learn from it.

I tried a lot of online tools, but nothing worked properly or was paid.

So I made this quick Go script to take images from `./images` directory and generate a .pdf file containing an image in each page.

Supports those image formats:
- JPEG
- PNG
- GIF

## How to use

Just put all your images to `./images` directory and run:

```sh
go run main.go
```

Find result PDF in `./output` directory

## Support

If you like what I share, I'd be happy if you bought me a coffee.
[!["Buy Me A Coffee"](https://cdn.buymeacoffee.com/buttons/v2/arial-orange.png)](https://buymeacoffee.com/dimot9)
