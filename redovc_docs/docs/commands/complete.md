# complete command
Items that are done can be marked as completed. They will still appear, but will have a checkmark. To hide todos, you can archive them. Items can also be "uncompleted" to mark them as not done


## Usage
`redovc complete [id] [flags]`

## Aliases
  complete, c

## Examples
#### Complete a todo with id 33

`redovc complete 33`
    
`redovc c 33`

#### Complete a todo with id 33 and archive it
    
`redovc uncomplete 33 --archive`

#### Uncomplete todo with id 33
    
`redovc uncomplete 33`
    
`redovc uc 33`

## Flags
`--archive   Archives a completed todo automatically`


`-h, --help      help for complete`