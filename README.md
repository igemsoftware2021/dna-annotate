<p align="center">
  <br>
  <img src="https://cdn.discordapp.com/attachments/862687453516660746/892128174408892447/friendzymes_freddy_square.png" width="150" />
  <!-- <img src="friendzymes-blah.png" width="150"/> -->
  <br>
</p>

# dna-annotate

A Github action to annotate problematic sequences from given Genbank file.

dna-annotate is a pretty great tool for low-attention readers like yourself who want the executive summary. This is executive-land, where all the C Suite execs live it up, only reading above the fold and never reading the fine print. Whee!

This is the segue so you don't get scared when the detail comes in.

## All options

### List of Options

Every argument is required.

| Option          | Description                                                              | Default           |
| --------------- | ------------------------------------------------------------------------ | ----------------- |
| [input-dir]     | Directory with all genbank to be read and annotated                      | `input`           |
| [input-pattern] | User-specified pattern to select specific files from the input directory | `.*\.\(gb\|gbk\)` |
| [output-dir]    | Directory where all genbank annotated will be written                    | `output`          |

### Detailed Options

#### input-dir

Phasellus ac felis auctor, molestie libero at, posuere tellus. Morbi interdum interdum viverra.
Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.

Default: `input`

#### input-pattern

Phasellus ac felis auctor, molestie libero at, posuere tellus. Morbi interdum interdum viverra.
Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.

- mention that syntax is [re2]
- regex tester
  - preload with relevant regexps
  - preload with relevant test strings

Default: `.*\.\(gb\|gbk\)`

#### output-dir

Phasellus ac felis auctor, molestie libero at, posuere tellus. Morbi interdum interdum viverra.
Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.

Default: `output`

### Usage

Basic:

```yaml
name: 'Some dummy YAML'
```

See [action.yml] for a comprehensive list of all the options.

See [Friendzymes Cookbook] for further examples and sample data.

---

[Friendzymes Cookbook]:       <https://github.com/Open-Science-Global/friendzymes-cookbook>
[re2]:                        <https://github.com/google/re2/wiki/Syntax>

[action.yml]:                 ./action.yml
[input-dir]:                  #input-dir
[input-pattern]:              #input-pattern
[output-dir]:                 #output-dir