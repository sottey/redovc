# status command
Sets the status of a todo item

A status should be a single lower-case word, e.g. "now", "blocked", or "waiting", "tonight", etc..

## Usage

`redovc status [id] [status] [flags]`

## Aliases
  status, s

## Examples

#### To add a "blocked" status to a todo:

`redovc status 33 blocked`

`redovc s 33 blocked`

#### You can remove a status by setting a status to "none".  Example:

`redovc s 33 none`


## Flags

`-h, --help   help for status`