var addon = require("bindings")("express-go");
var assert = require("assert");
var calculations = 10000000;

describe("addon", function() {
    it("should calculate pi synchronously", function() {
        var result = addon.calculateSync(calculations);
        assert.equal(Math.round(result * 1000), 3142);
    });
});
