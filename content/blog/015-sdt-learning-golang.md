+++
title = "Learning developing \"SDT\""
date = "2022-05-22T15:56:22Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/gabriel-heinzer-xbEVM6oJ1Fs-unsplash.jpg"
tags = []
keywords = []
description = "Little CLI for smart developers"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

*Photo by [Gabriel Heinzer](https://unsplash.com/@6heinz3r?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/@6heinz3r?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)*

In the daily life of a developer it happens that some stupid tools are needed such as encoders / decoders, string generators of various kinds, etc ...
Many rely on the various online sites where you post your content and give you the elaborate answer.
From a security point of view, however, it is an operation that cannot be considered very beautiful, even if the functioning of these online tools is only on the client side (javascript).

This is why I started, somewhat as a challenge, trying to create a CLI that would generate a series of random bytes in base64, using GoLang and cobra.

As I developed this little CLI I came up with other features that I have implemented, or that I have noted at the time.

It was an excellent test bed for some aspects that I had never tried to implement or that I had never investigated.

I called the tool **SDT**: *Smart Developer Tools* and the repository is located [here](https://github.com/sandrolain/sdt).


What have I learned?
- Create a cli in golang using [cobra 🐍](https://github.com/spf13/cobra) and create related unit tests
- Manage parameters from files using [viper](https://github.com/spf13/viper)
- Input management from different sources (Stdin, file, command line arguments)
- Management of different types of output
- Use of Go 1.18 generics!
- Multi-architecture build and release using goreleaser
- Use of UPX for the compression of the generated binary files
- Split architecture optimized code (wasm)
- Creation of a web app based on WebAssembly
- Use of [Vite](https://vitejs.dev/) for the creation of web apps based on [Lit](https://lit.dev/)
- Using [Fyne](https://fyne.io/) to create a desktop GUI (work in progess)
- Deploy of hybrid application golang + typescript with Netlify (first demo [here](https://sdt.sandrolain.com/))

I hope that this tool becomes useful and does not remain just a test bed, I have already been able to use it successfully in daily work.
