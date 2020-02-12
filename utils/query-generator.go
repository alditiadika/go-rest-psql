package utils

import "fmt"

//NumericQuery func
func NumericQuery(field string, value []string, operator string) string {
	qs := "and ("
	if len(value) == 1 {
		switch operator {
		case "eq":
			return fmt.Sprintf(`and "%s" = %s `, field, value[0])
		case "neq":
			return fmt.Sprintf(`and "%s" <> %s `, field, value[0])
		case "lt":
			return fmt.Sprintf(`and "%s" < %s `, field, value[0])
		case "lte":
			return fmt.Sprintf(`and "%s" <= %s `, field, value[0])
		case "gt":
			return fmt.Sprintf(`and "%s" > %s `, field, value[0])
		case "gte":
			return fmt.Sprintf(`and "%s" >= %s `, field, value[0])
		default:
			return fmt.Sprintf(`and "%s" = %s `, field, value[0])
		}
	}
	for index, item := range value {
		temp := ""
		if index > 0 {
			temp += "or "
		}
		switch operator {
		case "eq":
			temp += fmt.Sprintf(`"%s" = %s `, field, item)
			break
		case "neq":
			temp += fmt.Sprintf(`"%s" <> %s `, field, item)
			break
		case "lt":
			temp += fmt.Sprintf(`"%s" < %s `, field, item)
			break
		case "lte":
			temp += fmt.Sprintf(`"%s" <= %s `, field, item)
			break
		case "gt":
			temp += fmt.Sprintf(`"%s" > %s `, field, item)
			break
		case "gte":
			temp += fmt.Sprintf(`"%s" >= %s `, field, item)
			break
		default:
			temp += fmt.Sprintf(`"%s" = %s `, field, item)
			break
		}
		if item == "" {
			temp = ""
		}
		qs += temp
	}
	qs += ") "
	return qs
}

//StringQuery func
func StringQuery(field string, value []string, operator string) string {
	qs := "and ("
	if len(value) == 1 {
		switch operator {
		case "contains":
			return fmt.Sprintf(`and "%s" like '%s' `, field, regexFormatterSQL("both", value[0]))
		case "not_contains":
			return fmt.Sprintf(`and "%s" not like '%s' `, field, regexFormatterSQL("both", value[0]))
		case "start_with":
			return fmt.Sprintf(`and "%s" like '%s' `, field, regexFormatterSQL("after", value[0]))
		case "end_with":
			return fmt.Sprintf(`and "%s" like '%s' `, field, regexFormatterSQL("before", value[0]))
		default:
			return fmt.Sprintf(`and "%s" = '%s' `, field, value[0])
		}
	}
	for index, item := range value {
		temp := ""
		if index > 0 {
			temp += "or "
		}
		switch operator {
		case "contains":
			temp += fmt.Sprintf(`"%s" like '%s' `, field, regexFormatterSQL("both", item))
			break
		case "not_contains":
			temp += fmt.Sprintf(`"%s" not like '%s' `, field, regexFormatterSQL("both", item))
			break
		case "start_with":
			temp += fmt.Sprintf(`"%s" like '%s' `, field, regexFormatterSQL("after", item))
			break
		case "end_with":
			temp += fmt.Sprintf(`"%s" like '%s' `, field, regexFormatterSQL("before", item))
			break
		default:
			temp += fmt.Sprintf(`"%s" = '%s' `, field, item)
			break
		}
		if item == "" {
			temp = ""
		}
		qs += temp
	}
	qs += ") "
	return qs
}

//BooleanQuery func
func BooleanQuery(field string, value string) string {
	return fmt.Sprintf(`and "%s" = %s `, field, value)
}

func regexFormatterSQL(typeReg string, val string) string {
	switch typeReg {
	case "before":
		return "%" + val
	case "after":
		return val + "%"
	case "both":
		return "%" + val + "%"
	default:
		return val
	}
}
