### Still alpha, some backward compatibility demages can happen.

Logan wraps [logrus](https://github.com/sirupsen/logrus/) and adds:

* Custom error with map for storing fields
* `WithStack` to log stack of an error
* `WithRecover` to log recover objects and retrieve stack from errors passed into panic

Synopsis:

```go
    rootLog := logan.New().Level(loganLogLevel).WithField("application", "appName")
    childLog := rootLog.WithField("service", "serviceName") // contains `application`
    clildLog.WithField("key", "value").WithError(err).WithStack(err).Error("Error happened.")
```


Fielded error usage example:

```go
package main

import (
	"gitlab.com/distributed_lab/logan/v2"
	"gitlab.com/distributed_lab/logan/v2/errors"
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
		return errors.WithField("car_color", carColor).Wrap(err, "Failed to start engine.")
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