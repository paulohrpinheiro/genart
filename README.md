# genart
Generative Art in Go: httpd server and cli command line

Generate images (patterns) from a simple formula:

    color = ( C - C * X * Y ) mod MAX

`C` is a constante, ``(X, Y)` coordinates and MAX the number of colors.

## How to use

### cmd line

To get help:

```
$ bin/genart --help
Usage of bin/genart:
  -colors int
    	number of colors (default 216)
  -constant int
    	constant for formula (c): c-c*x*y (default 10)
  -filename string
    	filename to save image (.png) (default "formulaimg.png")
  -size int
    	size of image (default 1000)
```

Generating an image:

```
$ make genart
$ bin/genart -filename output.png -colors 100
```

### http server

To get the same image from above cmd line example:

```
make runserver
```

Then, on another terminal (or from browser):

```
$ curl -v http://127.0.0.1:8090/ccxy/size/1000/constant/10/colors/100/ --output output.png
```

### The image

![Generatedimage](images/ccxy_size:1000_constant:10_colors:100.png)
