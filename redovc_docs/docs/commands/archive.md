# Archives a todo

## Usage:

`redovc archive [id] [flags]`

`redovc archive [command]`

## Aliases
  archive, ar

## Examples

#### To archive a todo with id 33

`redovc archive 33`

`redovc ar 33`


#### To unarchive todo with id 33

`redovc unarchive 33`

`redovc uar 33`


#### To archive all completed todos

`redovc archive completed`

`redovc ar c`

#### Garbage collection will delete all archived todos, reclaming ids

`redovc archive gc`

`redovc ar gc`

` redovc archive gc`

`redovc ar gc`
	  Run garbage collection. Delete all archived todos and reclaim ids

## Available Commands:
```  
c           Archives all completed todos.
gc          Deletes all archived todos.
```

## Flags:
`-h, --help   help for archive`
