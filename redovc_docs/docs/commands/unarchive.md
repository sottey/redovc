# unarchive command
Un-archives a todo

## Usage

`redovc unarchive [id] [flags]`

## Aliases
  unarchive, uar

## Examples

#### To archive a todo with id 33:

`redovc archive 33`

`redovc ar 33`

#### To unarchive todo with id 33:

`redovc unarchive 33`

`redovc uar 33`

#### To archive all completed todos:

`redovc archive completed`

`redovc ar c`

#### Garbage collection will delete all archived todos, reclaming ids:

`redovc archive gc`

`redovc ar gc`

#### redovc archive gc

`redovc ar gc`
Run garbage collection. Delete all archived todos and reclaim ids


## Flags

`-h, --help   help for unarchive`