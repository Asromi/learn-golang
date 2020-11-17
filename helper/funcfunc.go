package helper

import (
	"time"
)

//JSONTime ... declaration --ex: JSONTime(time.Now())
type JSONTime time.Time

/*
	var rCount int64
	for rows.Next() {
		rCount++
	}
	fmt.Println("rCount:", rCount)
	if rCount == 0 {
		return nil, err
	}
*/

// //Marshaler ... interfaces
// type Marshaler interface {
// 	MarshalJSON() ([]byte, error)
// }

// //MarshalJSON ... formatting time
// func (t JSONTime) MarshalJSON() ([]byte, error) {
// 	//do your serializing here

// 	//fmt.Printf("time1 %v\n", time1)    // time1 2015-09-01 17:59:31.73600891 +0700 WIB

// 	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
// 	return []byte(stamp), nil
// }

// if got != tt.want {
// 	t.Errorf("Foo(%q) = %d; want %d", tt.in, got, tt.want) // or Fatalf, if test can't test anything more past this point
// }

// func catch(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// convert (rows *sql.Rows) to JSON ([]byte)
// =========================================
// values := make([]interface{}, len(cols))
// data := make(map[string][]interface{})
// 	scanArgs := make([]interface{}, len(values))
// 	for i := range values {
// 		scanArgs[i] = &values[i]
// 	}

// c := 0
// results := make(map[string]interface{})
// data := []string{}
// for rows.Next() {
// 	if c > 0 {
// 		data = append(data, ",")
// 	}
// 	err = rows.Scan(scanArgs...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for i, value := range values {
// 		switch value.(type) {
// 		case nil:
// 			results[cols[i]] = nil

// 		case []byte:
// 			s := string(value.([]byte))
// 			x, err := strconv.Atoi(s)

// 			if err != nil {
// 				results[cols[i]] = s
// 			} else {
// 				results[cols[i]] = x
// 			}

// 		default:
// 			results[cols[i]] = value
// 		}
// 	}

// 	b, _ := json.Marshal(results)
// 	data = append(data, strings.TrimSpace(string(b)))
// 	c++
// }

// masterData := make(map[string][]interface{})
// for rows.Next() {
// 	err := rows.Scan(scanArgs...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for i, v := range values {
// 		//NOTE: FROM THE GO BLOG: JSON and GO - 25 Jan 2011:
// 		// The json package uses map[string]interface{} and []interface{} values to store arbitrary JSON objects and arrays;
// 		// it will happily unmarshal any valid JSON blob into a plain interface{} value.

// 		if nx, ok := strconv.ParseFloat(v.(string), 64); ok == nil {
// 			// float64 for JSON numbers,
// 			masterData[cols[i]] = append(masterData[cols[i]], nx)
// 		} else if b, ok := strconv.ParseBool(v.(string)); ok == nil {
// 			// bool for JSON booleans,
// 			masterData[cols[i]] = append(masterData[cols[i]], b)
// 		} else if "string" == fmt.Sprintf("%T", v.(string)) {
// 			// string for JSON strings, and
// 			masterData[cols[i]] = append(masterData[cols[i]], v.(string))
// 		} else {
// 			// nil for JSON null.
// 			fmt.Printf("Failed on if for type %T of %v\n", v.(string), v.(string))
// 		}
// 	}
// }

// for rows.Next() {
// 	for i := range values {
// 		values[i] = reflect.New(types[i]).Interface()
// 	}
// 	err = rows.Scan(values...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for i, v := range values {
// 		data[cols[i]] = append(data[cols[i]], v)
// 	}
// }

/*
reflect
	ct, err := rows.ColumnTypes()
	check(err, "execMSSQL, Error rows.ColumnTypes():")
	if err != nil {
		return nil, err
	}

	types := make([]reflect.Type, len(ct))
	for i, tp := range ct {
		st := tp.ScanType()
		if st == nil {
			fmt.Println("execMSSQL, Error reflect.Type: ", err.Error())
			return nil, err
		}
		types[i] = st
	}
	values := make([]interface{}, len(cols))
	for i := range values {
		values[i] = reflect.New(types[i]).Interface()
	}
*/

/*
	SQL Columns Type
	ct, err := rows.ColumnTypes()
	check(err, "execMSSQL, Error rows.ColumnTypes():")
	if err != nil {
		return nil, err
	}
	for i, tp := range ct {
		fmt.Println(i, tp)
	}
*/
