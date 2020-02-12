# CSV2LocalizableStrings 🇺🇸

## About this tool

CSVファイルから各言語の`Localizable.strings`ファイルを作成します

## Install

```
$ brew install go
```
`.zshrc.`に以下を記述

```
export GOPATH=$HOME"/go"
export PATH="$GOPATH/bin:$PATH"
```
このツールでつかっているライブラリのインストール
```
go get github.com/jessevdk/go-flags
```

## Build 

```
$ go build -o ./bin
```

## Usage

`CSVのフォーマットはCSV UTF-8(コンマ区切り)`

カレントディレクトリに各言語の`Localizable.strings`を作成
```
$ csv2LocalizableStrings -f <csvファイル>
```

指定したディレクトリに各言語の`Localizable.strings`を作成

```
$ csv2LocalizableStrings -f <csvファイル> -o <出力先ディレクトリ>
```

## Options
```
Application Options:
  -f, --path=INPUT_FILE_PATH            LocalizeStringsファイルへ変換するCSVのパス
  -o, --output=OUTPUT_DIRECTORY_PATH    ファイルの出力先 (default: ./)

Help Options:
  -h, --help                            Show this help message
```

## Issue 😫

- [ ] 日英以外は未対応
- [ ] 前回との差分比較を行えるようにする
- [ ] Base.lprojにコピーする