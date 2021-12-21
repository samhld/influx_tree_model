## Design

User defined hierarchy rules (call them "rules") written like:
```
MEASUREMENT>region>app>FIELD
```

A rule maps the items to hierarchical tiers.  In the above, `MEASUREMENT` is top tier, `region` is second from top and so on. 

The API will reach out to InfluxDB and collect the Measurements and associated Tag keys, values, and Fields and match them to tiers according to the defined rule.

In this example, the values in the `_measurement` column will be tier 1.  The values in the `region` column will be tier 2 and so on.  