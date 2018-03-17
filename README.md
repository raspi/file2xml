# File2XML
Convert any file to handy byte safe XML

Inspired by http://thedailywtf.com/articles/Oh%2c-XML

## Example

File test.txt contents:
```
test
```

output.xml:
```
<?xml version="1.0" encoding="UTF-8"?>
<Data>
  <Byte pos="0">
    <Bit pos="0" Enabled="false"></Bit>
    <Bit pos="1" Enabled="true"></Bit>
    <Bit pos="2" Enabled="true"></Bit>
    <Bit pos="3" Enabled="true"></Bit>
    <Bit pos="4" Enabled="false"></Bit>
    <Bit pos="5" Enabled="true"></Bit>
    <Bit pos="6" Enabled="false"></Bit>
    <Bit pos="7" Enabled="false"></Bit>
  </Byte>
  <Byte pos="1">
    <Bit pos="0" Enabled="false"></Bit>
    <Bit pos="1" Enabled="true"></Bit>
    <Bit pos="2" Enabled="true"></Bit>
    <Bit pos="3" Enabled="false"></Bit>
    <Bit pos="4" Enabled="false"></Bit>
    <Bit pos="5" Enabled="true"></Bit>
    <Bit pos="6" Enabled="false"></Bit>
    <Bit pos="7" Enabled="true"></Bit>
  </Byte>
  <Byte pos="2">
    <Bit pos="0" Enabled="false"></Bit>
    <Bit pos="1" Enabled="true"></Bit>
    <Bit pos="2" Enabled="true"></Bit>
    <Bit pos="3" Enabled="true"></Bit>
    <Bit pos="4" Enabled="false"></Bit>
    <Bit pos="5" Enabled="false"></Bit>
    <Bit pos="6" Enabled="true"></Bit>
    <Bit pos="7" Enabled="true"></Bit>
  </Byte>
  <Byte pos="3">
    <Bit pos="0" Enabled="false"></Bit>
    <Bit pos="1" Enabled="true"></Bit>
    <Bit pos="2" Enabled="true"></Bit>
    <Bit pos="3" Enabled="true"></Bit>
    <Bit pos="4" Enabled="false"></Bit>
    <Bit pos="5" Enabled="true"></Bit>
    <Bit pos="6" Enabled="false"></Bit>
    <Bit pos="7" Enabled="false"></Bit>
  </Byte>
  <Byte pos="4">
    <Bit pos="0" Enabled="false"></Bit>
    <Bit pos="1" Enabled="false"></Bit>
    <Bit pos="2" Enabled="false"></Bit>
    <Bit pos="3" Enabled="false"></Bit>
    <Bit pos="4" Enabled="true"></Bit>
    <Bit pos="5" Enabled="false"></Bit>
    <Bit pos="6" Enabled="true"></Bit>
    <Bit pos="7" Enabled="false"></Bit>
  </Byte>
</Data>
```
