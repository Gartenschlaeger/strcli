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
echo "abcdefghijklmnopqrstuvwxyz0123456789" | str shuffle | str sub -l10
```

Regular expressions are also supported:

```sh
# returns 2023
echo "2023-10-11" | str regex -p"(\d{4})-(\d{2})-(\d{2})" -g1
```

```sh
# returns He__o wor_d!
echo "Hello world\!" | str regex -p"l" -r"_"
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

> test workflow
