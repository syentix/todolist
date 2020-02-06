# Easy TodoList-App by syentix

Little side project, I worked on, when I was bored. Works for CLI for now, will be availible on Telegram though. (**_Personal USE only, coz it needs a running db on localhost_**). A json variant is planned though.

## **V1 (CLI)**

By entering `./todolist -c` in your terminal, you will open the App in CLI-mode.
The App will then show all your added todos in a small little view. You can now type commands.

### Supported commands

* `add [Todo-Text]` This will add a Todo to the database.
* `check [ID]` This will put a checkmark next to the Todo corresponding to [ID].
* `delete [ID]` Deletes the Todo corresponding to [ID].
* `print` Prints the Todo-List.
* `exit` Exits the program.

## **V2 (Telegram-BOT)**

By entering `./todolist -tele` in your terminal, you will start the Telegram Bot on your local device.
