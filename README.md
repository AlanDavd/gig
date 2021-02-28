# Gig

[![Go mod](https://img.shields.io/github/go-mod/go-version/alandavd/dsgo)](go.mod)

A CLI application to hold a historic log of done tasks separated by category.

## Install

To install the app you can download the source code:

```bash
$ git clone git@github.com:AlanDavd/gig.git
```

Then `cd` into the `gig` folder and run:

```bash
$ go install .
```

Also you can run:

```bash
$ go get "github.com/alandavd/gig"
```

## Usage

`gig` has only one purpose, save done tasks, that's why the usage of it has only two actions, list and add both categories and tasks.

### Categories

The command `new` creates a new category that will hold tasks. Note that you cannot nest categories.

```bash
$ gig new "<category name>"
```

The command `list` lists all the categories you have created so far.

```bash
$ gig list
```

Since the database `gig` makes use of is a key-value database, each category is a new bucket that holds key-value entries.

### Tasks

The command `add` creates a new task entry to a selected category.

```bash
$ gig add "<category_name>" "task description"
```

The command `list` serves both tasks and categories as well, to list tasks of a category you can use it this way:

```bash
$ gig list -t "<category_name>"
```

