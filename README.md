# slalom-badge-pi

A project to create a pi-based badge that Slalom folk can wear, experiment with and enjoy. Based on Kali linux.

## Hardware requirements

- Raspberry Pi Zero W v1.1
- [Waveshare e-ink display HAT](https://www.pishop.ca/product/250x122-2-13inch-e-ink-display-hat-for-raspberry-pi/)
- [16GB SD Card](https://www.amazon.ca/Sandisk-Ultra-Micro-UHS-I-Adapter/dp/B073K14CVB/ref=sr_1_9?dchild=1&keywords=16gb+sd+card&qid=1634089895&sr=8-9)
- (Optional) [PiSugar battery](https://www.pisugar.com/) 


## Building an image

You'll need some disk space. Probably about 25GB. You'll also need a linux/macOS environment.

There's a `makefile` in the project root, to aid with installing dependencies, building an image, and removing unused files.

First, install the dependencies...
```bash
make install
```
Then, make an image that can be flashed to an SD card...
```bash
make image
```

Once you're done, and have flashed the image to your pi, you may want some disk space back...

```bash
make clean
```

## SSH to the pi

Right now, you must connect your pi to your wifi to access it. In the future, we'll enable gadget mode (plug it in your computers USB port to connect to it). Once it's connected to your wifi, find its IP address, and `ssh kali@<ip address>` (**the password is kali**).
