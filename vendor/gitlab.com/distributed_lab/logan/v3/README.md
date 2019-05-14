### Still alpha, some backward compatibility demages can happen.

Logan wraps [logrus](https://github.com/sirupsen/logrus/) and adds:

* Custom error with map for storing fields
* `WithStack` to log stack of an error
* `WithRecover` to log recover objects and retrieve stack from errors passed into panic

##Synopsis

####Log

```go
    rootLog := logan.New().Level(loganLogLevel).WithField("application", "appName")
    childLog := rootLog.WithField("service", "serviceName") // contains `application`
    clildLog.WithField("key", "value").WithError(err).WithStack(err).Error("Error happened.")
```

####Errors

```go
    plainError := errors.New("Error message.")
    
    wrapped := errors.Wrap(plainError, "Wrapping message")
    wrappedWithFields := errors.Wrap(plainError, "Wrapping message", logan.Field("key", "value").Add("key2", "value"))
    
    newErrorWithFields := errors.From(errors.New("Error message."), logan.Field("key", "value").Add("key2", "value"))
```


###Fielded error usage example:

```go
package main

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func main() {
	err := driveCar("Bob")
	if err != nil {
		log := logan.New()
		// Logan will log `car_color` here
		log.WithField("service", "car_manager").WithError(err).Error("Failed to start car.")
	}
}

func driveCar(driver string) error {
	var carColor string
	switch driver {
	case "John":
		// Only John drives blue car
		carColor = "BLUE"
	default:
		carColor = "RED"
	}

	err := startEngine(carColor)
	if err != nil {
	    // Mention `carColor` here, it is unknown from above.
		// Instead of logging the `carColor` here, put it into `err` as a field.
		return errors.Wrap(err, "Failed to start engine.", logan.F{"car_color": carColor})
	}

	return nil
}

// startEngine just returns simple error, if `carColor` is "RED".
func startEngine(carColor string) error {
	if carColor == "RED" {
	    // Do not add `carColor` into error here, it is known from above.
		return errors.New("Engine exploded.")
	}

	return nil
}
````



###To migrate from logan v1 to v3 do replaces:

Imports

`"gitlab.com/distributed_lab/logan"` --> `"gitlab.com/distributed_lab/logan/v3"\n\t"gitlab.com/distributed_lab/logan/v3/errors"` (do with regex) (Caution: this will also modify Gopkg files)

Wrap

`logan.Wrap(` --> `errors.Wrap(`

Wrap With fields

`errors.Wrap\((.+)\)\.[\n\t ]*WithFields\((.+)\)` --> `errors.Wrap($1, $2)` (do with regex)

Wrap With field

`errors.Wrap\(([a-zA-Z]+, *".+")\)\.[\n\t ]*WithField\((.+)\)((\.[\n\t ]*WithField\(.+\))*)` --> `errors.Wrap($1, logan.Field($2)$3)` (do with regex)

New With field

`errors.New\((".+")\)\.[\n\t ]*WithField\((.+)\)((\.[\n\t ]*WithField\(.+\))*)` --> `errors.From(errors.New($1), logan.Field($2)$3)` (do with regex)

For chained WithField calls - do multiple times
`logan.Field\((.+)\)\.[\n\t ]*WithField\((.+)\)` --> `logan.Field($1).Add($2)` (do with regex)

Remove errors import, where logan/v3/errors import exists

`\n\s"errors"((\n.*)*)"gitlab.com\/distributed_lab\/logan\/v3\/errors"` --> `$1"gitlab.com\/distributed_lab\/logan\/v3\/errors"` 

`\n\s"github.com/go-errors/errors"((\n.*)*)"gitlab.com\/distributed_lab\/logan\/v3\/errors"` --> `$1"gitlab.com\/distributed_lab\/logan\/v3\/errors"` 

`\n\s"github.com/pkg/errors"((\n.*)*)"gitlab.com\/distributed_lab\/logan\/v3\/errors"` --> `$1"gitlab.com\/distributed_lab\/logan\/v3\/errors"` 

Will then need to remove manually `"gitlab.com/distributed_lab/logan/v3/errors"` in some places (unused import)

If having some go errors imports - will also need to resolve manually.

###To replace deprecated `logan.Field(), fields.Add() and fields.AddFields()`:

`logan.Field\((\".+\"),[ 	]*([^)]+)\)` --> `logan.F{$1: $2}` (regexp)

`\.AddFields\(([^)]+)\)` --> `.Merge($1)` (regexp)

`\.Add\((\"[^.]+\"),[ 	]*([^)]+)\)` --> `[$1] = $2` (regexp)
Fix the appeared compilation errors.
