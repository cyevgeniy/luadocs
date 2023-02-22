package main

import ("fmt"
        "bufio"
        "os"
        "strings"
        "path/filepath"
        "io/ioutil"
        "unicode/utf8"
        "regexp"
)


type LinkInfo struct {
    oldId string // like 2.5.3
    newId string // like garbage-collection
    name string // like Garbage collection
    url string // like /02_basic_concepts/ch03#garbage-collection
}

var paths map[string]string
var res []*LinkInfo

func init() {
    var originErr error
    res, originErr = loadOrigins()
    if originErr != nil {
        fmt.Println(originErr.Error())
    }

    paths, _ = loadPaths()

    for i := range res {
        res[i].url = paths[res[i].newId] + "#" + res[i].newId
    }

}

func NewLinkInfo(s string) (*LinkInfo, error) {
    splitIdx := strings.LastIndex(s, "]") 

    if splitIdx == -1 {
        return nil, fmt.Errorf("Can't split source string %s", s)

    }

    // [lua_absimplementation] some_implementation
    fdecl := s[1:splitIdx]
    fid := s[splitIdx + 2:]


    return &LinkInfo{
        oldId: fid,
        newId:getAPIAnchor(fdecl),
    }, nil
}


func loadPaths() (map[string]string, error){
    res := make(map[string]string)

    origin := "input_api.txt"

    f, err := os.Open(origin)

    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        s := scanner.Text()
        splitIdx := strings.LastIndex(s, "]") 

        res[getAPIAnchor(s[1:splitIdx])] = s[splitIdx + 1 :]
    }

    if err := scanner.Err(); err != nil {
        return nil, err
	}

    return res, nil
}


// loads original function names and theirs id's and returns a slice with
// partially filled  LoadInfo structures (newId and oldId. We won't change
// link names, only references to headings)
func loadOrigins() ([]*LinkInfo, error) {
    res := make([]*LinkInfo, 0)

    origin := "api_orig_index.txt"

    f, err := os.Open(origin)

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

// Returns new anchor for function.
func getAPIAnchor(s string) string {
    sl := strings.ToLower(s)

    needDash := true

    res := make([]rune, 0)

    re := regexp.MustCompile("[a-zA-Z0-9·]")

    for i, w := 0, 0; i < len(sl); i += w {
            rune, width := utf8.DecodeRuneInString(sl[i:])
            w = width

            // Slug libs don't work with these dots, so hardcode this case
            if re.MatchString(string(rune)) {
                res = append(res, rune)
                needDash = true
                continue
            }

            if needDash {
                res = append(res, '-')
                needDash = false
            }
    }

    if !needDash {
        res = res[:len(res) - 1]
    }

    return strings.Trim(string(res), " ")

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
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		fmt.Println(path)

        newContents := string(read)

        for i := range res {
            old := fmt.Sprintf("(#%s)", res[i].oldId)
            new := fmt.Sprintf("(%s)", res[i].url)
            newContents = strings.Replace(string(newContents), old,  new, -1)
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
