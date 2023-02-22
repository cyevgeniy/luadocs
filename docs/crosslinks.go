package main

import ("fmt"
        "bufio"
        "os"
        "strings"
        "path/filepath"
        "io/ioutil"
)

var base map[string]string

// "2.5.3",  "Garbage Collection"
// s
type LinkInfo struct {
    oldId string // like 2.5.3
    newId string // like garbage-collection
    name string // like Garbage collection
    url string // like /02_basic_concepts/ch03#garbage-collection
}

var info []*LinkInfo

func init() {
    base = map[string]string{
        "1": "/01_intro",
        "2": "/02_basic_concepts",
        "3": "/03_the_language",
        "4": "/04_API",
        "5": "/05_aux_lib",
        "6": "/06_standard_lib",
        "7": "/07_standalone",
        "8": "/08_incompatibilities",
        "9": "/09_complete_syntax",
    }

    var initErr error
    info, initErr = readInput()
    if initErr != nil {
        panic(initErr)
    }
}

func Url(oldId string, newId string) string {
	// oldId = 2.5.1, newId = "garbage-collection" ===> /02_basic_concepts/ch05#garbage-collection
	// oldId=3, newId="the-language" ===> /03_the_language/intro#the-language
	// oldId = 5.6, newId="aux-lib-xxx" ====> /05_aux_lib/ch06#aux-lib-xxx

	section := string(oldId[0])

	// when we have single number
	if len(oldId) == 1 {
		return base[section] + "/intro#" + newId
	}

	// Otherwise we always have at least 3 length string
	return fmt.Sprintf("%s/ch%02s#%s", base[section], string(oldId[2]), newId)
}

func NewLinkInfo(str string) (*LinkInfo, error) {
    oldId, name, found := strings.Cut(str, " ")

    if !found {
        return nil, fmt.Errorf("Provided string (%s) can't be used as source for LinkInfo", str)
    }

    newId := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
    url := Url(oldId, newId)

    return &LinkInfo{
        oldId,
        newId,
        name,
        url,
    }, nil
}

func readInput() ([]*LinkInfo, error) {
    res := make([]*LinkInfo, 0)

    fname := "input.txt"

    f, err := os.Open(fname)

    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        info, err := NewLinkInfo(scanner.Text())
        if err != nil {
            return nil, err
        }

        res = append(res, info)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
	}

    return res, nil
}

// https://gist.github.com/tdegrunt/045f6b3377f3f7ffa408
func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	matched, err := filepath.Match("*.md", fi.Name())

	if err != nil {
		panic(err)
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		fmt.Println(path)

        newContents := string(read)
        for i := range info {
            old := fmt.Sprintf("(#%s)", info[i].oldId)
            new := fmt.Sprintf("(%s)", info[i].url)
            newContents = strings.Replace(string(newContents), old,  new, -1)

            old = fmt.Sprintf("[§%s]", info[i].oldId)
            new = fmt.Sprintf("[%s]", info[i].name)
            newContents = strings.Replace(newContents, old, new, -1)
        }

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}


func main() {
    err := filepath.Walk(".", visit)
	if err != nil {
		panic(err)
	}
}
