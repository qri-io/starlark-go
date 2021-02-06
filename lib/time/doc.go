// Package time defines time primitives for starlark, based heavily on the time
// package from the go standard library.
//
//  outline: time
//    time defines time primitives for starlark
//    path: time
//    functions:
//      parse_duration(string) duration
//        parse a duration
//      parse_location(string) location
//        parse a location
//      parse_time(string, format=..., location=...) time
//        parse a time
//      time(year=..., month=..., day=..., hour=..., minute=..., second=..., nanosecond=..., location=...) time
//        build a time
//      from_timestamp(int) time
//        parse a Unix timestamp
//      now() time
//        internal call to time.Time, returns the current time by default
//
//    types:
//      duration
//        fields:
//          hours float
//          microseconds int
//          milliseconds int
//          minutes float
//          nanoseconds int
//          seconds float
//        operators:
//          duration == duration = boolean
//          duration != duration = boolean
//          duration < duration = boolean
//          duration <= duration = boolean
//          duration > duration = boolean
//          duration >= duration = boolean
//          duration + duration = duration
//          duration + time = time
//          duration - duration = duration
//          duration / duration = float
//          duration / int = duration
//          duration // duration = int
//          duration // int = duration
//          duration * int = duration
//      time
//        functions:
//          year duration
//          month int
//          day int
//          hour int
//          minute int
//          second int
//          nanosecond int
//          unix int
//          unix_nano int
//          in_location(string) time
//            get time representing the same instant but in a different location
//          format(string) string
//            textual representation of time formatted according to the provided
//            layout string
//        operators:
//          time == time = boolean
//          time != time = boolean
//          time < time = boolean
//          time <= time = boolean
//          time > time = boolean
//          time >= time = boolean
//          time + duration = time
//          time - duration = time
//          time - time = duration
package time
