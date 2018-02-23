# Overview

Package `flagicon` provides a function to generate flag icons from image files

By default the flag icon will be of size 14x14 pixels; the original image will
be resized to this size, with a 2px border with color `#808080`.

If the image is not square, `flagicon` first crops the image (centered at the center
of the image) such that it is square before it resizes the image.

# Usage
```
Usage: flagicon [--output OUTPUT] [--left] INPUT

Positional arguments:
  INPUT                  Path to input image file

Options:
  --output OUTPUT, -o OUTPUT
                         Output file name [default: out.png]
  --left, -l             Generate flagicon from left-side of image, rather than the center
  --help, -h             display this help and exit
```

For example, if we were to create a flagicon for the file `nm.png` (the state flag of New Mexico), we might run the following:

```
flagicon nm.png
```

This would produce a flagicon and save it to `out.png`. 

If we wish to name the output file -- perhaps `nm_icon.png`, we might run:

```
flagicon nm.png -o nm_icon.png
```