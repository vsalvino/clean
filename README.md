clean
=====

*The quick little project cleaner.*

`clean` is a tiny little utility to recursively delete compiled output, editor
cruft, etc. from your projects. It runs in a split second as it will completely
skip over programming language environments such as Python virtual envs,
node_modules, and any hidden dot directories.


Installation
------------

1. [Download a copy for your platform](https://github.com/vsalvino/clean/releases).

2. Rename the downloaded file to `clean` or `clean.exe`.

3. Put it somewhere on your PATH.

   * Linux/mac: I recommend `/usr/local/bin/`. Then mark as executable with
     `chmod +x clean`

   * Windows: I recommend creating a `tools` folder in your home folder (e.g.
     `C:\Users\YourName\tools\`) and then updating your PATH to include that
     folder. Then you can put all your own scripts and tools here.

4. Open a shell, `cd` to a project, and type `clean`.

Run `clean -help` to see full options and usage instructions.


What does it clean?
-------------------

* Compiled output (C, C++, Go, etc.) - removes:
  * `.app`
  * `.dll`
  * `.dylib`
  * `.exe`
  * `.idb`
  * `.ko`
  * `.o`
  * `.obj`
  * `.out`
  * `.pdb`
  * `.so`
  * `.test`

* Python - removes:
  * `__pycache__` directories and their contents.
  * `.pyc` files.

* Editors - removes Emacs/Vim backups and autosaves.


Why not use `find`?
-------------------

`find` is the venerable old standard, but the down side is that it still has to
descend into all directories to do regex comparisons against every path to
decide whether or not to include/exclude. On a large project with huge
node_modules, local Python virtual env, etc. this still takes a lot of time to
run (especially on NTFS, or worse on a spinning disk!)

Configuring `find` to delete/ignore everything you want also requires an
unweidly amount of options, arguments, and regular expressions which makes it
difficult to share with other team members (everyone having to pass around their
shell aliases and scripts).


Inspiration
-----------

This is inspired by the `pyclean` command that ships with Debian, which is used
to delete `.pyc` files before uploading/shipping python code.


Contributing
------------

Want to add support for your languages/projects? Please submit a pull request!

Future features may include ability to specify what to delete or what to ignore
using a config file, which would be useful for customizing within large
projects.
