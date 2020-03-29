# digestream
[![CircleCI](https://circleci.com/gh/nwtgck/digestream.svg?style=shield)](https://circleci.com/gh/nwtgck/digestream)

Streaming message digest calculator: Insert digest calculation inside pipe

## Installation

Get executable binaries from [GitHub Releases](https://github.com/nwtgck/digestream/releases)

## Usage

Here is an example to use `digestream`. The command below allows you to download `ubuntu.iso` and calculate message digest at the same time.

```bash
curl http://releases.ubuntu.com/18.04/ubuntu-18.04.1-desktop-amd64.iso | digestream > ubuntu.iso
```

Then you will have a SHA-256 digest as follows in tty.
```
5748706937539418ee5707bd538c4f5eabae485d17aa49fb13ce2c9b70532433
```

Note that the message digest, `5748706937...433` is output to tty not to stdout.  
The command above is the same as `curl ... > ubuntu.iso` except for digest calculation.

## Motivation

One of the motivations of `digestream` is for [Piping Server](https://github.com/nwtgck/piping-server), which is a server to allow users to transfer data via HTTP.  
You can get more information in <https://github.com/nwtgck/piping-server>.

By using `digestream`, users get more confident to send data correctly.

```bash
# Sender
seq 1000 | digestream | curl -T - https://piping.glitch.me/myseq
```

```bash
# Receiver
curl https://piping.glitch.me/myseq | digestream > myseq.txt
```

## Algorithms

Here is examples to use `md5`, `sham256` and `sha512`. In current implementation, `sham256` is the default algorithm. But the default algorithm may be changed in the future by technology improvement.

```
echo "hello, world" | digestream md5
```

```
echo "hello, world" | digestream sha256
```

```
echo "hello, world" | digestream sha512
```

### Available algorithms

You can get all algorithms by `digestream -h`.

```
Dige(st + St)ream

Usage:
  digestream [flags]
  digestream [command]

Available Commands:
  help        Help about any command
  md5         MD5
  sha1        SHA-1
  sha256      SHA-256
  sha512      SHA-512

Flags:
  -h, --help   help for digestream

Use "digestream [command] --help" for more information about a command.
```
