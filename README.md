# Purpose

This program replace input json with some command output, or file body.
You needs to specify Value-Type (int,float,string,[string]), and
Target-Property("." separated key path: "prop1.prop2").

# Install

go get github.com/tmrtmhr/filljson

# Usage

```
cat template.json | filljson ${ValueType} path.to.target <(some command)
```

# Example

## template.json

```
{
  "prop1": {
    "prop2": {
      "target_prop": null
    }
  }
}
```

## `ls` Output

```
aaa
bbb
ccc
```

## Command

```
cat template.json | filljson [string] prop1.prop2.target_prop <(ls)
```

## Output Json

```
{
  "prop1": {
    "prop2": {
      "target_prop": ["aaa","bbb","ccc"]
    }
  }
}
```

Also, you can use UNIX pipe like following commands.

```
cat template.json |
  filljson int "a.b.c" <(foo command) |
  filljson [string] "x.y.z" <(bar command) > output.json
```

enjoy!
