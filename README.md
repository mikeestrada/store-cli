# store-cli


## Install Redis
`brew install redis`

`brew services start redis`

Run `redis-cli PING` to confirm cache is running


## Download & Install go lang
https://golang.org/doc/install


## Download & Install storecli on your path
Run `git clone https://github.com/mikeestrada/store-cli.git && go build -o storecli && cp ./storecli $GOPATH/bin`



## Usage
### ADD
Adds a member to a collection for a given key. Displays an error if the member already exists for the key.

### KEYS
Returns all the keys in the dictionary.  Order is not guaranteed.

### MEMBERS
Returns the collection of strings for the given key.  Return order is not guaranteed.  Returns an error if the key does not exists.

### ADD
Adds a member to a collection for a given key. Displays an error if the member already exists for the key.

### REMOVE
Removes a member from a key.  If the last member is removed from the key, the key is removed from the dictionary. If the key or member does not exist, displays an error.

### REMOVEALL
Removes all members for a key and removes the key from the dictionary. Returns an error if the key does not exist.

### CLEAR
Removes all keys and all members from the dictionary.

### KEYEXISTS
Returns whether a key exists or not.

### MEMBEREXISTS
Returns whether a member exists within a key.  Returns false if the key does not exist.

### ALLMEMBERS
Returns all the members in the dictionary.  Returns nothing if there are none. Order is not guaranteed.

### ITEMS
Returns all keys in the dictionary and all of their members.  Returns nothing if there are none.  Order is not guaranteed.
