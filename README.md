# God's Eye Trojan

## Introduction
This project provides two main functions which are showing `username` and `password` for Web Authorization and establishing a TCP connection with the controller (the Server).

## Developing Environment
- Golang Version 1.9

## Running Environment
- The Payload only works on MacOS

## Workflow
1. The developer needs to compile this project based on the Golang developing environment.
``` go build trojan.go ```
2. When an user downloaded the compiled software and run, `username` and `password` will be shown.
3. The TCP connnection between a victim's computer and the controller will be established after showing the `password`.

## Thanks
- Thanks to the project of [EggShell](https://github.com/neoneggplant/EggShell)
- Thanks to the project of fyne
