![redovc_logo](https://github.com/sottey/redo.vc/assets/8494620/cc081562-3b4f-48f2-88bb-7133676919da)

Redo.vc is a tool for command line fans that allows you to track your tasks. It is a full featured todo manager with tagging, projects, recurring tasks and much more, all stored in a JSON file so it is super portable and tooling new apps for the data format is super easy.

## Getting Started
To learn about available redo.vc commands, see the [usage page](https://redo.vc/usage/)

## Features
* Ability to bulkedit todo data (opens default editor with source json data)
* Ability to serve todo data as json, html or CSV via http
* [Search capabilities](https://redo.vc/commands/list/) - Search for strings in all of your todos when using the list command
* [Theming capabilities](https://redo.vc/theming) - Change all the colors and column order on the fly!
* Recurrence daily, weekly and more!
* Adding additional due strings (End of quarter, End of Month)
* Tasks can be assigned to a project
* Projects can be created dynamically while creating todos
* Tasks can have a due date or NO due date
* Tags can be specified by using the '#' prepended to single word. If it does not exist, it will be created
* Projects can be specified by using the '+' prepended to a single word in the description. If the project does not exist, it will be created
* Tasks can be prioritized
* Notes can be added to tasks and shown or hidden when listing tasks
* Due dates can be absolute (oct25, 25oct) or relative (monday, today, none, etc.)
* Tasks can be marked as completed
* Tasks can be archived
* Tasks can be deleted
* Tasks can be sorted, segmented, grouped and shown ur hidden using the extensive command line options
* Tasks are color coded with specific colors for due dates, ID's, tags, projects and more
* Groupings comtain the number of tasks
* Abbreviations for most commands to save you keystrokes
* Auto completion for various shells can be generated
* A task status can be specified (doing, waiting, evening, etc)

## Downloading
Find the installation for your device in the releases section of the GitHub project [here](https://github.com/sottey/redo.vc/releases)

## Future Plans
* [DONE] Search capabilities
* [DONE] Theming capabilities
* [DONE] Adding additional due strings (End of quarter, End of Month)
* Syncing via cloud providers (Google Drive, iCloud, Dropbox, etc.)
* [continuing] Web UI
* Importing from common formats
* Config file defaults
* Whatever you would like to see! [Let me know](https://github.com/sottey/redo.vc/discussions)


## Building

Redo.vc is available for most configurations. Go to the relase page, download the proper archive for your device. 

Once the file has downloaded, extract the binary and put it somewhere that is accessible from your terminal.

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

* [ultralist](https://github.com/gammons/ultralist) - HUGE thanks to everyone who created [ultralist](https://github.com/gammons/ultralist), an amazing todo tool!


## License

Redo.vc uses the MIT License. Please see the [redo.vc license](https://github.com/sottey/redo.vc/blob/main/LICENSE) for details
