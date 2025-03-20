# hoverfly-convert-xml-to-json-util
## Converts an XML file to a JSON string consisting only of arrays and outputs to the console

The JSON that gets generated from the XML will consist only of arrays to support looping within the {{#each}} block used by the templating engine within Hoverfly. It will be exactly the same as the JSON that gets generated in Hoverfly with the jsonpathfromxml function.

This software utilises a fork of https://github.com/basgys/goxml2json with a modification to reflect everything as an array so that the resulting json can be iterated over using {{#each}}.

### Usage: convertxml "xml file name"

### Example: convertxml test.xml

Where text.xml contains
```XML
<lineitems>
    <lineitem>
        <upc>1</upc>
    </lineitem>
    <lineitem>
        <upc>2</upc>
    </lineitem>
    <lineitem>
        <upc>3</upc>
    </lineitem>
</lineitems>
```

Will return to the console (in a single line):

```JSON
{
    "lineitems": [
        {
            "lineitem": [
                {
                    "upc": [
                        "1"
                    ]
                },
                {
                    "upc": [
                        "2"
                    ]
                },
                {
                    "upc": [
                        "3"
                    ]
                }
            ]
        }
    ]
}
```
