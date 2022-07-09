<!-- the .webp image must be in a public repository to be properly loaded-->
<div align="center">
    <img width="776" height="376" src="https://raw.githubusercontent.com/johannsuarez/johannsuarez/main/cover.webp">
    <h1 align="center">bitcoin mining simulator</h1>
    <img src="https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square"  alt="standard-readme compliant">

</div><br/>

<div>
  <font>a testing suite as part of development operations to test
      the quality of software we write for bitcoin miners.</font>
</div>

## table of contents


- [install](#install)
- [usage](#usage)
- [maintainers](#maintainers)
- [contributing](#contributing)
- [license](#license)


## install

1. clone this repository.
2. cd into the repository root and run `go install`. the binary will be installed in the /bin directory of your gopath ( check with ```go env``` )
3. the application requires two environment variables, $**block_template_dir** and $**miner_template_dir**. $**block_template_dir must** contain the path to the .json block templates to be used by the bitcoin daemon simulator. 
4. here's are examples on how you set both variables with fish and bash.

for bash: 
```
export block_template_dir="/home/user/gohash-miner-simulator/data/bitcoind/block_templates"

export miner_template_dir="/home/user/gohash-miner-simulator/data/mining"
```

for fish:
```
set -ux block_template_dir "/home/user/gohash-miner-simulator/data/bitcoind/block_templates"

set -ux miner_template_dir "/home/user/gohash-miner-simulator/data/mining"
```


## usage

### ⛏️ mining simulation
to create several instances of simulated asic miners, provide
the program with the "minerclient" argument.

```
gohash_miner_simulator minerclient 
```

you can tweak the simulation parameters using several flags as enumerated here.


```
flags:
      --all                 will load all miner types (default true)
  -h, --help                help for minerclient
  -l, --load uint32         the number of miners you want to connect, defaults to 10. (default 10)
  -m, --miner-type string   name of the miner to target (s9, s17, s19, m30) (default "all")
  -p, --port uint16         corresponding port, defaults to 60000 (default 60000)
  -s, --server string       stratum url or public ip address, defaults to localhost. (default "127.0.0.1")
```
shown above are the short form of each flag, the long form, the data type expected, and the description.

[ under construction ]

### 💻 bitcoin daemon simulator

simulates getblocktemplate. a local http server will start that dispenses
block templates when the endpoint receives the appropriate post request.

possible flags and arguments can be viewed using `gohash_miner_simulator bitcoind -h`.

```
flags:
  -a, --auth string   authentication in the form: (name:pass)
  -h, --help          help for bitcoind
  -p, --port uint16   port number (default 4000)
```


run the bitcoind simulator by providing a name and password pair and a port number.
```
gohash_miner_simulator bitcoind -a "edgar_tool:pass123" -p 4000
```

<div align="center">
    <img src="https://raw.githubusercontent.com/johann-gohash/repository_media/main/sim/bitcoind_demo.svg"/>
</div><br/>
          

with the above credentials, we should successfully get a block template with this curl command.
```
curl --user edgar_tool:pass123 --data-binary '{"method":"getblocktemplate","params": [{"rules": ["segwit"]}],"id":1}' -h 'content-type: text/plain;' http://127.0.0.1:4000/
```

## maintainers

[@johann-suarez](https://github.com/johann-gohash)

## contributing

feel free to dive in! [open an issue](https://github.com/gohash-software/gohash-miner-simulator/issues/new) or submit prs.

standard readme follows the [contributor covenant](http://contributor-covenant.org/version/1/3/0/) code of conduct.

<!-- 
note: 
you need to set up an account first on opencollective.com to dynamically generate an image of all contributors

### contributors

this project exists thanks to all the people who contribute. 
<a href="https://github.com/gohash-software/gohash-miner-simulator/graphs/contributors"><img src="https://opencollective.com/standard-readme/contributors.svg?width=890&button=false" /></a>
-->

## license

[mit](license) © gohash team
