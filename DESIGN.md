## Design

User defined hierarchy rules (call them "rules") written like:
```
MEASUREMENT>region>app>FIELD
```

A rule maps the items to hierarchical tiers.  In the above, `MEASUREMENT` is top tier, `region` is second from top and so on. 

The API will reach out to InfluxDB and collect the Measurements and associated Tag keys, values, and Fields and match them to tiers according to the defined rule.

In this example, the values in the `_measurement` column will be tier 1.  The values in the `region` column will be tier 2 and so on.  

### How it works:
For each Bucket-Measurement, all Tag keys are retrieved from the instance.  If a tree rule was set by the user, each key is checked against the tokenized rule and given its tier (index in the rule). 

If the top tier (index 0) is a Tag, all values are retrieved for this Tag key and those values' parent value is set to the key.  For each value, they are given a child value of the next "word" in the index (next tier). If that word is a Tag key again, all values for that Tag Key are retrieved and assigned as children of the key.

Tag keys aren't the only tiers, however.  If a "word" is a Measurement, that tier is set to that and then the next "word" is checked for its type (Tag key or Field).  If Tag key, see above.  If it's a Field, that's the end/leaf of the tree. 

### Goal:
Create a tree of Measurements, their Tag keys-values and Field keys. Expose an API for a UI to render this tree.  When the UI interaction reaches a leaf (Field), the UI will query InfluxDB for that value (latest value only, for now).  Intended to be lazy.

The tree (excluding actual metric data) will be kept in memory by the server. Server startup should rehydrate it automatically.

### Non-goal:
Get all Field data back and stored in UI memory.  Evaluation of leaf nodes should be lazy.

### Tokenization of user input
The user's input will be tokenized into words (operands) and operators (`>` or `|`). An index for each should be retained so we know their order. Second, consecutive words split by `|` will be determined to be siblings of one another.  Siblings can be 2 or more. A case of `sib1|sib2|sib3|sib4` should be possible in the user's input and all four words would be siblings at the same level.  

[**Not sure here**]
Then parents will be identified.  [**Can this be dealt with just with the order they're in?**]

### Parsing of InfluxDB output





### Mapping of parsed Influx data to tokenized user input 