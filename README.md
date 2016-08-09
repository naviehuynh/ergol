ERGOL - pluggable log viewer
============

### DOING

__MVP__
- basic structure (source + filter + displayer, configuration)
- command line parser
- printer logging filename and color, with local timestamp

### DONE

__PoC__
- rough structure
- follow changes
- multiple files

### TODOs
- follow file path, not inode
- input from STDOUT/STDERR of command
- input from files over SSH
- filter/highlight by phrase/regex/wildcard/fuzzy matching
- log parser (fields include: type, level, date range, exception, count)
- config file
- vim navigation
- pin/unpin log
- text index
- grep -C
- history (recent files, recent commands/filter/search/highlight)
- help text (press `?` to show current active key bindings)
- stdin source
- consider process launcher
- testing

### FEATURES
<!-- - tail multiple logs -->
<!-- - output different color for each file -->
<!-- - configuration file -->
<!-- - text highlight -->
<!-- - text filter -->
<!-- - vim binding navigation -->
<!-- - regexp ? -->
<!-- - pin log -->
<!-- - parse content (exception, level, json etc) -->
<!-- - history -->
<!-- - text search -->
<!-- - help keys revealing current binding -->
<!-- - pipe log into log viewer or process launcher pattern? -->
