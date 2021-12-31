package main

// type Getter interface {
// 	getTagKeys() []string
// 	getValuesByTagKey() []string
// 	getField
// }

// func (m *MeasurementAPI) getTagKeyValues(tag string) []string {
// 	fluxGetVals := fmt.Sprintf(m.fluxGetVals, m.bucket, tag)
// 	result, err := m.api.Query(context.Background(), fluxGetVals)
// 	if err != nil {
// 		fmt.Printf("error querying for tag key values: %v", err)
// 	}
// 	// checkQueryError(err)
// 	var vals []string
// 	for result.Next() {
// 		checkQueryError(result.Err())
// 		strVal := fmt.Sprintf("%v", result.Record().Value())
// 		vals = append(vals, strVal)
// 	}
// 	return vals
// }
