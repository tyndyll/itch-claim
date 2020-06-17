# itch-claim

With so many items included in the [itch.io Bundle for Racial Justice and Equality](https://itch.io/b/520/bundle-for-racial-justice-and-equality)
and having to click to claim each one, I threw together a quick app that will go and claim them automatically. To claim 
everything in that particular bundle will take about 5 hours (1700+ items), as there is a 5 second delay between 
requests to be kind to the itch servers.

This process will only work if you are authenticating with an itch.io username and password (i.e. not the Github sign on) and does not have the option to provide 2 factor authentication/One Time Password fields. 

## Install
To build from source ensure you have a recent version of Go (1.12+) installed. Clone the repository and run 

```
go install
```

Alternatively download the binaries

* [Linux](https://github.com/tyndyll/itch-claim/releases/download/v1.0/itch-claim)
* [macOS](https://github.com/tyndyll/itch-claim/releases/download/v1.0/itch-claim.macOS)
* [Linux](https://github.com/tyndyll/itch-claim/releases/download/v1.0/itch-claim.exe)


## Usage
```
➜  itch-claim
Add all items from the #BLM itch.io bundle to your account

Usage:
  itch-claim [flags]

Flags:
  -h, --help              help for itch-claim
      --password string   itch.io password
      --url string        your unique bundle url
      --username string   itch.io username

➜  itch-claim --username tyndyll --password mypass --url https://itch.io/bundle/download/Your_Individual_Code
``` 
