InfluxDB data is structurally flat in that each dimension of each record has no hierarchical relationship.  However, there are users (especially in the IoT space) who desire and would benefit from being able to visualize a hierarchical relationship between their records' dimensions (Tags).  

In many cases, there is an intrinsic hierarchical relationship between records' dimensions but Line Protocol does not capture this.  

The goal of this tool is to extend the dimension relationship to enable the rendering of tree structures out of otherwise flat data.  

This should happen in two possible ways:

1) Without user input on what the hierarchy should be, the tool should, by default, generate a hierarchy where the levels/tiers follow a least-to-greatest cardinality.  
Example:

    A dataset has Measurement `cpu`, Fields `usage_user` and `usage_system`, and dimensions (Tags):
   | Tag    | Cardinality |
   |--------|-------------|
   | host   | 10          |
   | region | 2           |
   | app    | 4           |
   
   The tool would set the hierarchy to `region`>`app`>`host`.  

   This is great if the hierarchy is this straight forward but in the case above, there is a problem -- `app` and `host` probably don't have an intrinsic parent-child relationship.  These are Tags that can have a many:many relationship with one another and hierarchy is non-existent.  We would call these sibling dimensions and we'll address that with the following:

2) A user can input their own defintion of the hierarchy and, in doing so, can also define sibling relationships.  Sibling relationships will be on the same level and are simply different ways of slicing up data in the tree.

Proposals for user input:
* `MEASUREMENT,region,{host,app},FIELD`
* `MEASUREMENT>region>{host,app},FIELD`
* `MEASUREMENT|region|{host,app},FIELD`

Data may not follow a structure exactly like this, where `MEASUREMENT` is the top level node of the tree.  It could look like (using first proposal):

* `region,{host,app},MEASUREMENT,FIELD`

