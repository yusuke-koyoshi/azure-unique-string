# azure-unique-string

Go implementation command for Azure Resource Manager template function [uniqueString](https://learn.microsoft.com/ja-jp/azure/azure-resource-manager/templates/template-functions-string#uniquestring).
Creates a hash string based on the values provided as parameters.

Bicep cannot be dry run, and you cannot check the hash string until you run it. You can check what string will be generated in advance.

## Installation

```bash
go install github.com/yusuke-koyoshi/azure-unique-string@latest
```

Installed in `$GOPATH/bin`.
Add the path to `$GOPATH/bin`.

## Synopsis

```bash
azure-unique-string string1 [string2]...
```

## Execution example

```bash
$ azure-unique-string "/subscriptions/xxxx-xxxx/resourceGroups/yourgroups" "test"
ffirqaf2lsvui
```