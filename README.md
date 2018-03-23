# oraconn
All this does is take a oracle database configuration stored in a config file to generate a connection file that will be passed in making the actual connection. I solely created this for my convenience, it has virtually no function if you aren't looking for Oracle connection strings with a properly setup `tnsnames.ora`.  90% of the work being done by this is from the incredibly useful [viper](https://github.com/spf13/viper), I would not be surprised or insulted if you could write something better in 10 minutes, or 5.

To use the package just:
```
import (
    "github.com/jazzboME/oraconn"
)
```
And use it as:
```
connStr, err := oraconn.Do() // use 'database' config file
```
or
```
connStr.err := oraconn.Do("prod") // use 'prod' config file 
```

By default it looks for a database configuration file in the present directory and then in your home directory. I use it so I can I can easily switch between databases in testing new apps.

I like TOML for my config files, but your config needs to specify a credential section and a database section.
```
title = "Option Config Title"

[credentials]
user = "scott"
password = "tiger"

[database]
tns = "orcl.example.com"
```

Insert all the usual caveats that this is putting your pw in plaintext, so you should only use this package if you feel you can do that sanely.  At some point I'll work on something more secure.

I also plan to look into adding options for placement of the config file, and non-TNS setups.  But again this is for quick prototyping of my work, not something I expect a ton of people to be using.
