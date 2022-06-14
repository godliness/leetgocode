# PDB code formating

JenkinsCI will automatically reject any PRs which do not follow our coding conventions. The easiest way to ensure your PR adheres to those conventions is to use the pdb_indent tool. This tool uses `clang-format` under the hood.

This tool only check *.c files which are added by us, we don't check the upstream files, if you cherrypicked the upstream files which are new added into project, please also update the .gitattributes file and add the new upstream file into .gitattributes with attribute "-pdb-style" (black list).

## Install clang-format, you must make sure the version is 14.0.5:
[clang-format-style-options](https://clang.llvm.org/docs/ClangFormatStyleOptions.html)

[clang-format-style-options.14](https://releases.llvm.org/14.0.0/tools/clang/docs/ClangFormatStyleOptions.html)

[clang-format.14](https://releases.llvm.org/14.0.0/tools/clang/docs/ClangFormat.html)

**For Ubuntu20.04**

Uninstall old version package if you have

```
apt-get remove clang-format-14
```

Install new version package and the installation file is <a href="https://github.com/0penpie/piecloud-db-CLUSTER/raw/master/src/tools/check-style/clang-format-14.0.5.deb" target="_blank">here</a>

```
apt-get install -f ./clang-format-14.0.5.deb
rm -f /usr/bin/clang-format
ln -s /usr/bin/clang-format-14 /usr/bin/clang-format
```

**For MacOS**
Uninstall old version package if you have

```
brew uninstall clang-format@14
```

Install new version package and the installation file is <a href="[https://10.24.15.41/software/clang-format--14.0.5.big_sur.bottle.tar.gz](https://github.com/0penpie/piecloud-db-CLUSTER/raw/master/src/tools/check-style/clang-format--14.0.5.big_sur.bottle.tar.gz)">here</a>

```
brew install -f ./clang-format--14.0.5.big_sur.bottle.tar.gz
```

## Check your code and format it (If you want to check and format code by mannual without IDE)

Go to the top level of our project and run below command to check the code style:

```
make stylecheck
```

and then run the command as below will format them:

```
make styleformat
```

## About some clang-format arguments and features you need to known

Go to the top level of project, run below commands will format specified files(clang-format --help for details):
```
clang-format -i file1 file2 file3
```
Below will only check specified files and not format them(clang-format --help for details):
```
clang-format -i file1 file2 file3 --dry-run
```

Disabling formatting on a piece of strange code that clang-format does not format it correctly by using clang-format comment as below:
```
int formatted_code;
/* clang-format off */
    void    unformatted_code  ;
/* clang-format on */
void formatted_code_again;
```
Nested function with pointer and double pointers, the alignment of pointers is not right, there is weird formated results:
https://github.com/llvm/llvm-project/issues/55745
Consecutive variable declarations binpack aligments has problems:
https://github.com/llvm/llvm-project/issues/55792

So please use this trick for this kind of case.

## Integrate clang-format into your editor

**clang-format config file `.clang-format` located on top level of our project which you could be used to integrate it into your editor.**

[int.emacs]: https://clang.llvm.org/docs/ClangFormat.html#emacs-integration
 * LLVM's official [Emacs integration][int.emacs] provides the _`clang-format-region`_ macro.

[int.vim]: https://clang.llvm.org/docs/ClangFormat.html#vim-integration
 * [Vim integration][int.vim] as documented by the official LLVM project

[int.altvim]: https://github.com/rhysd/vim-clang-format
 * A community [VIM plugin][int.altvim] that looks promising and claims to be better than the official VIM integration

[int.clion]: https://www.jetbrains.com/help/clion/clangformat-as-alternative-formatter.html
 * [Clion][int.clion] detects `.clang-format` automatically and switches to using it for the built-in "Reformat Code" and "Reformat File...".

[int.vscode]: https://code.visualstudio.com/docs/cpp/cpp-ide#_code-formatting
 * Similar to CLion, [Visual Studio Code has it built in][int.vscode].

[int.xcode8]: https://github.com/mapbox/XcodeClangFormat
 * [Plugin for Xcode 8+][int.xcode8]

## Code style conventions

```
Consecutive variables and functions definition will be aligned with tab, if not enough one tab use spaces instead
Pointer alignment is right side to the var name, eg void *test vs void* test
Limit code width to 90 columns, more then 90 columes will make a break line
Formatting uses 4 columns tab spacing
Each logical indentation level is one additional tab stop
Curly braces for the controlled blocks of if, while, switch, etc go on their own lines
Do not use C++ style comments (// comments). pdb_indent will replace them with /* ... */ and add a space before online comment.
Add space around "+", "-", "%", "<<", "=", "+=", "||", "&&", "<", ">", "==", etc
Do not use "{}" in a sinle-line "if", "while", "for" statement
Align function args and variable uses tab, if not enough one tab use spaces instead
Add space before '(' of control statements ('if', 'for', 'switch', 'while', etc.)
Add space after ',', i.e. 'a,b' vs. 'a, b'
Align function params, variable definition and marco definition with tabs, if not enough one tab(4 spaces) use spaces instead
```
