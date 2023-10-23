# edit command
Edit a todo

You can edit all facets of a todo

## Usage
`redovc edit [id] [flags]`

## Aliases
  edit, e

## Examples
#### To edit a todo's subject:

`redovc edit 33 Meeting with #bob about +project`
`redovc e 33 Change the subject once again`

#### To edit just the due date, keeping the subject:

`redovc edit 33 due:mon`

#### To remove a due date:

`redovc edit 33 due:none`

#### To edit a status

`redovc edit 33 status:next`

#### To remove a status:

`redovc edit 33 status:none`

## Flags

`-h, --help   help for edit`