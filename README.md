# blinky
cli for blinkstick

## Background
There are a lot of tutorials for using Blinkstick with Python. However, all the tutorials I found were based on Python 2.7 and I could never get them working. This project represents my attempt to create a Blinkstick CLI using Go.

Anyway, I browsed the blinkstick forums, which are pretty dead, and found this page that shows the [official and unofficial APIs](https://www.blinkstick.com/help/api-implementations). There are a couple of prerequisites. I'm running Raspberry OS - Bullseye on an RPI4.

## Goal
My goal is to have a Go-based command line utility that I can use to interact with the devices. I'm still forming exactly what I want this utility to do but a simple device recognition and blink will suffice for now. I think I may have the utility called from various monitoring / notification scripts that I build. Let's see where we end up.

## Prerequisites for RPI4
You will need to install libusb
`sudo apt-get install libusb-1.0-0-dev`

You will need to give all users access to usb devices or else you'll get access denied errors
`echo "SUBSYSTEM==\"usb\", ATTR{idVendor}==\"20a0\", ATTR{idProduct}==\"41e5\", MODE:=\"0666\"" | sudo tee /etc/udev/rules.d/85-blinkstick.rules`

`sudo reboot`

## Test
Once you've performed the above steps, let's test it out. Go into the `cli` folder and run `go run main.go`
You should get any detected devices light up red, then green, and finally blue for 1 second each before turning off after blue.

Now we can start building our utility.

The test code was sampled from the maintainer's repo, linked below.

## References
[GoDoc Reference](https://pkg.go.dev/github.com/boombuler/led?utm_source=godoc)
[Maintainer's Repo](https://github.com/boombuler/led)