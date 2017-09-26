## Aptitude, Experience, Skills

I can't even describe my skills.  I don't think our industry can either.
We talk about years of experience when I know it doesn't work that way.  Some years
I'm expanding greatly and others I'm expected to ship and perform.  It isn't even
broken up by years like that.  Maybe it's "have extra energy on weekend 34, 41 and 42".

So how do I even talk about language skills?


## Answering It Myself

I don't have an answer, no one has an answer so I can just find out myself.

Maybe I'm not correct in my base assumption but a good starting point would be
"I can make a TODO list".  For example, I've touched Node.js many times but I was doing Ruby
during the "hype transition" and I never transitioned to it being my primary language.
So ... do I say "I know Node"?  What does that even mean?  So I'm tired of speaking
like this so I'm just going to create TODO lists in every language I can.

Some languages (like Elm for me) are going to be a "first synthesis app" for me.
So part of this repo is to track and log how slow I am in the languages where I'm growing.


## This is not

This doesn't demonstrate design or architecture.  This is a well-known problem
that doesn't need to be broken down too much or thought about from different angles.
This is purely syntax and implementation (hopefully).


## Completed Projects

### Ruby

| Time to Complete | Feeling | 
| ---------------- | ------- |
| 3 hours (minus TUI) | Familiar, coded off the top of my head |

| Library | Use |
| ------- | --- |
| Rspec   | Testing |

Liked:
* TDD flow, rspec syntax

Researched:
* `Array.reject` chaining with `each_with_index`

Got stuck on:
* Nothing

New to me:
* Never done an ncurses TUI before in Ruby

Disliked:
* Interactivity means I need a TUI.  A library demos very poorly.  How would I take a screenshot of a library?  Jealous of frontend projects.

Wish I had done:
* Wish I had a dummy data generator like factory girl in there.
Not worth it for this kind of project but would have been nicer I guess?

### Go

| Time to Complete | Feeling | 
| ---------------- | ------- |
| 5.5 hours | Mostly quick |

| Library | Use |
| ------- | --- |
| Ginkgo | Testing |
| Tablewriter | Pretty printing a text table |
| gb (not a lib) | Dependencies and vendoring |

Liked:
* Ginkgo watching, test-first flow
* Gomega's matchers

Researched:
* Deleting a slice element by index.
* Makefile stuff like exit codes `$$?` instead of `$?`.
* Golang's text/template pkg.
* How to collect or map over a slice.  Pretty easy, just not a one-liner.

Got stuck on:
* Nothing

New to me:
* Never did Ginkgo this much
* The tablewriter pkg

Disliked:
* Fake interactivity, but I'm not going to do a TUI (too much like making a web browser).
* I tried writing my own text renderer but I didn't like it, so I just used a pkg.
* My Todo struct stayed unexported (lowercase) until I did the `main.go`.  Ah well.

Wish I had done:
* Looked at ncurses libs?

Output:
```
$ cd go
$ make build  # assume gb already installed
$ ./bin/todo
         TODO        | COMPLETED
+--------------------+-----------+
  Do a little dance  | [ ]
  Make a little love | [ ]
  Get down tonight   | [x]
```
You'll need to install gb before the above will work.
