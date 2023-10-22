# redo.vc (formerly redovc)
Built on [ultralist](https://github.com/ultralist)

Redo.vc is a tool for command line fans that allows you to track your tasks. It is a full featured todo manager with tagging, projects, recurring tasks and much more, all stored in a JSON file so it is super portable and tooling new apps for the data format is super easy.


## Features

* Tasks can be assigned to a project.
* Projects can be created dynamically while creating todos
* Tasks can have a due date or NO due date
* Tags can be specified by using the '#' prepended to single word. If it does not exist, it will be created.
* Projects can be specified by using the '+' prepended to a single word in the description. If the project does not exist, it will be created
* Tasks can be prioritized
* Notes can be added to tasks and shown or hidden when listing tasks
* Due dates can be absolute (oct25) or relative (monday, today, etc.)
* Tasks can be marked as completed
* Tasks can be archived
* Tasks can be deleted
* Tasks can be sorted, segmented, grouped and shown ur hidden using the extensive command line options
* Tasks are color coded with specific colors for due dates, ID's, tags, projects and more
* Groupings comtain the number of tasks
* Abbreviations for most commands to save you keystrokes
* Auto completion for various shells can be generated
* A task status can be specified (doing, waiting, evening, etc)


## Future Plans
* Theming
* Syncing via cloud providers (Google Drive, iCloud, Dropbox, etc.)
* Web UI
* Importing from common formats
* Config file defaults
* Whatever you would like to see!


## Builds

Redo.vc is available for most configurations. Go to the relase page, download the proper archive for your device. 

Once the file has downloaded, extract the binary and put it somewhere that is accessible from your terminal.

To run the app, just use the command redovc (some releases use the previous name redovc). To learn more about what you can do with the app, use the following command:
```
redovc --help
```

To learn about a specific command, run redovc with the command and then --help:
```
redovc [command] --help
```

## Contribute

Clone locally and run
```
go build main.go
```

OR (using rake)
```
rake build
```

## Shoutouts

* [ultralist](https://github.com/ultralist) - HUGE thanks to everyone who created [ultralist](https://github.com/ultralist), an amazing todo tool!


## License

Redo.vc uses the MIT License. Please see the [redo.vc license](https://github.com/sottey/redo.vc/blob/main/LICENSE) for details