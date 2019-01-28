# digestream
Streaming message digest calculator: Insert digest calculation inside pipe

## Usage

Here is an example to use `digestream`.

```bash
seq 1000 | digestream | gzip > seq1000.gz
```

Then you will have a digest as follows in tty.
```
67d4ff71d43921d5739f387da09746f405e425b07d727e4c69d029461d1f051f
```

Note that the message digest, `67d4ff71d...51f` is output to tty not to stdout, so the digest must not be passed to the `gzip`.  
The command above is the same as `seq 1000 | gzip > seq1000.gz` except for digest calculation.

## Motivation

One of the motivations of `digestream` is for [Piping server](https://github.com/nwtgck/piping-server), which is a server to allow users to transfer data via HTTP.  
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
  ./digestream [flags]
  ./digestream [command]

Available Commands:
  help        Help about any command
  md5         MD5
  sha1        SHA-1
  sha256      SHA-256
  sha512      SHA-512

Flags:
  -h, --help   help for ./digestream

Use "./digestream [command] --help" for more information about a command.
```
