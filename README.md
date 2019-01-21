# dots
dots is a configuration tool for dotfiles.  
This tool downloads files from public github repository that has *dots.yml* configuration file at root directory and copy them to local path according to setting.

## Usage
sample repository [matsune/dots_sample](https://github.com/matsune/dots_sample)
```shell
$ dots matsune/dots_sample
```
Above command will download all targets.

If you want to download specific target:
```shell
$ dots matsune/dots_sample vimrc
```

And targets can have tags and you can filter by them:
```shell
$ dots matsune/dots_sample -t linux // or --tag=linux
```


## dots.yml
```yaml
targets:
  - name: vimrc     # target name
    file: ./.vimrc  # relative file path from dots.yml
    dst: ~/.vimrc   # copy destination path
    tags:
      - a
      - b
sub:                # nested sub directories that has dots.yml
  - ./zsh
```
```
.
├── dots.yml
├── .vimrc
└── zsh
    ├── dots.yml
    └── ...
```

## Example
My dotfiles repository using dots
[matsune/dotfiles](https://github.com/matsune/dotfiles)
