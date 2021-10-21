<p align="center">
  <br>
  <img src="friendzymes.svg" width="300" />
  <br>
</p>

# dna-annotate [![Friendzymes Cookbook](https://img.shields.io/static/v1?label=friendzymes%20cookbook&message=dna-annotate&labelColor=F8F1EF&color=FF8552&style=for-the-badge&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAAAAXNSR0IArs4c6QAABQ9JREFUaEPtmG1MU1cYx//nkirUIa5x+AKOoWQjLjhxLO4lS/jAXNSY7MN0GwS9haRl0Ftlzi+LEadzmRrfWoG2GntRsyzKXJYZnbJkzGwwiMiLMRku6KKbLxWQKTitvfcst1LXGem5195iTHqSJs09z/0//995ntPTewme8EGecP+IAzzuCsYrEK9AlCugSwvx/LYJMHCLZHC9e3cJR6P0pOn2RwYwl1dPo5LsADBh+DM7mJnSPaLHXvqgC7PV+R4FlhKC7oSAtHH37hVXNTkdIfiRAXirYz9Aih6mS4C9FGQ/KL0CjrxPKZ1HgLxQLAE8XrdgfWwARYJjvMGP8wAxhZs41320Zcvm9Qlnz55NbWg4nnbgwNeBJcUbxj7MqESljH2eFReihdBUAb7MwQNSI5CQD0q84clzc8Z12yvML4RfW7duHQ5929yb+8rCiaHrJxq249atW7jjR8v1/suvjhoAb915HqDPjZSwepulzZg05uUH5+fPn4/JGQuDl3NnyrDb7cHvXV1dQ8XFJRdz55rfrdtlO/OoIKoqwFudtQDKIiXZ/Hlh50ST6SUlpqenp2nLlq0JNTXVc202G4YC2f8zH9Lp6Oi4umn7kUljE8efA0Gb6BKWaAVhAvBlznxQ/MgSnjnjbveqVZXBFsrJmfVXVlZWb1FR4fXDhw/nN544daO97Sc5OTn5rt/v9xuNxrSQXuXK9Z0DQ08HwUFwUCsEE2BpWW0aRwN/sgD27a5Ea2vLD3PmzClYvXr1kfr6+lmFhYXp5y+Z0N763UDHqWMTUlMnX+zv75s2OHizJzExcYai+ckaR+dlH3cPAIAsY4GWs4QJoIjy1p1egPIsiJam+r8z0owp2dnZOH36NPoGDP/MzluQpNznddlQVfXpLxs2fPZGX1/vyZSUlDyfz9curHRNNY4zTQrTrhDdQg0rV2ieCVBautUkGQwuUCxWK/qwuKGb/Zfeyk/3Z2Y+21FQUPCOz3etx175xe2k5MwX78cTXCIGOtvrtF9TmysiwDKLYx7hiBcUU9UKRoqTZckfuHPld+muRAcGbydNmfp8sI3Ch9ZDLiIAb3UKAJS/C6M+RLfA7I57+z7CMFsdH1CQL0fdfdAZNYsuu8jKHRGgpGRjsmww3mCJxGae/CG6bZksbWaZeKtTOQPyWUKxmFfTRpH3gMpDLBbmQUmb6LHd/wc7Uo7Ie0BwPEP9xBcTgwxRSrGtziN8xMrNbKHS8uosKUDXgNBilpie8wRo9rqF11maTICQAG9xVoFgLUtQz/mo90C4GXOZ821K8b2eBllaah56VFdASbbM6mwiwGusxHrNc5R7c4+n4udIetoALA47IWSHXgZH0qHAN4TKXaJnObNlNQEUl9WmJlCpGaDTYweh7gAL5dcEoNw0/HrkKwCNsTngYgwQ3AsWx4w6j72HtzppDCpxQXQLGWp1NVcgXJi3ODpByCy1ydTEEUIFr8u+U02sEhMlwI61IFyV2mRhfdtMqXwchGQBZPj5mNwEoVtFl6C0puoRFYDF4p7iJ/5DAB58v6O8JvnvSWvYjvKwInOyp652eZtqh4zAqABC2mar82MAyymQDuAMAV1CQTYBuPdCCPgVVD6m5mdRK5guAErSxeXVT42T5bwhjjt5sKZiULlW8mHNdFmSxoge4TetxtTG6wagNqHecXEAvVdUq168AlpXTO/4eAX0XlGtevEKaF0xveOf+Ar8C/HOtEAxBdcIAAAAAElFTkSuQmCC)](https://github.com/Open-Science-Global/friendzymes-cookbook)

A Github action to annotate problematic sequences from given Genbank files.

dna-annotate is a Github Action that receives a path for an input directory, a regex pattern that should be used to filter genbank files or another interetsing file name pattern, and a directory where the output will be written. This action will use all this information to annotate problematic parts of a given sequence. 

Currently, dna-annotate attempts to find and annotate:

- Repetitions greater than 10 base pairs in length
- Hairpins
- Homopolymers
- Most common restriction binding sites

If you have some feature that you think will make this action better, please feel free to create an issue.

## All options

### List of Options

Every argument is required.

| Option          | Description                                                  | Default           |
| --------------- | ------------------------------------------------------------ | ----------------- |
| [input-dir]     | Directory where all the input genbank files will be read     | `input`           |
| [input-pattern] | Regex to filter files in the input directory                 | `.*\.\(gb\|gbk\)` |
| [output-dir]    | Directory where all the output genbank files will be written | `output`          |

### Detailed Options

#### input-dir

This parameter is the path of the directory for your genbank files to read and annotate. You can use this parameter to setup different pipelines for different folders, so your project can be divided in folders with different processes. By default the action will use `input` as the input directory.

Default: `input`

#### input-pattern

This parameter is a regex pattern using [re2] syntax to filter files from within [input-dir]. So even inside a given input directory, you can select a specific file or group of files for the current job. By default the action will match files with genbank extensions (`.gb` or `.gbk`).

Example: [Match only BBF10k-prefixed files][input-pattern-example-1], freegene 10k gene project parts.

Default: `.*\.\(gb\|gbk\)`

#### output-dir

This parameter is the path of the directory for outputting annotated sequences as genbank files. By default the action will use `output` as the output directory.

Default: `output`

### Usage

Basic:

```yaml
- name: dna-annotate
  uses: Open-Science-Global/dna-annotate@v0.6.1
```

See [action.yml] for a comprehensive list of all the options.

See [Friendzymes Cookbook] for further examples and sample data.

---

[Friendzymes Cookbook]:       <https://github.com/Open-Science-Global/friendzymes-cookbook>
[re2]:                        <https://github.com/google/re2/wiki/Syntax>

[action.yml]:                 ./action.yml
[input-dir]:                  #input-dir
[input-pattern]:              #input-pattern
[input-pattern-example-1]:    <https://regex101.com/library/eEz1xN>
[output-dir]:                 #output-dir
