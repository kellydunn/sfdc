# sfdc

A collection of General Purpose Salesforce utilities that I find personally useful when debugging.

## installation

This library is written in golang, and as such, uses the go tool chain to download and install itself.

```bash
$ go install github.com/kellydunn/sfdc
```

## usage

Currently, this library only has one command, `15to18` which converts a 15 character Salesforce ID to an 18 character Salesforce ID.  More on the difference [here](https://help.salesforce.com/apex/HTViewSolution?urlname=How-do-unique-IDs-work-in-Salesforce-1327108651310&). The alogrithm that creates these ids is based on 

```bash
$ sfdc 15to18 <salesforce_id>
```

