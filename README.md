# slalom-badge-pi

A project to create a pi-based badge that slalom folk can wear and enjoy.

## Hardware requirements

- Raspberry Pi Zero W v1.1
- [Waveshare e-ink display HAT](https://www.pishop.ca/product/250x122-2-13inch-e-ink-display-hat-for-raspberry-pi/)
- [16GB SD Card](https://www.amazon.ca/Sandisk-Ultra-Micro-UHS-I-Adapter/dp/B073K14CVB/ref=sr_1_9?dchild=1&keywords=16gb+sd+card&qid=1634089895&sr=8-9)
- (Optional) [PiSugar battery](https://www.pisugar.com/) 


## Building an image

You'll need some disk space. Probably about 25GB. You'll also need a linux/macOS environment.

There's a `makefile` in the project root, to aid with installing dependencies, building an image, and removing unused files.

Then...
```bash
make install # install dependencies
make image # create an image that can be flashed to a pi
```

```bash
make clean # if you'd like some disk space back :)
```