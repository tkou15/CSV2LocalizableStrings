package main

import (
	"fmt"
    "encoding/csv"
    "io"
    "os"
    "errors"
    "path/filepath"
    "strings"

    "github.com/jessevdk/go-flags"
)


// グロッサリの順番が変わったときは修正必要
const (
    KEY int = iota
    JP
    EN
    FR
    DE
    IT
    ES
    PT  //
    PT_BR
    NL
    DA
    NO
    SV
    FI
    EL //
    PL
    CS
    RU
    HU  //
    TR 
    AR  //
    HR  //
    ET  //
    LV  //
    LT  //
    SL  //
    UK  //
    SK  //
    RO  //
    BG  //
    HE  //
    FA  //
    SC
    TC
    KO
    TH
)
const OUTPUT_FILE_NAME = "Localizable.strings"

var opts struct {
    INPUT_FILE_PATH  	string `short:"f" long:"path" description:"LocalizeStringsファイルへ変換するCSVのパス" value-name:"INPUT_FILE_PATH" required:"true"`
	OUTPUT_DIRECTORY_PATH	    string `short:"o" long:"output" default:"./" description:"ファイルの出力先" value-name:"OUTPUT_DIRECTORY_PATH" required:"true"`
}

func main() {

    _, err := flags.Parse(&opts)
    CheckIfError(err)
    
    // mapだと出力順が変わってしまうため２次元配列としている
    localizeTable := make([][]string, 40)

    csvRead(opts.INPUT_FILE_PATH, &localizeTable)
    outputLocalizeFile(&localizeTable)

    fmt.Println("Success!")
}

// CSVを読み込む
func csvRead(filePath string, localizeTable *[][]string){

    if filepath.Ext(filePath) != ".csv"{
        CheckIfError(errors.New("Not csv!"))
    }

    fp, err := os.Open(filePath)
    CheckIfError(err)

	defer fp.Close()
	
    reader := csv.NewReader(fp)
    reader.LazyQuotes = true
	
	for {
        record, err := reader.Read()
          if err == io.EOF {
            break
        }
        CheckIfError(err)
        
        (*localizeTable)[KEY] = append((*localizeTable)[KEY], record[KEY])
        (*localizeTable)[JP] = append((*localizeTable)[JP], record[JP])
        (*localizeTable)[EN] = append((*localizeTable)[EN], record[EN])

        // 正式文言がくるまでENを使用
        (*localizeTable)[FR] = append((*localizeTable)[FR], record[EN])
        (*localizeTable)[DE] = append((*localizeTable)[DE], record[EN])
        (*localizeTable)[IT] = append((*localizeTable)[IT], record[EN])
        (*localizeTable)[ES] = append((*localizeTable)[ES], record[EN])
        (*localizeTable)[PT_BR] = append((*localizeTable)[PT_BR], record[EN])
        (*localizeTable)[NL] = append((*localizeTable)[NL], record[EN])
        (*localizeTable)[DA] = append((*localizeTable)[DA], record[EN])
        (*localizeTable)[NO] = append((*localizeTable)[NO], record[EN])
        (*localizeTable)[SV] = append((*localizeTable)[SV], record[EN])
        (*localizeTable)[FI] = append((*localizeTable)[FI], record[EN])
        (*localizeTable)[PL] = append((*localizeTable)[PL], record[EN])
        (*localizeTable)[CS] = append((*localizeTable)[CS], record[EN])
        (*localizeTable)[RU] = append((*localizeTable)[RU], record[EN])
        (*localizeTable)[TR] = append((*localizeTable)[TR], record[EN])
        (*localizeTable)[SC] = append((*localizeTable)[SC], record[EN])
        (*localizeTable)[TC] = append((*localizeTable)[TC], record[EN])
        (*localizeTable)[KO] = append((*localizeTable)[KO], record[EN])
        (*localizeTable)[TH] = append((*localizeTable)[TH], record[EN])
    }
}

func outputLocalizeFile(localizeTable *[][]string){
    for languageType, language := range *localizeTable {
        if languageType == KEY{
            continue
        } else if language != nil{
            fp := createLocalizeFile(languageType)
            writeLocarizeFile(fp, &(*localizeTable)[KEY], &language)
            defer fp.Close()
        }
    }
}

 //ファイルの作成
func createLocalizeFile(languageType int) *os.File {
    createtDirectoryList := map[int]string{
        JP: "ja.lproj",
        EN: "en.lproj",
        FR: "fr.lproj",
        DE: "de.lproj",
        IT: "it.lproj",
        ES: "es.lproj",
        PT_BR: "pt-BR.lproj",
        NL: "nl.lproj",
        DA: "da.lproj",
        NO: "nb.lproj",
        SV: "sv.lproj",
        FI: "fi.lproj",
        PL: "pl-PL.lproj",
        CS: "cs.lproj",
        RU: "ru.lproj",
        TR: "tr.lproj",
        SC: "zh-Hans.lproj",
        TC: "zh-Hant.lproj",
        KO: "ko.lproj",
        TH: "th.lproj",
    }

    output_path := strings.TrimRight(opts.OUTPUT_DIRECTORY_PATH, "/")
    outputDirectoryPath := output_path + "/" + createtDirectoryList[languageType]

    err := os.RemoveAll(outputDirectoryPath)
    CheckIfError(err)

    err = os.Mkdir(outputDirectoryPath, 0777)
    CheckIfError(err)

    fp, err := os.OpenFile(outputDirectoryPath + "/" +  OUTPUT_FILE_NAME, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    CheckIfError(err)

    return fp
}

// ファイルの書き込み
func writeLocarizeFile(fp *os.File, localizeTableKey *[]string, language *[]string){
    for index, word := range *language {
        if index != 0 {
            // 文章中に半角の"が入ってたら置き換える
            // 情メディアに直してもらった方が良い
            word = strings.Replace(word, "\"", "\\\"", -1) 
            fmt.Fprintln(fp, "\"" + (*localizeTableKey)[index] + "\" = " + "\"" + word + "\";") //書き込み
        }else{
            fmt.Fprintln(fp, "// " + (*localizeTableKey)[index] +  "=" + word)
        }
    }
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}