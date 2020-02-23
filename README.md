Recommended editor is [Visual Studio Code](https://code.visualstudio.com/) with following extensions:

 * [Go](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
 * [EditorConfig](https://marketplace.visualstudio.com/items?itemName=EditorConfig.EditorConfig)

To avoid breaking test you have to set some specific editor's settings for `*.expected` files:

 * eof `\n`
 * insert final newline is on
 * trim trailing whitespace is off

If your editor supports EditorConfig, you shouldn't need to worry about this.

Some useful tasks are defined for VSCode. To run them press `Ctrl + Shift + P` and select `Tasks: Run Task`.
