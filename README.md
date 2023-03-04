# CLI tool str

CLI tool `str` performs common string operations in an easy way.

_Example:_

```sh
# returns "Doe"
echo "John Doe" | str field -i1
```

_`str` works with standard input/output, so you can easily perform multiple operations in a chain:_

```sh
# returns a random string with 10 characters
echo "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" | str shuffle | str sub -l10
```

# Installation

Currently the app has not been added to any package manager.

If you want to use the app, clone the repository and run `make install`.

> The `make install` command may not work on windows systems.
> In this case you can also run make build and copy the binary in dist directory manually.

# Contribution

Since I do not know or can guess all use cases, it is difficult to implement all commands.
If you have a use case or a good idea for a command, feel free to suggest it as an issue.

Of course, pull requests are always welcome.
