package com.shutl.controller;

import com.shutl.model.Quote;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

import static org.springframework.web.bind.annotation.RequestMethod.POST;

@RestController
public class QuoteController {

    @RequestMapping(value = "/quote", method = POST)
    public @ResponseBody Quote quote(@RequestBody Quote quote) {
        Long price = Math.abs((Long.valueOf(quote.getDeliveryPostcode(), 36) - Long.valueOf(quote.getPickupPostcode(), 36))/100000000);

        return new Quote(quote.getPickupPostcode(), quote.getDeliveryPostcode(), price);
    }
}
