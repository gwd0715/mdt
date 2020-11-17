# MDT -- Markdown To

Based on [Cobra](https://github.com/spf13/cobra) and [OpenCC](https://github.com/sgoby/opencc) tools, MDT is a CLI tool which converts markdown file between Chinese zh-cn and zh-tw.

基于[Cobra](https://github.com/spf13/cobra) 和[OpenCC](https://github.com/sgoby/opencc)两个Go的工具，MDT是一个将markdown文件进行中文简繁体转换的CLI工具。

## Usage 用法

Download binary source code and set your env $PATH.
```
Usage: mdt [command]
Avaialable Commands:  
  cn          Convert markdown file from zh-tw to zh-cn  
  help        Help about any command  
  tw          convert markdown file from zh-ch to zh-tw  

Flags:
  -h, --help   help for mdt
```