# extract-hostname
A utility to extract hostnames from stdin input, written in Golang

### Sample stdin

```
> cat sample_stdin.txt
"propertyName": "staging.mbe.domain.com",
                            "label": "mobile-domain-com",
                                "domain.com"
                                "domain.co.uk"
                                "domain.ai"
                                    "hostname": "billing-api-staging.domain.com",
                                            "subjectCN": "mbe-staging.domain.com",
                        "description": "mbe-staging.domain.com",
                        "name": "mbe-staging.domain.com"
```
### Usage
```
> cat sample_stdin.txt | extract-hostname
staging.mbe.domain.com
domain.com
domain.co.uk
domain.ai
billing-api-staging.domain.com
mbe-staging.domain.com
mbe-staging.domain.com
```

### Install with Homebrew

First install gomod package.
```
brew install filosottile/gomod/brew-gomod
brew gomod github.com/icheko/extract-hostname
```
