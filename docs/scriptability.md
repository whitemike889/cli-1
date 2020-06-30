# Scriptability audit

This document acts as a high level TODO list for improving scriptability of `gh` across all of our
commands.

## gh pr

### create

- [x] Behavior when `STDIN` is a pipe

  *Current*: It breaks. Because STDIN is occupied, survey fails to prompt.
  *Desired*: It doesn't break. It'd be cool if we could take PR data to pre-fill the prompts with
  but that opens up questions of how to serialize the input. For now it should just behave as if
  --fill was passed. 
  
- [x] Behavior when `STDOUT` is piped out

  *Current*: It breaks. survey writes to the pipe and the command hangs indefinitely.
  *Desired*: It doesn't break. --fill should be assumed since it's impossible to prompt.

- [x] Behavior when run un-attached to a terminal

  *Current*: It hangs still attempting to prompt.
  *Desired*: --fill should be assumed.

- [x] Consistent use of `STDERR`

  *Current*: Almost all output in the non-interactive case goes to STDERR.
  *Desired*: All output _except_ the final PR URL is sent to STDERR in the non-interactive case.

### close

- [x] Behavior when `STDIN` is piped to command

  *Current*: Argument validation error.
  *Desired*: STDIN should be parsed as a single PR argument.

- [x] Behavior when `STDOUT` is piped out
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

### checkout

- [x] Behavior when `STDIN` is piped to command

  *Current*: Validation error
  *Desired*: STDIN parsed as a single PR argument

- [x] Behavior when `STDOUT` is piped out
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

### diff

- [x] Behavior when `STDIN` is piped to command

  *Current*: STDIN is ignored and current branch's PR is used, if any.
  *Desired*: STDIN is parsed as a single PR argument

- [x] Behavior when `STDOUT` is piped out
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

### list

- [x] Behavior when `STDIN` is piped to command
- [x] Behavior when `STDOUT` is piped out

  *Current*: Just the PR list is printed in a slightly different format
  *Desired*: The format is more machine-friendly and conveys what is lost without color. It should
  be easy to process with `cut` or `awk`.
    
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

  *Current*: "Showing x of y pull requests in owner/repo" is always printed to STDERR
  *Desired*: Nothing is printed to STDERR (except warnings/errors) if running non-interactively

### merge

- [x] Behavior when `STDIN` is piped to command

  *Current*: Broken. Survey tries and fails to prompt user.
  *Desired*: If a pipe, STDIN should be parsed as a single PR argument and then the merge should run
  non-interactively.

- [x] Behavior when `STDOUT` is piped out

  *Current*: Broken. Survey hangs forever.
  *Desired*: Non-interactive is assumed.

- [x] Behavior when run un-attached to a terminal

  *Current*: Broken.
  *Desired*: Non-interactive is assumed.

- [x] Consistent use of `STDERR`

  *Current*: All output goes to STDOUT
  *Desired*: Output should go to STDERR

### ready

- [x] Behavior when `STDIN` is piped to command

  *Current*: STDIN is ignored
  *Desired*: STDIN is parsed as a single PR argument.

- [x] Behavior when `STDOUT` is piped out
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

### reopen

- [x] Behavior when `STDIN` is piped to command

  *Current*: Validation error
  *Desired*: STDIN is parsed as a single PR argument.

- [x] Behavior when `STDOUT` is piped out
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

### review

- [x] Behavior when `STDIN` is piped to command

  *Current*: STDIN is ignored
  *Desired*: STDIN is parsed as a single PR argument.

- [x] Behavior when `STDOUT` is piped out

  *Current*: Hangs waiting for input.
  *Desired*: If lacking enough flags to run non-interactively, clear error. Otherwise just run.

- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

  *Current*: Outputs everything to STDOUT
  *Desired*: Should send current output to STDERR but print URL of review/comment to STDOUT

### status

- [x] Behavior when `STDIN` is piped to command
- [x] Behavior when `STDOUT` is piped out

  *Current*: Same human-oriented output is printed
  *Desired*: `grep`/`awk`/`cut` friendly output

- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

### view

- [x] Behavior when `STDIN` is piped to command

  *Current*: Ignored
  *Desired*: STDIN parsed as single PR argument

- [x] Behavior when `STDOUT` is piped out

  *Current*: Rendered markdown and human metadata written to pipe
  *Desired*: Raw markdown is printed, possibly made possible via a --raw flag. Metadata printed
  linewise for grepping.

- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`


## gh issue

### close

- [x] Behavior when `STDIN` is piped to command

  *Current*: Argument validation error.
  *Desired*: STDIN should be parsed as a single PR argument.

- [x] Behavior when `STDOUT` is piped out
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

### create

- [x] Behavior when `STDIN` is piped to command

  *Current*: Breaks. Survey attempts to parse STDIN.
  *Desired*: Doesn't break but errors clearly saying that piping on STDIN is unsupported

- [x] Behavior when `STDOUT` is piped out

  *Current*: Hangs. Survey writing to pipe.
  *Desired*: Clear error

- [x] Behavior when run un-attached to a terminal

  *Current*: Hangs.
  *Desired*: Clear error

- [x] Consistent use of `STDERR`

### list

- [x] Behavior when `STDIN` is piped to command
- [x] Behavior when `STDOUT` is piped out

  *Current*: Just the issue list is printed in a slightly different format
  *Desired*: The format is more machine-friendly and conveys what is lost without color. It should
  be easy to process with `cut` or `awk`.
    
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

  *Current*: "Showing x of y pull requests in owner/repo" is always printed to STDERR
  *Desired*: Nothing is printed to STDERR (except warnings/errors) if running non-interactively

### reopen

- [x] Behavior when `STDIN` is piped to command

  *Current*: Validation error
  *Desired*: STDIN is parsed as a single issue argument.

- [x] Behavior when `STDOUT` is piped out
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

### status

- [x] Behavior when `STDIN` is piped to command
- [x] Behavior when `STDOUT` is piped out

  *Current*: Same human-oriented output is printed
  *Desired*: `grep`/`awk`/`cut` friendly output

- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

### view

- [x] Behavior when `STDIN` is piped to command

  *Current*: Ignored
  *Desired*: STDIN parsed as single issue argument

- [x] Behavior when `STDOUT` is piped out

  *Current*: Rendered markdown and human metadata written to pipe
  *Desired*: Raw markdown is printed, possibly made possible via a --raw flag. Metadata printed
  linewise for grepping.

- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`


## gh repo

### clone

- [x] Behavior when `STDIN` is piped to command

  *Current*: Ignored
  *Desired*: STDIN parsed as single repo argument

- [x] Behavior when `STDOUT` is piped out
- [x] Behavior when run un-attached to a terminal
- [x] Consistent use of `STDERR`

  *Current*: Output from git is on both STDERR and STDOUT. We print nothing ourselves to STDOUT.
  *Desired*: Git information should all be on stderr and potentially filtered (without a --debug or
  --verbose option). Only thing on STDOUT should be resulting URL of clone.

### create

- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`

### fork

- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`

### view

- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`


## gh alias

### delete

- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`

### list

- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`

### set

- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`


## gh api


- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`


## gh config

### set

- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`

### get 

- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`


## gh gist

### create

- [ ] Behavior when `STDIN` is piped to command
- [ ] Behavior when `STDOUT` is piped out
- [ ] Behavior when run un-attached to a terminal
- [ ] Consistent use of `STDERR`
