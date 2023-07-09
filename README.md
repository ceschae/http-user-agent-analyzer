# http-user-agent-analyzer

## TL;DR

The HTTP User Agent Analyzer is the first set of features and the general framework 
for an HTTP Analyzer tool with a natural-ish language[<sup>1</sup>](#subscripts) interface. Although the analysis
functionality is not new itself, the grammar for the TUI is... unusual, because it's asking questions of the terminal, rather than issuing commands:

```bash
http-user-agent-analyzer is "./filepath/or/rawstring.txt" allowed
```

[Skip to the technical documentation]()


## ... Why Would You Do That

CLI (Command Line Interface) tools can be a gatekeeping experience for non-technical people. In a modern, GUI-driven (Graphical User Interface) world where most folks only use computers to interact with web browsers, the average computer user does not know what a CLI tool is, or what the conventions are[<sup>2</sup>](#subscripts). Good news! We can bring the usability to them, instead of expecting them to adhere to the old structures of coding.

## How Is That Even Possible

Based on my years of experience in developer tooling, and with my Master's in Library Science in hand, I can confidently say that people think in their native language for the most part. Why shouldn't our tooling reflect that? Therefore, I'm here to introduce a **question grammar** to CLI tooling to provide a more familiar TUI (Textual User Interface) experience.  

```
<command> := <app> <question word> <question arguments>
<question word> := is|what|do|who|why
<question arguments> := <flags> <adjectives>
<flags> := | -f <file name> | <insert other flags>
<adjectives> := | allowed | <insert other adjectives>
```
(Not familiar with this format? It's [Backus-Naur form](https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form))

Obviously this isn't a complete grammar, but ideally it demonstrates the point. This paradigm is designed to have us ask questions of our TUI, rather than have us know the format it needs. 

## No Seriously, Why Would You Do That

Meeting people where they're at is the core value of a good UX (User Experience). If we can eliminate the step of having to memorize Yet Another<sup>TM</sup> set of conventions that's separate from how our language currently operates, then that's one more barrier broken down to get more folks involved in tech.[<sup>3</sup>](#subscripts)

## Okay So How Does This Work

You ask it questions based on the [technical documentation](#technical-documentation) of allowed words, specified in the next section. Currently this strategy is unintuitive to those used to tab-complete, but meeting that set of folks where they are can be solved with some neat NLP (Natural Language Processing) parsers, like [this one from Stanford](https://nlp.stanford.edu/software/lex-parser.shtml). 

## Technical Documentation
Listed in this section is the current controlled vocabulary of the current state of this tool.

| Question Word | Description | Flags | Adjectives |
| --- | --- | --- | ---  |
| `is` | the `is` question expects a `true`/`false` answer | `"./filepath/or/rawstring.txt"` is the **obejct** of this question, and should be either a filepath or a complete string. The program will figure out which it is for you. | `allowed` currently will return `true` if the `User Agent` header in the **object** of the question indicates the request is from Firefox, `false` otherwise |

## Subscripts 
<sup>1</sup>in the colloquial sense, not the Natural Language Processing (NLP) sense

<sup>2</sup> This is based off of when I was an Intro Coding TA at the University of Washington: Seattle, where, consistently, 75% of our students in that course had never coded before.

<sup>3</sup> This theory is based on _The Design of Everyday Things_ by Don Norman ([Goodreads link](https://www.goodreads.com/book/show/840.The_Design_of_Everyday_Things)). 