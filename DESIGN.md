## Design

User defined hierarchy rules (call them "rules") written like:
```
MEASUREMENT>region>app>FIELD
```

A rule maps the items to hierarchical tiers.  In the above, `MEASUREMENT` is top tier, `region` is second from top and so on. 

The API will reach out to InfluxDB and collect the Measurements and associated Tag keys, values, and Fields and match them to tiers according to the defined rule.

In this example, the values in the `_measurement` column will be tier 1.  The values in the `region` column will be tier 2 and so on.  

### Tokenization of user input
The user's input will be tokenized into words (operands) and operators (`>` or `|`). An index for each should be retained so we know their order. Second, consecutive words split by `|` will be determined to be siblings of one another.  Siblings can be 2 or more. A case of `sib1|sib2|sib3|sib4` should be possible in the user's input and all four words would be siblings at the same level.  

[**Not sure here**]
Then parents will be identified.  [**Can this be dealt with just with the order they're in?**]

### Parsing of InfluxDB output





### Mapping of parsed Influx data to tokenized user input 