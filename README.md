# Overview

Package `flagicon` provides a function to generate flag icons from image files

By default the flag icon will be of size 14x14 pixels; the original image will
be resized to this size, with a 2px border with color `#808080`.

If the image is not square, `flagicon` first crops the image (centered at the center
of the image) such that it is square before it resizes the image.

# Usage
## Overview
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

## General Usage

Let us image that we wish to create a flagicon for the state flag of New Mexico, which we have saved as `nm.png`. To generate a basic flagicon for the state flag, we can run:

```
flagicon nm.png
```

This would run `flagicon` and create ![New Mexico flagicon](img/nm.png "New Mexico Flagicon"), which would be saved as `out.png`.

## Saving to Specific Files

Let us take the same example from above. Here, we wish to create a flagicon for the state flag of New Mexico, but we wish to name the generated flagicon image as `nm_icon.png`. To do this, we can run

```
flagicon nm.png -o nm_icon.png
```

This would run `flagicon` and create ![New Mexico flagicon](img/nm.png "New Mexico Flagicon"), which would be saved as `nm_icon.png`.

## Controlling

When dealing with square images, `flagicon` simply creates the circular flagicons from the center of the image. But most flags, with the notable exception of the ![Vatican flagicon](img/vt.png "Vatican Flagicon") Vatican City and ![Switzerland flagicon](img/ch.png "Switzerland Flagicon") Switzerland, are not square. How can we control which parts of the flag to include in the flagicon?

When dealing with non-square images, such as the state flag of New Mexico from the previous examples, we see that by default, `flagicon` uses the center of the image to generate the flagicon. However, many flags, such as the state flag of Ohio, have their primary design on the left-hand side of the flag, close to the flag pole. 

If we use the default behavior of `flagicon` to generate a flagicon for the state flag of Ohio, we get the following flagicon: ![Ohio flagicon](img/oh-center.png "Centered Ohio Flagicon").

```
flagicon oh.png -o oh_icon.png
```

However, we see that a large portion of the "primary details" in the Ohio state flag are not present in the centered flagicon. When cropping the image to generate the circular flagicon, we want `flagicon` to use the left side of the image rather than the center of the image.

To do this, we can use the `-l` flag, which tells the program to use the left side of the image for the flagicon. Using this option produces the following flagicon: ![Ohio flagicon](img/oh-left.png "Left-aligned Ohio Flagicon").

```
flagicon -l oh.png -o oh_icon_left.png
```