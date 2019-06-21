# stripes

<div style="width: 250px; height: 250px; background: url(&quot;bluestripes-example.png&quot;) repeat"></div>

`stripes` generates an image containing two horizontal stripes, suitable for use as a tiled background with a striped pattern.

It outputs the image in PNG format to standard output.

For example, to generate the pattern shown above, run the following:

`stripes -a '#ddddff' -b '#d0d0ff' -w 60 > bluestripes.png`

(Bear in mind that the stripes will appear narrower when displayed at native resolution on higher-DPI screens.)