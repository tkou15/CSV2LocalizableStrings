# CSV2LocalizableStrings 🇺🇸

## About this tool

CSVファイルから各言語の`Localizable.strings`ファイルを作成する

## Build 

```
$ go build csv2LocalizableStrings.go
```

## Usage

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

- [ ] 日英以外は仮対応
- [ ] 前回との差分比較を行えるようにする
- [ ] Base.lprojにen.lproj/Localizable.stringsをコピーする
