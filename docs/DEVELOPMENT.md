# Development

This extension was created by running

```shell
gh extension create --precompiled=go itkdev
```

(cf. [Creating GitHub CLI extensions](https://docs.github.com/en/github-cli/github-cli/creating-github-cli-extensions)).

## Commands

We use [Cobra](https://github.com/spf13/cobra) for commands. To add a new command, install [Cobra
Generator](https://github.com/spf13/cobra-cli) and run

```shell
cobra-cli add the-new-command
```
