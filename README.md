# Take home exercise

For our tech test, we'd like you to take a stripped-down version of our quoting engine, and then add some features. This is a RESTful service endpoint that takes a few details and works out the price for a delivery.

Throughout the test we're looking for great coding style, driving your code through tests (and refactoring) and at all times doing the bare minimum possible to get the job done. If you don't like the code or tests that are there already, feel free to refactor as you add features.

Please take your time and make sure that the features you complete are done to a standard that you're happy with. Please complete the features in order.

Read this document thoroughly before starting your work. You are welcome to contact us if you have any question.

Good luck! :)


### Submitting your work:

Submit your work as a pull request to this repository. You should be able to complete this in a day but you're free to come back to us in 3-5 days.

## Completed Feature

### Basic Service

Build a basic service that responds to a POST to /quotes, with the following request structure:

```
{
  "pickup_postcode":   "SW1A1AA",
  "delivery_postcode": "EC2A3LT"
}
```
And responds with the following price:
```
{
  "pickup_postcode":   "SW1A1AA",
  "delivery_postcode": "EC2A3LT",
  "price":             316
}
```

The price we charge depends on the distance between two postcodes. We are not implementing postcode geocoding here, so instead we are using basic formula for working out the price for a quote between two postcodes. The process is to take the base-36 integer of each postcode, subtract the delivery postcode from the pickup postcode and then divide by some large number. If the result is negative, turn it into a positive.

Hint: in java, this would be:

`Long.valueOf("SW1A1AA", 36) - Long.valueOf("EC2A3LT", 36)`

If you have a better idea for a deterministic way of making a number from two postcodes, please feel free to use that instead. Update your service to calculate pricing based upon these rules.

## Features to complete

### 1) Simple variable prices by vehicle

Our price changes based upon the vehicle. Implement a "vehicle" attribute on the request, that takes one of the following values, applying the appropriate markup:

* bicycle: 10%
* motorbike: 15%
* parcel_car: 20%
* small_van: 30%
* large_van: 40%

For example, if the base price was 100, the `small_van` price with markup will be 130.
The vehicle should also be returned in the response, and the price should be rounded to the nearest integer.

Request:
```
{
  "pickup_postcode":   "SW1A1AA",
  "delivery_postcode": "EC2A3LT",
  "vehicle": "bicycle"
}
```
Response:
```
{
  "pickup_postcode":   "SW1A1AA",
  "delivery_postcode": "EC2A3LT"
  "vehicle": "bicycle"
  "price": 348
}
```

### 2) Variable prices by carrier

Now we need the list of prices per carrier for the given `pickup_postcode`, `delivery_postcode` and `vehicle`.

Use the JSON file in the `src/data` folder to fetch the carrier data and calculate the price.
Bear in mind the carrier service should support the vehicle type. When calculating the price, add the service markup as well as the vehicle markup you have implemented in the earlier exercise to the carrier base price.

The `price_list` array needs to contain JSON objects sorted by `price`.

Example request:
```
{
  "pickup_postcode":   "SW1A1AA",
  "delivery_postcode": "EC2A3LT",
  "vehicle": "small_van"
}
```
Example response:
```
{
  "pickup_postcode":   "SW1A1AA",
  "delivery_postcode": "EC2A3LT"
  "vehicle": "small_van"
  "price_list": [
    {"service": "RoyalPackages", "price": 300, "delivery_time": 5}
    {"service": "Hercules", "price": 500, "delivery_time": 2},
  ]
}
```

### 3) Let's build a webpage!

Build a webpage that makes the above calls.

It should contain a form with the following fields:
`pickup_postcode`, `delivery_postcode` and `vehicle`.

Under the form, based on the response, list the services with the following info:
`service`, `price`, `delivery_time`
While the page is waiting for the response, an appropriate message should be displayed.
If there were no services returned, an appropriate message should be displayed.

**Bonus**:
- Make sure that the page displays well both on smaller and larger screens, ie that is `responsive`.
- The action linked to the submit button could retrieve the data from the service without refreshing the page.

# Dependencies

`gradle`: make sure is correctly installed on your machine. `brew` can help you with the installation if you are using a Mac Machine.

## Useful commands

Run tests from command line:
```
gradle test
```

Run server locally:
```
gradle bootRun
```

Make quote request:
```
echo '{"pickupPostcode": "SW1A1AA", "deliveryPostcode": "EC2A3LT" }' | \
curl -d @- http://localhost:8080/quote --header "Content-Type:application/json"
```
