package xmlTest

import (
    "encoding/xml"
    "fmt"
    "strings"
)

func parseXML(input string) {
    reader := strings.NewReader(input)

    decoder := xml.NewDecoder(reader)
    for {
        t, err := decoder.Token()
        if err != nil {
            fmt.Println("decoder xml failed, err: ", err)
            break
        }
        switch token := t.(type) {
        case xml.StartElement:
            name := token.Name.Local
            fmt.Printf("Token name: %s\n", name)
            for _, attr := range token.Attr {
                attrName := attr.Name.Local
                attrValue := attr.Value
                fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
            }
        case xml.EndElement:
            fmt.Printf("Token of %s end\n", token.Name.Local)
        case xml.CharData:
            content := string([]byte(token))
            fmt.Printf("This is the content: %v\n", content)
        default:
            //
        }
    }
}
